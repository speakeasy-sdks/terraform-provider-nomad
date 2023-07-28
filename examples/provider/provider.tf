terraform {
  required_providers {
    nomad = {
      source  = "hashicorp/nomad"
      version = "1.5.0"
    }
  }
}

provider "nomad" {
  # Configuration options
}