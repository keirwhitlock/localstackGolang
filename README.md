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

tflocal init
tflocal apply
```

```shell
awslocal sqs get-queue-attributes --queue-url "http://sqs.eu-west-2.localhost.localstack.cloud:4566/000000000000/input-dlq" --attribute-names All
```

```shell
brew install protobuf
protoc --go_out=. *.proto
```

