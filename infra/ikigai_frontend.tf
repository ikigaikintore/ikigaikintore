resource "google_service_account" "frontend-sa" {
  account_id   = "frontend-sa"
  display_name = "frontend"
  project      = var.project_id
}

locals {
  frontend_sa_roles = [
    "roles/run.developer",
    "roles/secretmanager.secretAccessor",
    "roles/iam.serviceAccountUser",
    "roles/iam.serviceAccountTokenCreator",
  ]

  frontend_secrets = [
    "firebase_api_key",
    "firebase_auth_domain",
    "firebase_project_id",
    "firebase_storage_bucket",
    "firebase_messaging_sender_id",
    "firebase_app_id",
    "base_endpoint",
  ]

  frontend_apis = [
    "identitytoolkit.googleapis.com",
    "firestore.googleapis.com",
    "serviceusage.googleapis.com",
    "eventarc.googleapis.com",
  ]
}

resource "google_project_iam_member" "frontend-sa-roles" {
  for_each = toset(local.frontend_sa_roles)
  project  = var.project_id
  role     = each.value
  member   = "serviceAccount:${google_service_account.frontend-sa.email}"
}

resource "google_project_service" "frontend-api-resources" {
  for_each = toset(local.frontend_apis)
  project  = var.project_id
  service  = each.value
}

resource "google_secret_manager_secret" "frontend-secrets" {
  for_each  = toset(local.frontend_secrets)
  project   = var.project_id
  secret_id = each.value

  replication {
    auto {}
  }

  depends_on = [
    google_project_service.frontend-api-resources,
  ]
}
