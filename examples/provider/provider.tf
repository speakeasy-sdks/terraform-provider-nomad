terraform {
  required_providers {
    nomad = {
      source  = "hashicorp/nomad"
      version = "2.0.0"
    }
  }
}

provider "nomad" {
  # Configuration options
}