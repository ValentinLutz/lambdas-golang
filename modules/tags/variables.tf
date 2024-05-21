variable "region" {
  type = string
}

variable "environment" {
  type = string
}

variable "project" {
  type = string
}

variable "component" {
  type = string
}

variable "iac" {
  type    = string
  default = "tofu"
}
