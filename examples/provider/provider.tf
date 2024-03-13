terraform {
  required_providers {
    nomad = {
      source  = "hashicorp/nomad"
      version = "5.1.1"
    }
  }
}

provider "nomad" {
  # Configuration options
}