terraform {
  required_providers {
    nomad = {
      source  = "hashicorp/nomad"
      version = "2.4.1"
    }
  }
}

provider "nomad" {
  # Configuration options
}