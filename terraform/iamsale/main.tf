locals {
  app_name     = var.resource_prefix != "" ? "${var.resource_prefix}-${var.app_name}" : var.app_name
  app_hostname = var.resource_prefix != "" ? "${var.resource_prefix}-${var.app_hostname}" : var.app_hostname
}

resource "cloudfoundry_app" "iamsale" {
  name  = local.app_name
  space = var.space_id

  environment = {
    IAMSALE_CONFIG_BASE64 = templatefile("template/iamsale.yaml.tmpl",
      {
        api_username        = var.api_username
        api_password        = var.api_password
        service_id          = var.service_id
        service_private_key = var.service_private_key
        iam_org_id          = var.iam_org_id
    })
  }

  docker_image = "${var.docker_image}:${var.tag}"

  command = "echo $IAMSALE_CONFIG_BASE64 | base64 -d > /app/iamsale.yaml && /app/iamsale"

  disk_quota = var.disk
  memory     = var.memory

  routes {
    route = cloudfoundry_route.iamsale.id
  }
}

resource "cloudfoundry_route" "iamsale" {
  hostname = local.app_hostname
  domain   = data.hsdp_config.cf.domain
  space    = var.space_id
}
