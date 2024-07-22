source "null" "basic-example" {
  communicator = "none"
}

build {
  sources = [
    "source.null.basic-example"
  ]

  provisioner "mvm-morpheus" {
    url      = "https://morpheus.test.local"
    username = "admin"
    password = "Password123"
    task_id  = 2
  }
}
