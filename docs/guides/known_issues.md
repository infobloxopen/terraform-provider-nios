##  Terraform Known Issues

- Password attribute values for all objects should not be stored in plaintext format within the terraform.tfstate file.
- FQDN validation needs to be added for multiple objects.
- Admin Group:
    - password_setting cannot be updated when use_password_setting is unset under Admin Group.
    - Not able to set disable_all and enable_all under admin_set_commands in Admin Group.
    - Unable to Set/Unset many flags.
- Network:
    - Unable to unset cloud info in terraform.
- Host Record:
    - Unable to unset cloud info in terraform.
    - Unable to add authentication_password and privacy_password under snmp3_credential in the host record.
    - Cleanup of ipv4addrs and ipv6addrs sub-fields in doc as these are not accessible.
- Certificate Authentication Service:
    - File Uploading for objects like CA Auth Service take place in Update Calls even if file path is not updated.
- VDiscovery Task:
    - Need handle proper error message while creating openstack_vdiscoverytask with wrong HTTPS port.
    - Cannot perform search via service_account_file_token and cdiscovery_file_token although wapi mentions them to be searchable fields.
- CAA Record:
    - EAs added via backend cannot be erased when Config does NOT have EA block.
- Alias Record:
    - Cannot create Alias Record of type CAA through API but operation is supported on UI.

