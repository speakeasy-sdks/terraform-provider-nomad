terraform {
  required_providers {
    nomad = {
      source  = "hashicorp/nomad"
      version = "1.10.4"
    }
  }
}

provider "nomad" {
  # Configuration options
}