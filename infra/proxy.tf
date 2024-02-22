resource "google_service_account" "proxy-sa" {
  account_id   = "proxy-sa"
  display_name = "proxy-sa"
  project      = var.project_id
}

locals {
  proxy_sa_roles = [
    "roles/run.invoker",
    "roles/secretmanager.secretAccessor",
  ]
}

resource "google_project_iam_member" "proxy-sa-roles" {
  member   = "serviceAccount:${google_service_account.proxy-sa.email}"
  project  = var.project_id
  for_each = toset(local.proxy_sa_roles)
  role     = each.value
}