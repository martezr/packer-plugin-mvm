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
/*  boot_command = [
    "root<enter><wait>",
    "setup-alpine<enter>",
    "us<enter>",
    "us<enter>",
    "alpinetemp<enter>",
    "eth0<enter>",
    "dhcp<enter>",
    "n<enter>",
    "password123<enter>", // Enter Password
    "password123<enter>", // Confirm Password
    "<enter>",             // Timezone - Lowercase doesn't register properly
    "none<enter>",         // Proxy
    "r<enter>",            // APK Mirror
    "no<enter>",           // Create user
    "openssh<enter>",      // SSH Server
    "yes<enter>",          // Allow root login
    "none<enter>",         // SSH Key
    "sda<enter>",          // Install DISK
    "sys<enter>",          // Install DISK
    "y<enter>",            // Wipe Disk
    "mount /dev/sda3 /mnt<enter>",
    "setup-apkrepos -c -1<enter>",
    "apk update<enter>",
    "apk -U add qemu-guest-agent && rc-service qemu-guest-agent start && rc-update add qemu-guest-agent<enter>",
    "umount /mnt<enter>",
    "reboot<enter>"              // Reboot System
  ]*/
    vm_name = "packiso-${local.timestamp}"
    boot_command = [
      "e<wait><down><down><down><end>",
      "ip={{ .StaticIP }}::{{ .StaticGateway }}:{{ .StaticMask }}:${local.vm_name}:::{{ .StaticDNS }} autoinstall 'ds=nocloud-net;s=http://{{ .HTTPIP }}:{{ .HTTPPort }}/'", 
      "<wait><F10><wait>"
    ]
    static_ip = "{{ .StaticIP }}"
}

source "mvm-iso" "demo" {
  url = var.morpheus_url
  username = var.morpheus_username
  password = var.morpheus_password
  cluster_name = "mvmcluster01"
  boot_command = local.boot_command
  boot_wait = "5s"
  convert_to_template = false
  skip_agent_install = true
  vm_name = local.vm_name
  template_name = "packertest"
  virtual_image_id = 1015
  group_id = 3
  cloud_id = 139
  plan_id = 174

  http_interface       = "en0"
  http_directory = "${path.root}/http"
  http_template_directory = "${path.root}/http_templates"
  //http_content            = {
  //  "/user-data" = templatefile("${path.root}/http/user-data.pkrtpl", { ip_address = "{{ .StaticIP }}"})
  //}
  http_port_max        = 8030
  http_port_min        = 8020

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
  # Raise the timeout, when installation takes longer
  ssh_timeout = "55m"
  communicator = "ssh"
  ssh_username = "ubuntu"
  ssh_password = "password123"
}

build {
  sources = [
    "source.mvm-iso.demo"
  ]

  provisioner "shell" {
      inline = ["echo foo"]
  }

   provisioner "shell" {
        inline = [
            "while [ ! -f /var/lib/cloud/instance/boot-finished ]; do echo 'Waiting for cloud-init...'; sleep 1; done",
            "sudo rm /etc/ssh/ssh_host_*",
            "sudo truncate -s 0 /etc/machine-id",
            "sudo apt -y autoremove --purge",
            "sudo apt -y clean",
            "sudo apt -y autoclean",
            "sudo cloud-init clean",
            "sudo rm -f /etc/cloud/cloud.cfg.d/subiquity-disable-cloudinit-networking.cfg",
            "sudo sync"
        ]
    }

/*
  provisioner "mvm-morpheus" {
    url = var.morpheus_url
    username = var.morpheus_username
    password = var.morpheus_password
    task_id = 2
  }
  */
}
