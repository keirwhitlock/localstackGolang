package main

import (
	"encoding/base64"
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sqs"
	"google.golang.org/protobuf/proto"
	"log"
	"sqsExample/person"
	"sync"
	"time"
)

const sqsQueue = "http://sqs.eu-west-2.localhost.localstack.cloud:4566/000000000000/input"
const endpoint = "http://localhost:4566"

func main() {
	var wg sync.WaitGroup

	config := aws.NewConfig().WithEndpoint(endpoint).WithRegion("eu-west-2")
	client := sqs.New(session.Must(session.NewSessionWithOptions(session.Options{
		Profile: "localstack",
	})), config)

	c := make(chan string)

	go func() {

		defer close(c) // not sure if this is strictly needed in this case but better to be explicit(?)

		for {
			msg, err := client.ReceiveMessage(&sqs.ReceiveMessageInput{
				QueueUrl:        aws.String(sqsQueue),
				WaitTimeSeconds: aws.Int64(60),
			})
			if err != nil {
				panic(err)
			}
			if len(msg.Messages) == 0 {
				fmt.Println("Received no messages")
				time.Sleep(time.Duration(5 * time.Second))
				continue
			}

			for _, v := range msg.Messages {
				go func() {
					var user person.Person

					base64Decoded, err := base64.StdEncoding.DecodeString(*v.Body)
					if err != nil {
						panic(err)
					}

					err = proto.Unmarshal(base64Decoded, &user)
					if err != nil {
						log.Fatalln(err)
					}

					fmt.Println(user.String())
					c <- *v.ReceiptHandle
				}()
			}
		}
	}()

	for v := range c {
		wg.Add(1)
		go func() {
			_, err := client.DeleteMessage(&sqs.DeleteMessageInput{
				QueueUrl:      aws.String(sqsQueue),
				ReceiptHandle: aws.String(v),
			})
			if err != nil {
				panic(err)
			}
			wg.Done()
		}()
	}
	wg.Wait()
}
