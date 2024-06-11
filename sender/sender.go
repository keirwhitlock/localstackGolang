package main

import (
	"encoding/base64"
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sqs"
	"github.com/google/uuid"
	"google.golang.org/protobuf/proto"
	"sqsExample/person"
	"sync"
)

const sqsQueue = "http://sqs.eu-west-2.localhost.localstack.cloud:4566/000000000000/input"
const endpoint = "http://localhost:4566"

func main() {

	var wg sync.WaitGroup

	config := aws.NewConfig().WithEndpoint(endpoint).WithRegion("eu-west-2")
	client := sqs.New(session.Must(session.NewSessionWithOptions(session.Options{
		Profile: "localstack",
	})), config)

	for i := 1; i < 100; i++ {
		wg.Add(1)
		go func() {
			uuid := uuid.New()

			user := person.Person{
				Userid: uuid.String(),
				Name:   "Billy Bob",
				Age:    int32(i),
			}

			userEncoded, err := proto.Marshal(&user)
			if err != nil {
				panic(err)
			}

			base64Encoded := base64.StdEncoding.EncodeToString(userEncoded)
			fmt.Println("Base64 Encoded Value: ", base64Encoded)

			_, err = client.SendMessage(&sqs.SendMessageInput{
				MessageBody: aws.String(base64Encoded),
				QueueUrl:    aws.String(sqsQueue),
			})
			if err != nil {
				panic(err)
			}

			wg.Done()
		}()
	}

	wg.Wait()
}
