data "hsdp_config" "cf" {
  region      = var.region
  environment = var.environment
  service     = "cf"
}
