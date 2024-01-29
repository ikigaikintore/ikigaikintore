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
  network                  = google_compute_network.internal-vpc-network.self_link
  project                  = var.project_id
  private_ip_google_access = true
}

# resource "google_vpc_access_connector" "internal-vpc-connector" {
#   name    = "connector"
#   network = google_compute_network.internal-vpc-network.self_link
# }

locals {
  network_apis = [
    "vpcaccess.googleapis.com",
  ]
}

resource "google_project_service" "network_apis" {
  project = var.project_id
  for_each = toset(local.network_apis)
  service = each.value

  disable_on_destroy = true
}