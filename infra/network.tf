resource "google_compute_network" "vpc-network" {
  name                    = "vpc-network"
  project                 = var.project_id
  auto_create_subnetworks = true
}

resource "google_vpc_access_connector" "vpc-connector" {
  name          = "connector"
  project       = var.project_id
  region        = var.region
  network       = google_compute_network.vpc-network.name
  ip_cidr_range = "20.8.2.0/28"
}

locals {
  network_apis = [
    "compute.googleapis.com",
    "vpcaccess.googleapis.com",
  ]
}

resource "google_project_service" "network-api-resources" {
  for_each = toset(local.network_apis)
  project  = var.project_id
  service  = each.value
}