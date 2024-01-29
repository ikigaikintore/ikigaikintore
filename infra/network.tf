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
  network                  = google_compute_subnetwork.internal-vpc-subnet.self_link
  project                  = var.project_id
  private_ip_google_access = true
}