
terraform {
  # The modules used in this example have been updated with 0.12 syntax, which means the example is no longer
  # compatible with any versions below 0.12.
  required_providers {
    google = {
      source = "hashicorp/google"
      version = "3.52.0"
    }
  }
  required_version = ">= 0.12.6"
  backend "gcs" {
    bucket = "gotfbucket"
    prefix = "devops-guidelines/goterraform-test"
  }
}

provider "google" {
  project = "dev"
  region = "us-central-1"
}

# ---------------------------------------------------------------------------------------------------------------------
# CREATE A CUSTOM SERVICE ACCOUNT FOR THE CD PIPELINES
# ---------------------------------------------------------------------------------------------------------------------

module "cd_serviceaccount" {
  source = "git@git.build.ingka.ikea.com:/terraform/gcp-service-account.git?ref=v0.0.1"
  name                  = "go-template-sa"
  project               = "dev"
  description           = "Templating test"
}

