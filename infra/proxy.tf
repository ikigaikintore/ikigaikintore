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

  proxy_secrets = [
    "proxy_allowed_domains",
    "proxy_target_backend",
    "proxy_telegram_bot_token",
  ]
}

resource "google_project_iam_member" "proxy-sa-roles" {
  member   = "serviceAccount:${google_service_account.proxy-sa.email}"
  project  = var.project_id
  for_each = toset(local.proxy_sa_roles)
  role     = each.value
}

resource "google_secret_manager_secret" "proxy-sa-secret" {
  secret_id = each.value
  project   = var.project_id
  for_each  = toset(local.proxy_secrets)
  replication {
    auto {}
  }
}