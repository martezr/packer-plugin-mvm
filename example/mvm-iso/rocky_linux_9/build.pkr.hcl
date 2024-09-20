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
    "<up><tab> text ip={{ .StaticIP }}::{{ .StaticGateway }}:{{ .StaticMask }}:${local.vm_name}:ens3:none nameserver={{ .StaticDNS }} inst.ks=http://{{ .HTTPIP }}:{{ .HTTPPort }}/ks.cfg<enter><wait><enter>"
  ]
}

source "mvm-iso" "rocky-linux-demo" {
  url      = var.morpheus_url
  username = var.morpheus_username
  password = var.morpheus_password

  // MVM Cluster
  cluster_name     = "mvmcluster01"
  vm_name          = local.vm_name
  description      = "packer test instance"
  environment      = "dev"
  labels           = ["packer", "automation"]
  virtual_image_id = 1021
  group_id         = 3
  cloud_id         = 139
  plan_id          = 174

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

  storage_volume {
    name            = "data"
    root_volume     = false
    size            = 10
    storage_type_id = 1
    datastore_id    = 49
  }

  convert_to_template = true
  template_name       = "rocky9mvm"

  boot_command            = local.boot_command
  boot_wait               = "5s"
  http_interface          = "en0"
  http_directory          = "${path.root}/http"
  http_template_directory = "${path.root}/http_templates"
  http_port_max           = 8030
  http_port_min           = 8020


  ip_wait_timeout = "15m"

  // Provisioner settimgs
  ssh_timeout  = "55m"
  communicator = "ssh"
  ssh_username = "root"
  ssh_password = "mysecurepassword"
}

build {
  sources = [
    "source.mvm-iso.rocky-linux-demo"
  ]

  provisioner "shell" {
    script = "scripts/setup.sh"
  }
}
