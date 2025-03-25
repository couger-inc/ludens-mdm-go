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