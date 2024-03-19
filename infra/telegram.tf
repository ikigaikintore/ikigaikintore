resource "google_service_account" "telegram-sa" {
  account_id   = "telegram-sa"
  display_name = "telegram-sa"
  project      = var.project_id
}

locals {
  telegram_sa_roles = [
    "roles/run.invoker",
    "roles/secretmanager.secretAccessor",
  ]

  telegram_secrets = [
    "proxy_telegram_allowed_domains",
    "proxy_telegram_bot_token",
    "proxy_telegram_webhook_user_id",
  ]
}

# resource "google_project_iam_member" "telegram-sa-roles" {
#   member   = "serviceAccount:${google_service_account.telegram-sa.email}"
#   project  = var.project_id
#   for_each = toset(local.telegram_sa_roles)
#   role     = each.value
# }
#
# resource "google_secret_manager_secret" "telegram-sa-secret" {
#   secret_id = each.value
#   project   = var.project_id
#   for_each  = toset(local.telegram_secrets)
#   replication {
#     auto {}
#   }
# }