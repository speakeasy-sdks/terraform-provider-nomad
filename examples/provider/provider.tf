terraform {
  required_providers {
    nomad = {
      source  = "hashicorp/nomad"
      version = "3.0.4"
    }
  }
}

provider "nomad" {
  # Configuration options
}