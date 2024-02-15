resource "google_service_account" "frontend-sa" {
  account_id   = "frontend-sa"
  display_name = "frontend"
  project      = var.project_id
}

locals {
  frontend_sa_roles = [
    "roles/run.developer",
  ]
  
  frontend_secrets = [
    "base_endpoint",
    "firebase_api_key",
    "firebase_auth_domain",
    "firebase_project_id",
    "firebase_storage_bucket",
    "firebase_messaging_sender_id",
    "firebase_app_id",
    "environment",
    "user_auth",
  ]
}

resource "google_secret_manager_secret" "frontend-sa-secret" {
  secret_id = each.value
  project   = var.project_id
  for_each  = toset(local.frontend_secrets)
  replication {
    auto {}
  }
}

resource "google_project_iam_member" "frontend-sa-roles" {
  for_each = toset(local.frontend_sa_roles)
  project  = var.project_id
  role     = each.value
  member   = "serviceAccount:${google_service_account.frontend-sa.email}"
}
