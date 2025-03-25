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

resource "aws_ecr_repository" "get-stores-lambda" {
  name = "${var.app}-get-stores-function"

  image_scanning_configuration {
    scan_on_push = true
  }
}

resource "aws_ecr_lifecycle_policy" "get-stores-lambda" {
  repository = aws_ecr_repository.get-stores-lambda.name

  policy = local.lifecycle_policy
}