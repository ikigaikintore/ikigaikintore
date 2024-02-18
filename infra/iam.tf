# gitops serviceAccount
resource "google_service_account" "gitops-sa" {
  account_id   = "terraform-gitops"
  display_name = "terraform gitops service account"
  project      = var.project_id
}

resource "google_iam_workload_identity_pool" "gitops-github-identity-pool" {
  workload_identity_pool_id = "gitops-github"
  display_name              = "GitHub Identity Pool"
  disabled                  = false
  project                   = var.project_id
}

locals {
  github_project_id       = "ikigaikintore"
  github_repository_owner = "ikigaikintore"
}

resource "google_iam_workload_identity_pool_provider" "gitops-github-pool-provider" {
  workload_identity_pool_id          = google_iam_workload_identity_pool.gitops-github-identity-pool.workload_identity_pool_id
  workload_identity_pool_provider_id = "gitops-github-pool"
  disabled                           = false
  display_name                       = "GitHub Pool Provider"

  attribute_condition = <<-EOT
      assertion.repository_owner == "${local.github_repository_owner}"
    EOT

  attribute_mapping = {
    "google.subject" = "assertion.repository"
  }

  oidc {
    issuer_uri = "https://token.actions.githubusercontent.com"
  }
  project = var.project_id
}

resource "google_project_iam_member" "gitops-sa-binding" {
  member  = "principal://iam.googleapis.com/${google_iam_workload_identity_pool.gitops-github-identity-pool.name}/subject/${local.github_repository_owner}/${local.github_project_id}"
  project = var.project_id
  role    = "roles/iam.workloadIdentityUser"
}

locals {
  roles = [
    "roles/editor",
    "roles/artifactregistry.writer",
    "roles/secretmanager.admin",
    "roles/iam.workloadIdentityUser",
    "roles/storage.admin",
  ]
}

resource "google_project_iam_member" "gitops-sa-roles" {
  project = var.project_id
  role    = each.value
  for_each = toset(local.roles)
  member  = "serviceAccount:${google_service_account.gitops-sa.email}"
}