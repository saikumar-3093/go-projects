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

variable key_ring_name {
    description = "The name of the KMS key ring."
    type        = string
    default     = "my-key-ring"
}

variable encryption_key {
    description = "The KMS key to use for bucket encryption."
    type        = string
    default     = ""
}

variable project_id {
    description = "The ID of the project in which to create the bucket."
    type        = string
}