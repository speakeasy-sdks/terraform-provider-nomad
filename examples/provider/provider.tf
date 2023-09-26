terraform {
  required_providers {
    nomad = {
      source  = "hashicorp/nomad"
      version = "1.16.0"
    }
  }
}

provider "nomad" {
  # Configuration options
}