locals {
  lifecycle_policy = <<EOF
{
    "rules": [
        {
            "rulePriority": 10,
            "description": "latest から始まるタグを持つイメージを5件残す",
            "selection": {
                "tagStatus": "tagged",
                "tagPrefixList": ["latest"],
                "countType": "imageCountMoreThan",
                "countNumber": 5
            },
            "action": {
                "type": "expire"
            }
        },
        {
            "rulePriority": 100,
            "description": "10件残す",
            "selection": {
                "tagStatus": "any",
                "countType": "sinceImagePushed",
                "countType": "imageCountMoreThan",
                "countNumber": 10
            },
            "action": {
                "type": "expire"
            }
        }
    ]
}
EOF
}

resource "aws_ecr_repository" "get-managers-lambda" {
  name = "${var.app}-get-managers-lambda-function"

  image_scanning_configuration {
    scan_on_push = true
  }
}

resource "aws_ecr_lifecycle_policy" "get-managers-lambda" {
  repository = aws_ecr_repository.get-managers-lambda.name

  policy = local.lifecycle_policy
}

resource "aws_ecr_repository" "get-store-managers-lambda" {
  name = "${var.app}-get-store-managers-lambda-function"

  image_scanning_configuration {
    scan_on_push = true
  }
}

resource "aws_ecr_lifecycle_policy" "get-store-managers-lambda" {
  repository = aws_ecr_repository.get-store-managers-lambda.name

  policy = local.lifecycle_policy
}

resource "aws_ecr_repository" "put-store-managers-lambda" {
  name = "${var.app}-put-store-managers-lambda-function"

  image_scanning_configuration {
    scan_on_push = true
  }
}

resource "aws_ecr_lifecycle_policy" "put-store-managers-lambda" {
  repository = aws_ecr_repository.put-store-managers-lambda.name

  policy = local.lifecycle_policy
}

resource "aws_ecr_repository" "delete-store-managers-lambda" {
  name = "${var.app}-delete-store-managers-lambda-function"

  image_scanning_configuration {
    scan_on_push = true
  }
}

resource "aws_ecr_lifecycle_policy" "delete-store-managers-lambda" {
  repository = aws_ecr_repository.delete-store-managers-lambda.name

  policy = local.lifecycle_policy
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
  image_uri     = "${data.aws_ecr_repository.get-managers-lambda.repository_url}:latest"
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