resource "aws_sqs_queue" "input" {
  name = "input"
  redrive_policy = jsonencode({
    deadLetterTargetArn = aws_sqs_queue.input_dlq.arn
    maxReceiveCount     = 1
  })
}

resource "aws_sqs_queue" "input_dlq" {
  name = "input-dlq"
}

