## Importing Existing Resources

> **Note**: In provider version 2.0.0, imports use UUID-based identification format for NIOS 9.1.0 and WAPI v2.14.

Resources can be imported using their UUID:

```bash
terraform import nios_dns_record_a.example "bf94eb18feee486aa608bb27ca2dcf0a"
```

You can use Terraform's import blocks (available in Terraform 1.5.0 and later) to declaratively import resources:

```hcl
import {
  to = nios_dns_record_a.example
  id = "bf94eb18feee486aa608bb27ca2dcf0a"
}

resource "nios_dns_record_a" "example" {
  # Configuration will be imported from the ID
  # After import, update the configuration as needed
}
```

After running `terraform plan` and `terraform apply`, the resource will be imported and you can then update the configuration as needed.

You can generate a plan config and then use it to import a resource. This is in beta version and is supported in terraform version 1.7.0 or later.
Below is the command to generate the same.

```
terraform plan -generate-config-out=generated.tf
```

Once the config is generated you can execute the `terraform apply` command to import the resource and you can then update the configuration as needed.

Refer the terraform link [here](https://developer.hashicorp.com/terraform/language/v1.14.x/import/generating-configuration).
