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

source "mvm-iso" "demo" {
  url = var.morpheus_url
  username = var.morpheus_username
  password = var.morpheus_password
  cluster_name = "mvmcluster01"
  convert_to_template = false
  skip_agent_install = true
  vm_name = "packiso-${local.timestamp}"
  template_name = "packertest"
  virtual_image_id = 935
  group_id = 3
  cloud_id = 139
  plan_id = 174

  network_interface {
    network_id = 229
    network_interface_type_id = 4
  }

  storage_volume {
    name = "root"
    root_volume = true
    size = 25
    storage_type_id = 1
    datastore_id = 49
  }
/*
  storage_volume {
    name = "data"
    root_volume = false
    size = 5
    storage_type_id = 1
    datastore_id = 15
  }
  */
  communicator          = "none"
}

build {
  sources = [
    "source.mvm-iso.demo"
  ]

/*
  provisioner "mvm-morpheus" {
    url = var.morpheus_url
    username = var.morpheus_username
    password = var.morpheus_password
    task_id = 2
  }
  */
}
