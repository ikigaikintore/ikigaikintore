resource "google_compute_network" "internal-vpc-network" {
  name                    = "int-vpc"
  auto_create_subnetworks = false
  routing_mode            = "REGIONAL"
  project                 = var.project_id
}

resource "google_vpc_access_connector" "internal-vpc-connector" {
  name           = "vpc-connector"
  network        = google_compute_network.internal-vpc-network.id
  region         = var.region
  ip_cidr_range  = "10.5.5.0/28"
  project        = var.project_id
  min_throughput = 200
  max_throughput = 500

  min_instances = 2
  max_instances = 4

  machine_type = "f1-micro"

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