terraform {
  required_providers {
    nomad = {
      source  = "hashicorp/nomad"
      version = "1.13.2"
    }
  }
}

provider "nomad" {
  # Configuration options
}