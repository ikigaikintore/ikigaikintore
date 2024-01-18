resource "google_compute_network" "shared-test" {
  name                    = "shared-test-2"
  auto_create_subnetworks = true
  project                 = var.project_id
}