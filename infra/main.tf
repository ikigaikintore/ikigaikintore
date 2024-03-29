locals {
  apis = [
    "artifactregistry.googleapis.com",
  ]
}

resource "google_project_service" "api-resources" {
  project  = var.project_id
  for_each = toset(local.apis)
  service  = each.value
}

resource "google_artifact_registry_repository" "artifact-repository" {
  project       = var.project_id
  location      = var.region
  repository_id = "ikigai"
  format        = "DOCKER"

  docker_config {
    immutable_tags = false
  }
}