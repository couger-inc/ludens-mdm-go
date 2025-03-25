import {
  to = aws_vpc.test_vpc
  id = "vpc-0094445a8abda30a8"
}
import {
  to = aws_security_group.elb_sg
  id = "sg-0665c31eafc1fe84b"
}
data "aws_subnets" "selected" {
  filter {
    name   = "tag:Name"
    values = ["ludens-mdm"] # insert values here
  }
}


data "aws_iam_policy_document" "lambda_assume_role" {
  statement {
    effect  = "Allow"
    actions = ["sts:AssumeRole"]
    principals {
      type        = "Service"
      identifiers = ["lambda.amazonaws.com"]
    }
  }
}

resource "aws_lambda_function" "get-managers-lambda-function" {
  function_name = "${var.app}-get-managers-lambda"
  timeout       = 5 # seconds
  image_uri     = "${aws_ecr_repository.get-managers-lambda.repository_url}:latest"
  package_type  = "Image"

  role = aws_iam_role.get-managers-lambda-function-role.arn

  environment {
    variables = {
      ENVIRONMENT = var.app
      DB_USERNAME = var.db_username
      DB_PASSWORD = var.db_password
      DB_NAME     = db_name
    }
  }

  vpc_config {
    subnet_ids         = selected.ids
    security_group_ids = [aws_security_group.elb_sg.id]
  }
}

resource "aws_lambda_function" "get-store-managers-lambda-function" {
  function_name = "${var.app}-get-store-managers-lambda"
  timeout       = 5 # seconds
  image_uri     = "${aws_ecr_repository.get-store-managers-lambda.repository_url}:latest"
  package_type  = "Image"

  role = aws_iam_role.get-managers-lambda-function-role.arn

  environment {
    variables = {
      ENVIRONMENT = var.app
    }
  }
}

resource "aws_lambda_function" "put-store-managers-lambda-function" {
  function_name = "${var.app}-put-store-managers-lambda"
  timeout       = 5 # seconds
  image_uri     = "${aws_ecr_repository.put-store-managers-lambda.repository_url}:latest"
  package_type  = "Image"

  role = aws_iam_role.get-managers-lambda-function-role.arn

  environment {
    variables = {
      ENVIRONMENT = var.app
    }
  }
}

resource "aws_lambda_function" "delete-store-managers-lambda-function" {
  function_name = "${var.app}-delete-store-managers-lambda"
  timeout       = 5 # seconds
  image_uri     = "${aws_ecr_repository.delete-store-managers-lambda.repository_url}:latest"
  package_type  = "Image"

  role = aws_iam_role.get-managers-lambda-function-role.arn

  environment {
    variables = {
      ENVIRONMENT = var.app
    }
  }
}

resource "aws_iam_role" "get-managers-lambda-function-role" {
  name = "${var.app}-get-managers-lambda-function-role"

  assume_role_policy = jsonencode({
    Statement = [
      {
        Action = "sts:AssumeRole"
        Effect = "Allow"
        Principal = {
          Service = "lambda.amazonaws.com"
        }
      },
    ]
  })
}
