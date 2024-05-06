terraform {
  backend "s3" {}
}

module "tags" {
  source = "../../../modules/tags"

  region      = var.region
  environment = var.environment
  project     = var.project
  resource    = var.resource
}

provider "aws" {
  region  = var.region
  profile = var.profile

  default_tags {
    tags = module.tags.default_tags
  }
}