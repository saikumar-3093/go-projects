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
