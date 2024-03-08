variable "project_id" {
  description = "ID for the project"
  type        = string
}

variable "region" {
  description = "region of the project"
  type        = string
}

variable "zone" {
  description = "zone for the region"
  type        = string
}

variable "billing_account" {
  description = "billing account"
  type        = string
}

variable "key_file" {}

variable "proxy_app_backend_target" {
  description = "The name of the backend service for the proxy app"
  type        = string
}