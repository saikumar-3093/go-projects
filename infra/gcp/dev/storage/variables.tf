variable region {
    type = string
    default = "asia-south1"
    }
 variable zone  {
    type = string
    default = "asia-south1-c"
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
    default = "terraform-bucket-3093"
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
    default = "terraform-project-470706"
    }

variable service_account_email {
    type = string   
    default = "saikumar.3093@gmail.com"
    }

variable key_ring_name {
    description = "The name of the KMS key ring."
    type        = string
    default     = "terraform-key-ring"
}