variable "region" {
  description = "The region to deploy to."
  type        = string
  default     = "us-east"
}

variable "environment" {
  description = "The environment to deploy to."
  type        = string
  default     = "client-test"
}

variable "space_id" {
  description = "The ID of the space to deploy to."
  type        = string
}

variable "iam_org_id" {
  description = "The ID of the IAM organization."
  type        = string
}

variable "service_id" {
  description = "The ID of the service key to use for IAM operations."
  type        = string
}

variable "service_private_key" {
  description = "The private key of the service key to use for IAM operations."
  type        = string
}

variable "api_username" {
  description = "The username to use API authentication."
  type        = string
}

variable "api_password" {
  description = "The password to use API authentication."
  type        = string
}

variable "resource_prefix" {
  description = "The prefix to use for all resource names."
  type        = string
  default     = ""
}

variable "app_name" {
  description = "The name of the application."
  type        = string
  default     = "iamsale"
}

variable "app_hostname" {
  description = "The hostname of the application."
  type        = string
  default     = "iamsale"
}

variable "docker_image" {
  description = "The Docker image to deploy."
  type        = string
  default     = "ghcr.io/loafoe/iamsale"
}

variable "tag" {
  description = "The tag of the Docker image to deploy."
  type        = string
  default     = "latest"
}

variable "disk" {
  description = "The disk size to use for the application."
  type        = number
  default     = 1024
}

variable "memory" {
  description = "The memory size to use for the application."
  type        = number
  default     = 128
}
