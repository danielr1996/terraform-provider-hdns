# terraform-provider-hdns

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
