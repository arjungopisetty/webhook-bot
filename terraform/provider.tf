terraform {
  required_providers {
    google = {
      source  = "hashicorp/google"
      version = "4.3.0"
    }
  }
}

provider "google" {
  project = "micro-enigma-347407"
  region  = "us-west1"
  zone    = "us-west1-a"
}
