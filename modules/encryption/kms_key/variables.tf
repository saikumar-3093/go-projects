variable key_ring {
    description = "The name of the KMS key."
    type        = string
    default     = "my-key" 
}

variable location {
    description = "The location where the KMS key ring and key will be created."
    type        = string
    default     = "asia"
}

variable project_id {
    description = "The GCP project ID where the KMS key ring and key will be created."
    type        = string
    default     = "terraform-project-470706"
}

variable key {
    description = "The KMS key ring ID."
    type        = string    
    default     = ""
}

variable key_name {
    description = "The KMS key name."
    type        = string    
    default     = "my-key"
}