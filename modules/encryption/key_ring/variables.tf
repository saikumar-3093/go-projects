variable key_ring_name {
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