## Importing Existing Resources

Resources can be imported using their reference ID:

```bash
terraform import nios_dns_record_a.example record:a/ZG5zLmJpbmRfYSQuX2RlZmF1bHQuY29tLmV4YW1wbGUsc2FtcGxlLDE5Mi4xNjguMS4xMA:example.mydomain.com/default
```

Alternatively, you can use Terraform's import blocks (available in Terraform 1.5.0 and later) to declaratively import resources:

```hcl
import {
  to = nios_dns_record_a.example
  id = "record:a/ZG5zLmJpbmRfYSQuX2RlZmF1bHQuY29tLmV4YW1wbGUsc2FtcGxlLDE5Mi4xNjguMS4xMA:example.mydomain.com/default"
}

resource "nios_dns_record_a" "example" {
  # Configuration will be imported from the ID
  # After import, update the configuration as needed
}
```

After running `terraform plan` and `terraform apply`, the resource will be imported and you can then update the configuration as needed.