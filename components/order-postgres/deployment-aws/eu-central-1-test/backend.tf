terraform {
  backend "s3" {}
}

module "tags" {
  source = "../../../../modules/tags"

  region      = var.region
  environment = var.environment
  project     = var.project
  component   = var.component
}

provider "aws" {
  region = var.region

  default_tags {
    tags = module.tags.default_tags
  }
}