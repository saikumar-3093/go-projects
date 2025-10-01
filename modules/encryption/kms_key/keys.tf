terraform {
    required_providers {
        google = {
        source = "hashicorp/google"
        version = "6.8.0"
        }
    }
}

resource "google_kms_crypto_key" "crypto_key" {
  name            = var.key_name
  key_ring       = var.key
  rotation_period = "100000s"

  lifecycle {
    # prevent_destroy = true
  }
}


output "kms_key_id" {
  value = google_kms_crypto_key.crypto_key.id
}