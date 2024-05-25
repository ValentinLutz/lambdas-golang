module "base" {
  source = "../base"

  region      = var.region
  environment = var.environment
  project     = var.project

  public_subnets = {
    "ca013790-5097-4e97-a7b3-98aef0f6f0d9" : {
      cidr_block = "10.0.0.0/20", availability_zone = "eu-central-1a",
    },
    "3069b7b6-393a-48ef-97e7-508fb595d8f8" : {
      cidr_block = "10.0.16.0/20", availability_zone = "eu-central-1b",
    },
    "db1a9d26-e973-45ca-a8c4-c02596853861" : {
      cidr_block = "10.0.32.0/20", availability_zone = "eu-central-1c",
    }
  }
  private_subnets = {
    "700edcbb-1ca3-483d-8bbf-20d599ca4f38" : {
      cidr_block = "10.0.128.0/20", availability_zone = "eu-central-1a"
    },
    "75ebb059-7f5d-4b6f-917f-0eb95259397c" : {
      cidr_block = "10.0.160.0/20", availability_zone = "eu-central-1b"
    },
    "04ff306a-5527-4313-914e-41953b9298b4" : {
      cidr_block = "10.0.192.0/20", availability_zone = "eu-central-1c"
    }
  }
}