# resource "google_project_iam_member" "saikumar_storage_viewer" {
#   project = var.project_id
#   role    = "roles/storage.admin"
#   member  = "user:${var.service_account_email}"
# }