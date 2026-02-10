// Manage an RIR Organization with Basic Fields
resource "nios_rir_organization" "rir_organization_basic" {
    id = "ID_REPLACE_ME"
    maintainer = "MAINTAINER_REPLACE_ME"
    name = "NAME_REPLACE_ME"
    password = "PASSWORD_REPLACE_ME"
    rir = "RIR_REPLACE_ME"
    sender_email = "SENDER_EMAIL_REPLACE_ME"
}

// Manage an RIR Organization with Additional Fields
resource "nios_rir_organization" "rir_organization_with_additional_fields" {
    id = "ID_REPLACE_ME"
    maintainer = "MAINTAINER_REPLACE_ME"
    name = "NAME_REPLACE_ME"
    password = "PASSWORD_REPLACE_ME"
    rir = "RIR_REPLACE_ME"
    sender_email = "SENDER_EMAIL_REPLACE_ME"

// TODO : Add additional optional fields below

}