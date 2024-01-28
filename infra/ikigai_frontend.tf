resource "google_service_account" "frontend-sa" {
  account_id   = "frontend-sa"
  display_name = "frontend"
  project      = var.project_id
}

locals {
  frontend_sa_roles = [
    "roles/run.developer",
  ]
}

resource "google_project_iam_member" "frontend-sa-roles" {
  for_each = toset(local.frontend_sa_roles)
  project  = var.project_id
  role     = each.value
  member   = "serviceAccount:${google_service_account.frontend-sa.email}"
}
