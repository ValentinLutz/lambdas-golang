variable "region" {}
variable "environment" {}
variable "project" {}

variable "private_subnets" {
  type = map(object({
    cidr_block        = string
    availability_zone = string
  }))
}
variable "public_subnets" {
  type = map(object({
    cidr_block        = string
    availability_zone = string
  }))
}