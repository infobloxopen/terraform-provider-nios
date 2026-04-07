// EC2 Instance for Grid Member
resource "aws_instance" "grid" {
  ami                = var.ami_id
  instance_type      = var.instance_type
  key_name           = var.key_name
  availability_zone  = var.availability_zone
  ipv6_address_count = var.enable_ipv6 ? 1 : 0

  subnet_id              = var.mgmt_subnet_id
  vpc_security_group_ids = [var.security_group_id]

  user_data = templatefile("${path.module}/user_data.tftpl", {
    nios_license           = var.nios_license
    remote_console_enabled = var.remote_console_enabled ? "y" : "n"
    default_admin_password = var.default_admin_password
  })

  root_block_device {
    volume_size           = var.volume_size
    volume_type           = var.volume_type
    delete_on_termination = var.delete_on_termination

    tags = merge(
      var.tags,
      {
        Name = var.name_prefix
      }
    )
  }

  tags = merge(
    var.tags,
    {
      Name = var.name_prefix
    }
  )

  iam_instance_profile = var.iam_instance_profile
}

// Eth1 - Grid Communication Interface 
resource "aws_network_interface" "eth1" {
  subnet_id       = var.lan1_subnet_id
  security_groups = [var.security_group_id]

  ipv6_address_count = var.enable_ipv6 ? 1 : 0

  attachment {
    instance     = aws_instance.grid.id
    device_index = 1
  }

  tags = merge(
    var.tags,
    {
      Name = var.name_prefix
    }
  )
}

// Eth2 - HA Interface (Created only when HA is enabled)
resource "aws_network_interface" "eth2" {
  count           = var.ha_enable ? 1 : 0
  subnet_id       = var.lan1_subnet_id
  security_groups = [var.security_group_id]

  ipv6_address_count = var.enable_ipv6 ? 1 : 0

  attachment {
    instance     = aws_instance.grid.id
    device_index = 2
  }

  private_ips_count = var.private_ips_count_eth2
  tags = merge(
    var.tags,
    {
      Name = var.name_prefix
    }
  )
}

data "aws_subnet" "lan1" {
  id = var.lan1_subnet_id
}