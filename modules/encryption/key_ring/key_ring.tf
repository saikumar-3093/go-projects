resource "google_kms_key_ring" "key_ring" {
  name     = var.key_ring_name
  location = var.location
  project  = var.project_id
}

output "kms_key_ring_id" {
  value = google_kms_key_ring.key_ring.id
}

output "kms_key_ring_name" {
  value = google_kms_key_ring.key_ring.name
}