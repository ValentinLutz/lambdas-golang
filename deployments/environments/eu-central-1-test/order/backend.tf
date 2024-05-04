terraform {
  backend "s3" {}
}

locals {
  default_tags = {
    "custom:resource" = var.resource
  }
}

provider "aws" {
  region  = var.region
  profile = var.profile

  default_tags {
    tags = merge(var.default_tags, local.default_tags)
  }
}