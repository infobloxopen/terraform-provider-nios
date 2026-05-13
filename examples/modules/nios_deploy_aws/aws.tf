provider "aws" {
  region     = var.aws_region
  access_key = var.aws_access_key
  secret_key = var.aws_secret_key
}

module "node1" {
  source = "github.com/infobloxopen/terraform-provider-nios//modules/nios_deploy_aws?ref=nios_v9.1.0"

  security_group_id = var.security_group_id
  mgmt_subnet_id    = var.mgmt_subnet_id
  lan1_subnet_id    = var.lan1_subnet_id

  ami_id = var.ami_id

  instance_type     = var.instance_type
  key_name          = var.key_name
  availability_zone = var.availability_zone

  volume_size           = var.volume_size
  volume_type           = var.volume_type
  delete_on_termination = var.delete_on_termination

  enable_ipv6 = var.enable_ipv6

  name = var.name
  tags = var.tags

  nios_license           = var.nios_license
  remote_console_enabled = var.remote_console_enabled
  default_admin_password = var.default_admin_password
}
