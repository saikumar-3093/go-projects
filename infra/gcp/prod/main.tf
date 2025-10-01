terraform {
    required_providers {
        google = {
        source = "hashicorp/google"
        version = "6.8.0"
        }
    }
}

provider "google" {
    project = "hardy-orb-468816-g1"
    region  = var.region
    zone    = var.zone
}

# module "vm_instance" {
#     source = "../../../modules/vm"

#     name         = var.name
#     machine_type = var.machine_type
#     image        = var.image
#     region       = var.region
#     zone         = var.zone
# }

module "google_storage_bucket" {
    source = "../../../modules/storage"

    bucket_name = var.bucket_name
    location = var.region
}