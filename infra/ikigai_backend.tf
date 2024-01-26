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
  ]

  backend_apis = [
    "run.googleapis.com",
  ]
}

resource "google_project_iam_member" "backend-sa-roles" {
  for_each = toset(local.backend_sa_roles)
  project  = var.project_id
  role     = each.value
  member   = "serviceAccount:${google_service_account.backend-sa.email}"
}
