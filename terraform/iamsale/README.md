# iamsale Terraform module

Module for deploying the iamsale API to CF

## Usage

```hcl
module "iamsale" {
  source = "github.com/loafoe/terraform/iamsale"

  api_username = var.api_username
  api_password = var.api_password
  service_id   = var.service_id
  service_private_key = var.service_private_key
  iam_org_id   = var.org_id
}
```

## Documentation

<!-- BEGIN_TF_DOCS -->
## Requirements

| Name | Version |
|------|---------|
| <a name="requirement_cloudfoundry"></a> [cloudfoundry](#requirement\_cloudfoundry) | >= 0.51.3 |
| <a name="requirement_hsdp"></a> [hsdp](#requirement\_hsdp) | >= 0.45.0 |

## Providers

| Name | Version |
|------|---------|
| <a name="provider_cloudfoundry"></a> [cloudfoundry](#provider\_cloudfoundry) | 0.51.3 |
| <a name="provider_hsdp"></a> [hsdp](#provider\_hsdp) | 0.45.0 |

## Modules

No modules.

## Resources

| Name | Type |
|------|------|
| [cloudfoundry_app.iamsale](https://registry.terraform.io/providers/cloudfoundry-community/cloudfoundry/latest/docs/resources/app) | resource |
| [cloudfoundry_route.iamsale](https://registry.terraform.io/providers/cloudfoundry-community/cloudfoundry/latest/docs/resources/route) | resource |
| [hsdp_config.cf](https://registry.terraform.io/providers/philips-software/hsdp/latest/docs/data-sources/config) | data source |

## Inputs

| Name | Description | Type | Default | Required |
|------|-------------|------|---------|:--------:|
| <a name="input_api_password"></a> [api\_password](#input\_api\_password) | The password to use API authentication. | `string` | n/a | yes |
| <a name="input_api_username"></a> [api\_username](#input\_api\_username) | The username to use API authentication. | `string` | n/a | yes |
| <a name="input_app_hostname"></a> [app\_hostname](#input\_app\_hostname) | The hostname of the application. | `string` | `"iamsale"` | no |
| <a name="input_app_name"></a> [app\_name](#input\_app\_name) | The name of the application. | `string` | `"iamsale"` | no |
| <a name="input_disk"></a> [disk](#input\_disk) | The disk size to use for the application. | `number` | `1024` | no |
| <a name="input_docker_image"></a> [docker\_image](#input\_docker\_image) | The Docker image to deploy. | `string` | `"ghcr.io/loafoe/iamsale"` | no |
| <a name="input_environment"></a> [environment](#input\_environment) | The environment to deploy to. | `string` | `"client-test"` | no |
| <a name="input_iam_org_id"></a> [iam\_org\_id](#input\_iam\_org\_id) | The ID of the IAM organization. | `string` | n/a | yes |
| <a name="input_memory"></a> [memory](#input\_memory) | The memory size to use for the application. | `number` | `128` | no |
| <a name="input_region"></a> [region](#input\_region) | The region to deploy to. | `string` | `"us-east"` | no |
| <a name="input_resource_prefix"></a> [resource\_prefix](#input\_resource\_prefix) | The prefix to use for all resource names. | `string` | `""` | no |
| <a name="input_service_id"></a> [service\_id](#input\_service\_id) | The ID of the service key to use for IAM operations. | `string` | n/a | yes |
| <a name="input_service_private_key"></a> [service\_private\_key](#input\_service\_private\_key) | The private key of the service key to use for IAM operations. | `string` | n/a | yes |
| <a name="input_space_id"></a> [space\_id](#input\_space\_id) | The ID of the space to deploy to. | `string` | n/a | yes |
| <a name="input_tag"></a> [tag](#input\_tag) | The tag of the Docker image to deploy. | `string` | `"latest"` | no |

## Outputs

No outputs.
<!-- END_TF_DOCS -->
