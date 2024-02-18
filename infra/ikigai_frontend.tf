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
    "environment",
    "user_auth",
  ]

  frontend_apis = [
    "identitytoolkit.googleapis.com",
    "firestore.googleapis.com",
    "serviceusage.googleapis.com",
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

# gcs
resource "google_storage_bucket_iam_binding" "frontend-sa-gcs" {
  bucket = google_storage_bucket.frontend-bucket.name
  role   = "roles/storage.objectViewer"
  members = [
    "allUsers",
  ]
}

resource "google_storage_bucket" "frontend-bucket" {
  name          = "web-iki"
  location      = var.region
  project       = var.project_id

  storage_class = "STANDARD"
  uniform_bucket_level_access = true

  website {
    main_page_suffix = "index.html"
    not_found_page   = "404.html"
  }

  lifecycle_rule {
    action {
      type = "Delete"
    }
    condition {
      age = 30
    }
  }
}
