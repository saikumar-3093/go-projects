variable region {
    type = string
    default = "asia-south1"
    }
 variable zone  {
    type = string
    default = "asia-south1-c"
    }

variable machine_type {
    type = string
    default = "e2-micro"
    }

variable image {
    type = string
    default = "debian-cloud/debian-11"
    }

variable name {
    type = string
    default = "terraform-instance"
    }

variable bucket_name {
    description = "The name of the storage bucket."
    type        = string
}

variable location {
    description = "The location where the bucket will be created."
    type        = string
    default     = "asia"
}

variable labels {
    description = "A map of labels to assign to the bucket."
    type        = map(string)
    default     = {}
}

variable project_id {
    type = string
    default = "hardy-orb-468816-g1"
    }

variable service_account_email {
    type = string   
    default = "saikumar.3093@gmail.com"
    }