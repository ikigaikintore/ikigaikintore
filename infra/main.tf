locals {
  apis = [
    "artifactregistry.googleapis.com",
  ]
}

resource "google_project_service" "api-resources" {
  project = var.project_id
  for_each = toset(local.apis)
  service = each.value
}