provider "google" {
  credentials = var.credentials
  project     = var.project_id
  zone        = var.zones[0]
}