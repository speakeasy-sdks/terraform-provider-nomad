terraform {
  required_providers {
    nomad = {
      source  = "hashicorp/nomad"
      version = "1.17.0"
    }
  }
}

provider "nomad" {
  # Configuration options
}