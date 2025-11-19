// Manage IPAM Vlan Range with Basic Fields
resource "nios_ipam_vlanrange" "ipam_vlanrange_basic" {
    end_vlan_id = "END_VLAN_ID_REPLACE_ME"
    name = "NAME_REPLACE_ME"
    start_vlan_id = "START_VLAN_ID_REPLACE_ME"
    vlan_view = "VLAN_VIEW_REPLACE_ME"
}

// Manage IPAM Vlan Range with Additional Fields
resource "nios_ipam_vlanrange" "ipam_vlanrange_with_additional_fields" {
    end_vlan_id = "END_VLAN_ID_REPLACE_ME"
    name = "NAME_REPLACE_ME"
    start_vlan_id = "START_VLAN_ID_REPLACE_ME"
    vlan_view = "VLAN_VIEW_REPLACE_ME"

    // Additional Fields
    comment                 = "Example VLAN View"
    allow_range_overlapping = true

    //Extensible Attributes
    extattrs = {
        Site = "location-1"
    }
}
