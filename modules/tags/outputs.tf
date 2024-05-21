output "default_tags" {
  value = {
    "custom:region"      = var.region
    "custom:environment" = var.environment
    "custom:iac"         = var.iac
    "custom:project"     = var.project
    "custom:component"   = var.component
  }
}