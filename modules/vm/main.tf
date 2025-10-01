terraform {
    required_providers {
        google = {
        source = "hashicorp/google"
        version = "6.8.0"
        }
    }
}

resource "google_compute_instance" "vm_instance" {
    name         = var.name
    machine_type = var.machine_type

    boot_disk {
        initialize_params {
            image = var.image
        }
    }

    network_interface {
        network = "default"
        access_config {
            # Leave this block empty to get a public IP.
        }
    }

    labels = {
        env = "dev"
    }
}