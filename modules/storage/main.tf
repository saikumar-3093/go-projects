terraform {
    required_providers {
        google = {
        source = "hashicorp/google"
        version = "6.8.0"
        }
    }
}

resource "google_storage_bucket" "b" {
  name     = var.bucket_name
  location = var.location
  project  = var.project_id
  uniform_bucket_level_access = true

  force_destroy = true

  versioning {
    enabled = true
  }

  lifecycle_rule {
    action {
      type = "Delete"
    }
    condition {
      age = 365
    }
  }

  labels = var.labels
  encryption {
    default_kms_key_name = var.encryption_key
  }
}