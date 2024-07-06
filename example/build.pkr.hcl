/*
packer {
  required_plugins {
    mvm = {
      version = "0.0.1"
      source  = "github.com/martezr/mvm"
    }
  }
}
*/

locals {
 timestamp = formatdate("mmss", timestamp())
}

source "mvm-clone" "demo" {
  url = "https://morpheus.test.local"
  username = "admin"
  password = "Password123#"
  vm_name = "pack-${local.timestamp}"
  template_name = "packertest"
  virtual_image_id = 439
  group_id = 1
  cloud_id = 1
  plan_id = 164
/*
  network {
    network_id = 15000
    disk_controller_index = 1
  }

  storage {
    disk_size = 15000
    disk_controller_index = 1
  }
*/
  communicator          = "ssh"
  ssh_username          = "ubuntu"
  ssh_password          = "Password123#"
  ssh_port = 22
}

build {
  sources = [
    "source.mvm-clone.demo"
  ]

  provisioner "shell" {
    inline = [
      "echo 'Password123#' | sudo -S sh -c 'sudo apt-get install -y nano'"
    ]
  }
}
