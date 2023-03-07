# data.docker_registry_image.ubuntu:
data "docker_registry_image" "ubuntu" {
  id                   = "sha256:9a0bdde4188b896a372804be2384015e90e3f84906b750c1a53539b585fbbe7f"
  insecure_skip_verify = false
  name                 = "ubuntu:jammy"
  sha256_digest        = "sha256:9a0bdde4188b896a372804be2384015e90e3f84906b750c1a53539b585fbbe7f"
}

# docker_image.ubuntu:
resource "docker_image" "ubuntu" {
  id            = "sha256:58db3edaf2be6e80f628796355b1bdeaf8bea1692b402f48b7e7b8d1ff100b02ubuntu:jammy"
  image_id      = "sha256:58db3edaf2be6e80f628796355b1bdeaf8bea1692b402f48b7e7b8d1ff100b02"
  name          = "ubuntu:jammy"
  pull_triggers = [
    "sha256:9a0bdde4188b896a372804be2384015e90e3f84906b750c1a53539b585fbbe7f",
  ]
  repo_digest   = "ubuntu@sha256:9a0bdde4188b896a372804be2384015e90e3f84906b750c1a53539b585fbbe7f"
}
