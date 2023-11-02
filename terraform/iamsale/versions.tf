terraform {
  required_providers {
    cloudfoundry = {
      source  = "cloudfoundry-community/cloudfoundry"
      version = ">= 0.51.3"
    }
    hsdp = {
      source  = "philips-software/hsdp"
      version = ">= 0.45.0"
    }
  }
}
