terraform {
    required_providers {
        google = {
        source = "hashicorp/google"
        version = "6.8.0"
        }
    }
}

data "google_project" "current" {
  project_id = var.project_id
}

module "key_ring" {
  source       = "../../../../modules/encryption/key_ring"
  key_ring_name = "my-key-ring"
  location = var.location
  project_id = var.project_id
}

module "kms_key" {
  source       = "../../../../modules/encryption/kms_key"
  key = module.key_ring.kms_key_ring_id
  key_name = "${var.key_ring_name}-key"
}

module "storage_bucket" {
  source       = "../../../../modules/storage"
  bucket_name  = var.bucket_name
  project_id   = var.project_id
  location     = var.location
  labels       = var.labels
  encryption_key = module.kms_key.kms_key_id
}

# Grant GCS the ability to use the key
resource "google_kms_crypto_key_iam_member" "gcs_key_access" {
  crypto_key_id = module.kms_key.kms_key_id
  role          = "roles/cloudkms.cryptoKeyEncrypterDecrypter"
  member        = "serviceAccount:service-${data.google_project.current.number}@gs-project-accounts.iam.gserviceaccount.com"
}