packer {
  required_plugins {
    mvm = {
      version = "0.0.1"
      source  = "github.com/martezr/mvm"
    }
  }
}

locals {
  timestamp = formatdate("mmss", timestamp())
  vm_name   = "packiso-${local.timestamp}"
  boot_command = [
    "ipconfig"
  ]
}

source "mvm-iso" "demo" {
  url      = var.morpheus_url
  username = var.morpheus_username
  password = var.morpheus_password

  ssh_username = "root"

  cluster_name          = "mvmcluster01"
  vm_name               = local.vm_name
  labels                = ["packer", "automation"]
  convert_to_template   = false
  template_name         = "packertest"
  virtual_image_id      = 1018
  group                 = "platform engineering"
  cloud                 = "morpheuscloud"
  plan_id               = 174
  attach_virtio_drivers = true

  boot_command = local.boot_command
  boot_wait    = "2m"

  network_interface {
    network_id                = 229
    network_interface_type_id = 4
  }

  storage_volume {
    name            = "root"
    root_volume     = true
    size            = 25
    storage_type_id = 1
    datastore_id    = 49
  }

  ip_wait_timeout = "15m"

}

build {
  sources = [
    "source.mvm-iso.demo"
  ]
}
