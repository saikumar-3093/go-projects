terraform { 
  backend "gcs" {
  bucket = "prod-tf-storage-bucket"
  prefix = "terraform/state"
}
}