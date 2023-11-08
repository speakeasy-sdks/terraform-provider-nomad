terraform {
  required_providers {
    nomad = {
      source  = "hashicorp/nomad"
      version = "2.1.0"
    }
  }
}

provider "nomad" {
  # Configuration options
}