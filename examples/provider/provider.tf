terraform {
  required_providers {
    nomad = {
      source  = "hashicorp/nomad"
      version = "1.19.0"
    }
  }
}

provider "nomad" {
  # Configuration options
}