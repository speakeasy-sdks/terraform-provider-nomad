terraform {
  required_providers {
    nomad = {
      source  = "hashicorp/nomad"
      version = "2.5.0"
    }
  }
}

provider "nomad" {
  # Configuration options
}