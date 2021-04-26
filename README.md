# terraform-provider-hdns
[![Go Reference](https://pkg.go.dev/badge/github.com/danielr1996/terraform-provider-hdns.svg)](https://pkg.go.dev/github.com/danielr1996/terraform-provider-hdns)
[![Go Report Card](https://goreportcard.com/badge/github.com/danielr1996/terraform-provider-hdns)](https://goreportcard.com/report/github.com/danielr1996/terraform-provider-hdns)
![GitHub tag (latest SemVer)](https://img.shields.io/github/v/tag/danielr1996/hdns-go?style=flat)

A terraform provider for the [Hetzner DNS API](https://dns.hetzner.com/api-docs/)

> Inspired by [terraform-provider-hcloud](https://github.com/hetznercloud/terraform-provider-hcloud)

## Usage
Get your API Token at [https://dns.hetzner.com/settings/api-token](https://dns.hetzner.com/settings/api-token)

Add the provider configuration to your `terraform.tf`
```hcl
terraform {
  required_providers {
    hdns = {
      source = "danielr1996/hdns"
      version = "1.0.2"
    }
  }
}
```

Add a provider block for hdns and add your token
```hcl
provider "hdns" {
  token = "<api-token>"
}
```

Add a data block to retrieve the zone id from its name

```hcl
data "hdns_zone" "zone" {
  name = "example.com"
}
```

Add a resource block for a dns record with the `zone_id` from the previous data block
```hcl
resource "hdns_record" "record" {
  zone_id = data.hdns_zone.zone.id
  type = "A"
  name = "record.example.com"
  value = "1.2.3.4"
}
```

### Full example
```hcl
terraform {
  required_providers {
    hdns = {
      source = "danielr1996/hdns"
      version = "1.0.2"
    }
  }
}

provider "hdns" {
  token = "<api-token>"
}

data "hdns_zone" "zone" {
  name = "example.com"
}

resource "hdns_record" "sample" {
  zone_id = data.hdns_zone.zone.id
  type = "A"
  name = "record.example.com"
  value = "1.2.3.4"
}
```
