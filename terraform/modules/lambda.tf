data "aws_vpc" "vpc" {
  filter {
    name = "tag:Name"
    values = ["ludens-mdm"]
  }
}

data "aws_subnets" "private_subnets" {
  filter {
    name = "tag:Name"
    values = ["ludens-mdm-private-*"]
  }
}

data "aws_subnets" "public_subnets" {
  filter {
    name = "tag:Name"
    values = ["ludens-mdm-public-*"]
  }
}

data "aws_security_groups" "test" {
  filter {
    name   = "group-name"
    values = ["lambda"]
  }

  filter {
    name   = "vpc-id"
    values = [data.aws_vpc.vpc.id]
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

resource "aws_iam_role_policy_attachment" "lambda_aws_lambda_vpc_access_execution_role" {
  role       = aws_iam_role.get-managers-lambda-function-role.arn
  policy_arn = "arn:aws:iam::aws:policy/service-role/AWSLambdaVPCAccessExecutionRole"
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
    }
  }

  vpc_config {
    subnet_ids         = data.aws_subnets.private_subnets.ids
    security_group_ids = data.aws_security_groups.test.ids
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
