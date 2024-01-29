resource "google_compute_network" "internal-vpc-network" {
  name                    = "int-vpc"
  auto_create_subnetworks = false
  routing_mode            = "REGIONAL"
  project                 = var.project_id
}

resource "google_compute_subnetwork" "internal-vpc-subnet" {
  name                     = "int-vpc-subnet"
  ip_cidr_range            = "10.5.5.0/24"
  region                   = var.region
  network                  = google_compute_network.internal-vpc-network.id
  project                  = var.project_id
  private_ip_google_access = true
}

resource "google_vpc_access_connector" "internal-vpc-connector" {
  name           = "vpc-connector"
  network        = google_compute_network.internal-vpc-network.id
  region         = var.region
  ip_cidr_range  = "10.5.5.0/28"
  project        = var.project_id
  min_throughput = 200
  max_throughput = 500

  min_instances = 0
  max_instances = 1

  machine_type = "e2-small"

  depends_on = [
    google_project_service.network_apis
  ]
}

locals {
  network_apis = [
    "vpcaccess.googleapis.com",
  ]
}

resource "google_project_service" "network_apis" {
  project  = var.project_id
  for_each = toset(local.network_apis)
  service  = each.value

  disable_on_destroy = true
}