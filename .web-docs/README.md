This plugin allows Packer to communicate with the Morpheus platform. It is able to create custom images for MVM clusters. This plugin comes with builders designed to support building a custom image from an ISO file and an existing instance.

### Installation

To install this plugin, copy and paste this code into your Packer configuration, then run [`packer init`](https://www.packer.io/docs/commands/init).

```hcl
packer {
  required_plugins {
    mvm = {
      source  = "github.com/martezr/mvm"
      version = ">=0.0.1"
    }
  }
}
```

Alternatively, you can use `packer plugins install` to manage installation of this plugin.

```sh
$ packer plugins install github.com/martezr/mvm
```

### Components

The Morpheus MVM plugin allows Packer to interface with [Morpheus](https://morpheusdata.com).

#### Builders

- [mvm-clone](/docs/builders/mvm-clone.md) - The `mvm-clone` builder is used to create MVM custom templates based on an existing MVM virtual machine.

- [mvm-iso](/docs/builders/mvm-iso.md) - The `mvm-iso` builder is used to create MVM custom templates based on an ISO file.

#### Provisioners

- [mvm-morpheus](/docs/provisioners/mvm-morpheus.md) - The `mvm-morpheus` provisioner that utilizes the Morpheus automation (tasks and workflows) to configure the virtual machine.