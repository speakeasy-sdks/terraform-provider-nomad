terraform {
  required_providers {
    nomad = {
      source  = "hashicorp/nomad"
      version = "1.10.3"
    }
  }
}

provider "nomad" {
  # Configuration options
}