module "s3_bucket_name" {
  source = "../name"

  region      = var.region
  environment = var.environment
  project     = var.project
  name        = "tofu-state"
}

module "dynamodb_name" {
  source = "../name"

  region      = var.region
  environment = var.environment
  project     = var.project
  name        = "tofu-state-lock"
}