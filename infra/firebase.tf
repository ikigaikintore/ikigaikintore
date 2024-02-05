locals {
  firebase_apis = [
    "firebase.googleapis.com",
    "serviceusage.googleapis.com",
  ]
}

resource "google_project_service" "firebase_apis" {
  for_each = toset(local.firebase_apis)
  project  = var.project_id
  service  = each.value

  disable_on_destroy = true
}
