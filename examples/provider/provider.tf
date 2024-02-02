terraform {
  required_providers {
    nomad = {
      source  = "hashicorp/nomad"
      version = "4.0.0"
    }
  }
}

provider "nomad" {
  # Configuration options
}