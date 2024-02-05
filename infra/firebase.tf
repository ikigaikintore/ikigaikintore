locals {
  firebase_apis = [
    "firebase.googleapis.com",
  ]
}

resource "google_project_service" "firebase_apis" {
  for_each = toset(local.firebase_apis)
  project  = var.project_id
  service  = each.value

  disable_on_destroy = true
}

resource "google_firebase_project" "ikigaikintore" {
  provider = google-beta
  project  = var.project_id

  depends_on = [google_project_service.firebase_apis]
}

resource "google_project_service" "firebase-auth" {
  provider = google-beta
  project  = var.project_id
  service  = "identitytoolkit.googleapis.com"

  disable_on_destroy = true
}

resource "google_identity_platform_config" "firebase-auth-google" {
  provider = google-beta
  project  = var.project_id

  autodelete_anonymous_users = true

  depends_on = [
    google_project_service.firebase-auth
  ]

  sign_in {
    allow_duplicate_emails = false

    anonymous {
      enabled = false
    }

    email {
      enabled = false
    }

    phone_number {
      enabled = false
    }
  }

  sms_region_config {
    allowlist_only {
      allowed_regions = [
        "JP"
      ]
    }
  }
}