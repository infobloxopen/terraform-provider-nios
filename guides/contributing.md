## Contributing

### Terraform RC Configuration for local usage

As the Provider isn't available on registry , to develop the provider locally , you need to set up the `.terraformrc` file in your home directory to point to the `terraform-provider-nios` repository.

```bash
provider_installation {
  dev_overrides {
    "infobloxopen/nios" = "/Users/<user-name>/<path-to-provider>/terraform-provider-nios"
  }
  filesystem_mirror {
    path    = "/Users/<user-name>/.terraform.d/plugins/"
    include = ["infobloxopen/nios"]
  }
  direct {
    exclude = ["infobloxopen/nios"]
  }
}
```
Using this configuration allows Terraform to use the local provider instead of the one from the Terraform registry, which is particularly useful during development and testing.

### How to Contribute a New Feature

1. Fork the repository
2. Create a feature branch (`git checkout -b feature/amazing-feature`)
3. Commit your changes (`git commit -m 'Add some amazing feature'`)
4. Push to the branch (`git push origin feature/amazing-feature`)
5. Open a Pull Request

## Support

For issues, feature requests, or questions, please [open an issue](https://github.com/infobloxopen/terraform-provider-nios/issues) on GitHub.
