variable "project_id" {
  description = "ID for the project"
  type = string
}

variable "region" {
  description = "region of the project"
  type = string
}

variable "zones" {
  description = "zones for the region"
  type = set(
    string
  )
}

variable "credentials" {}
