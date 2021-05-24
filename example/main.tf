provider "env" {}

resource "env_variable" "user" {
  name = "USER"
}

output "user_resource" {
  value = env_variable.user
}

data "env_variable" "shell" {
  name = "SHELL"
}

output "shell_data" {
  value = data.env_variable.shell
}

terraform {
  required_version = "~> 0.15"
  required_providers {
    env = {
      source  = "github.com/tchupp/env"
    }
  }
}
