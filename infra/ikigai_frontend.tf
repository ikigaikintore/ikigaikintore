resource "google_service_account" "frontend-sa" {
  account_id   = "frontend-sa"
  display_name = "frontend"
  project      = var.project_id
}
