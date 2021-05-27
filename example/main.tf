provider "env" {}

resource "env_variable" "user" {
  name = "USER"
}

data "env_variable" "shell" {
  name = "SHELL"
}

terraform {
  required_version = "~> 0.13"
  required_providers {
    env = {
      source  = "github.com/tchupp/env"
    }
  }
}
