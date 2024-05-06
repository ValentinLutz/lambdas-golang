variable "region" {
  type = string
}

variable "environment" {
  type = string
}

variable "resource" {
  type = string
}

variable "project" {
  type = string
}

variable "iac" {
  type    = string
  default = "tofu"
}
