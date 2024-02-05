resource "google_project" "default" {
  name       = "ikigaikintore"
  project_id = var.project_id

  labels = {
    "firebase" = "enabled"
  }

  billing_account = var.billing_account
}

locals {
  firebase_apis = [
    "firebase.googleapis.com",
    "serviceusage.googleapis.com",
  ]
}

resource "google_project_service" "firebase_apis" {
  for_each = toset(local.firebase_apis)
  project  = google_project.default.project_id
  service  = each.value

  disable_on_destroy = true
}
