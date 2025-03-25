# couger-dev 環境

terraform {
  backend "s3" {
    bucket = "couger-dev-terraform-states"
    key    = "ludens-mdm-couger-dev-go.tfstate"
    region = "ap-northeast-1"
    assume_role = {
      role_arn = "arn:aws:iam::229484968889:role/TerraformRole"
    }
  }
}

locals {
  account_id       = "229484968889"
  app              = "ludens-mdm"
  env              = "couger-dev"
  zone_name        = "couger-dev.ludens.to"
  admin_web_acl_id = "c9b4f2e0-10ca-4e57-b8e7-b859128df73e"
  referer_value    = "3sw0FBt2NHXrQUzN0aVxZWqaym69uohi"
}


provider "aws" {
  region = "ap-northeast-1"
  default_tags {
    tags = {
      "App" = local.app
      "Env" = local.env
    }
  }

  assume_role {
    role_arn = "arn:aws:iam::${local.account_id}:role/TerraformRole"
  }
}

provider "aws" {
  alias  = "us-east-1"
  region = "us-east-1"
  default_tags {
    tags = {
      "App" = local.app
      "Env" = local.env
    }
  }

  assume_role {
    role_arn = "arn:aws:iam::${local.account_id}:role/TerraformRole"
  }
}

module "backend" {
  source = "../../modules"
  providers = {
    aws           = aws
    aws.us-east-1 = aws.us-east-1
  }

  app                        = local.app
  env                        = local.env

}