terraform {
  required_providers {
    nomad = {
      source  = "hashicorp/nomad"
      version = "1.8.0"
    }
  }
}

provider "nomad" {
  # Configuration options
}