# terraform-provider-hdns
[![Go Reference](https://pkg.go.dev/badge/github.com/danielr1996/terraform-provider-hdns.svg)](https://pkg.go.dev/github.com/danielr1996/terraform-provider-hdns)
[![Go Report Card](https://goreportcard.com/badge/github.com/danielr1996/hdns-go)](https://goreportcard.com/report/github.com/danielr1996/hdns-go)
![GitHub tag (latest SemVer)](https://img.shields.io/github/v/tag/danielr1996/hdns-go?style=flat)
## Usage

Add the provider configuration to your `terraform.tf`
```hcl
terraform {
  required_providers {
    hdns = {
      version = "0.0.1"
      source  = "danielr1996/hdns"
    }
  }
}
```

Add a provider block for hdns
```hcl
provider "hdns" {
  token = "<api-token>"
}
```

Add a resource block for a dns record
```hcl
resource "hdns_record" "record" {
  type = "A"
  name = "record.example.com"
  value = "1.2.3.4"
  ttl = 6400

}
```
