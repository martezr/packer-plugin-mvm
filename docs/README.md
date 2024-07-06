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

#### Builders

- [mvm-clone](/docs/builders/exoscale.md) - The `mvm-clone` builder is used to create Exoscale custom templates based on a Compute instance snapshot.

#### Provisioners

- [provisioner](/packer/integrations/hashicorp/scaffolding/latest/components/provisioner/provisioner-name) - The scaffolding provisioner is used to provisioner
  Packer builds.

#### Post-processors

- [post-processor](/packer/integrations/hashicorp/scaffolding/latest/components/post-processor/postprocessor-name) - The scaffolding post-processor is used to
  export scaffolding builds.

#### Data Sources

- [data source](/packer/integrations/hashicorp/scaffolding/latest/components/datasource/datasource-name) - The scaffolding data source is used to
  export scaffolding data.

