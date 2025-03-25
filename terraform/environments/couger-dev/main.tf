# couger-dev 環境

terraform {
  backend "s3" {
    bucket = "couger-dev-terraform-states"
    key    = "ludens-mdm-couger-dev.tfstate"
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
  vpc_cidr_block             = "10.1.0.0/16"
  availability_zones         = ["ap-northeast-1a", "ap-northeast-1c", "ap-northeast-1d"]
  subnet_public_cidr_blocks  = ["10.1.11.0/24", "10.1.12.0/24", "10.1.13.0/24"]
  subnet_private_cidr_blocks = ["10.1.21.0/24", "10.1.22.0/24", "10.1.23.0/24"]
  nat_gw_multi_az            = false
  vpc_endpoint_s3_additional_resources = [
    "arn:aws:s3:::sonar-couger-dev-data",
    "arn:aws:s3:::sonar-couger-dev-data/*",
    "arn:aws:s3:::sonar-couger-dev-cache",
    "arn:aws:s3:::sonar-couger-dev-cache/*",
  ]

  data_bucket  = "${local.app}-${local.env}-data"
  log_bucket   = "${local.app}-${local.env}-log"
  cache_bucket = "${local.app}-${local.env}-cache"

  monitoring_topic_arn = "arn:aws:sns:ap-northeast-1:${local.account_id}:monitoring"

  db_engine_version            = "8.0.mysql_aurora.3.07.1"
  db_master_username           = "root"
  db_master_password           = "8D#t2eb_P.XUYWL+"
  db_username                  = "app"
  db_password                  = "rK6w2-mi99hxmaP,"
  db_name                      = "ludens_mdm_couger_dev"
  db_min_capacity              = 0.5
  db_max_capacity              = 2.0
  db_instance_count            = 1
  db_backup_retention_period   = 7    # days
  db_backtrack_window          = 3600 # 1h
  db_snapshot_identifier       = null
  db_enable_http_endpoint      = true
  prisma_connection_parameters = "connection_limit=20&pool_timeout=10"

  zone_name               = local.zone_name
  private_zone_name       = "ludens-mdm-private.${local.zone_name}"
  alb_backend_domain_name = "alb.mdm-backend.${local.zone_name}"
  backend_domain_name     = "mdm-backend.${local.zone_name}"
  admin_web_acl_id        = "arn:aws:wafv2:us-east-1:${local.account_id}:global/webacl/admin/${local.admin_web_acl_id}"
  referer_value           = local.referer_value

  image_tag                  = "latest"
  backend_task_cpu           = 1024
  backend_task_memory        = 2048
  backend_task_desired_count = 1
  backend_task_max_count     = 2
  backend_task_min_count     = 1
  cors_allowed_origins       = ["https://mdm.${local.zone_name}", "https://mdm.localhost.${local.zone_name}:3100"]
  identity_api_base_path     = "https://admin-backend.${local.zone_name}"
  service                    = local.app
  service_kid                = "key-1"
  service_key                = "12345678901234567890123456789012"
  cookie_domain              = local.zone_name
  firebase_project_id        = "ludens-couger-dev"
  sns_topic_ludens_arn       = "arn:aws:sns:ap-northeast-1:${local.account_id}:ludens-user-console-couger-dev-ludens.fifo"
  log_level                  = "info"

  bastion_instance_count = 1
}

module "frontend" {
  source = "../../modules/frontend"
  providers = {
    aws           = aws
    aws.us-east-1 = aws.us-east-1
  }

  app              = local.app
  env              = local.env
  frontend_bucket  = "${local.app}-${local.env}-frontend"
  zone_name        = local.zone_name
  domain_name      = "mdm.${local.zone_name}"
  min_ttl          = 10
  default_ttl      = 10
  max_ttl          = 10
  admin_web_acl_id = "arn:aws:wafv2:us-east-1:${local.account_id}:global/webacl/admin/${local.admin_web_acl_id}"
  log_bucket       = module.backend.log_bucket
}

output "vpc" {
  value = module.backend.vpc
}

output "nat_gateway_ips" {
  value = module.backend.nat_gateway_ips
}