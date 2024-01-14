provider "google" {
  credentials = file(var.key_file)
  project     = var.project_id
  zone        = var.zone
  region      = var.region
}

terraform {
  required_providers {
    google = {
      version = "~> 5.11"
    }
  }
}
