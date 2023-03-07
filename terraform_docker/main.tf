terraform {
  required_providers {
    docker = {
      source = "kreuzwerker/docker"
    }
  }
}



data "docker_registry_image" "ubuntu" {
  name = "ubuntu:jammy"
}

resource "docker_image" "ubuntu" {
  name          = data.docker_registry_image.ubuntu.name
  pull_triggers = [data.docker_registry_image.ubuntu.sha256_digest]
}