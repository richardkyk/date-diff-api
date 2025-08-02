terraform {
  required_providers {
    aws = {
      source  = "hashicorp/aws"
      version = "~> 6.7"
    }
  }
  required_version = ">= 1.12.0"
}

locals {
  image_tag = filesha256("../cmd/lambda/Dockerfile")
}

# Create ECR repository to hold your Docker image
resource "aws_ecr_repository" "lambda_repo" {
  name = "${var.app_name}-repo"
}

# Build and push Docker image to ECR
resource "null_resource" "docker_build_and_push" {
  depends_on = [aws_ecr_repository.lambda_repo]

  triggers = {
    dockerfile_hash = local.image_tag
  }

  provisioner "local-exec" {
    command = <<EOT
      aws ecr get-login-password --region ${var.aws_region} | docker login --username AWS --password-stdin ${aws_ecr_repository.lambda_repo.repository_url}
      docker build -f ../cmd/lambda/Dockerfile -t ${aws_ecr_repository.lambda_repo.repository_url}:${local.image_tag} ../
      docker push ${aws_ecr_repository.lambda_repo.repository_url}:${local.image_tag}
    EOT
  }
}

# Create IAM role for Lambda execution
resource "aws_iam_role" "lambda_exec" {
  name               = "lambda-exec-role"
  assume_role_policy = data.aws_iam_policy_document.lambda_assume_role.json
}

data "aws_iam_policy_document" "lambda_assume_role" {
  statement {
    effect = "Allow"
    principals {
      type        = "Service"
      identifiers = ["lambda.amazonaws.com"]
    }
    actions = ["sts:AssumeRole"]
  }
}

# Attach managed policies to the role
resource "aws_iam_role_policy_attachment" "lambda_basic" {
  role       = aws_iam_role.lambda_exec.name
  policy_arn = "arn:aws:iam::aws:policy/service-role/AWSLambdaBasicExecutionRole"
}

# Create Lambda function from container image
resource "aws_lambda_function" "lambda" {
  depends_on = [null_resource.docker_build_and_push]

  function_name = "${var.app_name}-lambda"
  package_type  = "Image"
  image_uri     = "${aws_ecr_repository.lambda_repo.repository_url}:${local.image_tag}"
  role          = aws_iam_role.lambda_exec.arn
  timeout       = 10
  memory_size   = 512

}

resource "aws_lambda_function_url" "default" {
  function_name      = aws_lambda_function.lambda.function_name
  authorization_type = "NONE" # Or "AWS_IAM" for secured access
}

output "lambda_url" {
  value = aws_lambda_function_url.default.function_url
}
