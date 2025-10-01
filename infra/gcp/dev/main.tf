terraform {
    required_providers {
        google = {
        source = "hashicorp/google"
        version = "6.8.0"
        }
    }
}

provider "google" {
    project = var.project_id
    region  = var.region
    zone    = var.zone
}

module "vm_instance" {
    source = "../../../modules/vm"

    name         = "tf-preprod-vm"
    machine_type = var.machine_type
    image        = var.image
    region       = var.region
    zone         = var.zone
}

# module "storage_bucket" {
#     source = "../../../modules/storage"

#     bucket_name = var.bucket_name
#     location    = var.location
#     labels      = var.labels
# }