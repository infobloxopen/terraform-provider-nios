variable "name" {
    type        = string
    default     = "base-acl-template"
  
}

variable "comment" {
    type        = string
    default     = "Base ACL structure created for future assignment of access control entries"
  
}

variable "extattrs" {
    type        = map(string)
    default     = {
        Site = "location-1"
    }
  
}

variable "name2" {
    type        = string
    default     = "example-network-acl"
  
}

variable "comment2" {
    type        = string
    default     = "ACL to allow/deny access to specific dev network resources"
  
}

variable "extattrs2" {
    type        = map(string)
    default     = {
        Site = "location-2"
    }
  
}

variable "access_list" {
    type = list(object({
        struct        = string
        address       = optional(string)
        permission    = optional(string)
        tsig_key      = optional(string)
        tsig_key_name = optional(string)
        tsig_key_alg  = optional(string)
    }))
    default = null
}

variable "exploded_access_list" {
    type    = any
    default = null
  
}