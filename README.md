# Localstack local AWS development

## Setup

You do need docker installed.

```shell
brew install tfenv
tfenv use 1.8.5

brew install localstack/tap/localstack-cli
pip3 install terraform-local
```

```shell
localstack start -d
```

```shell
tflocal init
tflocal apply -auto-approve 
```

```shell
awslocal sqs get-queue-attributes --queue-url "http://sqs.eu-west-2.localhost.localstack.cloud:4566/000000000000/input-dlq" --attribute-names All
```

```shell
brew install protobuf
protoc --go_out=. *.proto

# if you cant use `go_package` var in proto file, you can pass it into the protoc cmd
# protoc --go_out=. --go_opt=Mperson.proto=person/ person.proto 
```

