resource "google_service_account" "backend-sa" {
  account_id   = "backend-sa"
  display_name = "backend"
  project      = var.project_id
}

locals {
  backend_sa_roles = [
    "roles/iam.serviceAccountUser",
    "roles/iam.serviceAccountTokenCreator",
    "roles/run.developer",
    "roles/secretmanager.secretAccessor",
    "roles/monitoring.viewer",
    "roles/logging.viewer",
    "roles/artifactregistry.reader",
  ]

  backend_apis = [
    "run.googleapis.com",
    "secretmanager.googleapis.com",
    "gmail.googleapis.com",
  ]

  backend_secrets = [
    "weather_api_key",
  ]
}

resource "google_project_iam_member" "backend-sa-roles" {
  for_each = toset(local.backend_sa_roles)
  project  = var.project_id
  role     = each.value
  member   = "serviceAccount:${google_service_account.backend-sa.email}"
}

resource "google_project_service" "backend-api-resources" {
  project  = var.project_id
  for_each = toset(local.backend_apis)
  service  = each.value
}

resource "google_secret_manager_secret" "backend-sa-secret" {
  secret_id = each.value
  project   = var.project_id
  for_each  = toset(local.backend_secrets)
  replication {
    auto {}
  }

  depends_on = [
    google_project_service.backend-api-resources
  ]
}

#resource "google_cloud_run_service_iam_binding" "proxy-invoker" {
#  members = [
#    "serviceAccount:${google_service_account.proxy-sa.email}"
#  ]
#  role    = "roles/run.invoker"
#  service = var.proxy_app_backend_target
#  project = var.project_id
#}
