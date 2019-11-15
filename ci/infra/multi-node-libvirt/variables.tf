variable "image_uri" {
  default     = ""
  description = "URL of the image to use"
}

variable "repositories" {
  type        = "map"
  default     = {}
  description = "Urls of the repositories to mount via cloud-init"
}

variable "stack_name" {
  default     = ""
  description = "Identifier to make all your resources unique and avoid clashes with other users of this terraform project"
}

variable "authorized_keys" {
  type        = "list"
  default     = []
  description = "SSH keys to inject into all the nodes"
}

variable "ntp_servers" {
  type        = "list"
  default     = []
  description = "List of NTP servers to configure"
}

variable "packages" {
  type = "list"

  default = [
    "kernel-default",
    "-kernel-default-base",
  ]

  description = "List of packages to install"
}

variable "username" {
  default     = "sles"
  description = "Username for the cluster nodes"
}

variable "password" {
  default     = "linux"
  description = "Password for the cluster nodes"
}

variable "caasp_registry_code" {
  default     = ""
  description = "SUSE CaaSP Product Registration Code"
}

variable "ha_registry_code" {
  default     = ""
  description = "SUSE Linux Enterprise High Availability Extension Registration Code"
}

variable "rmt_server_name" {
  default     = ""
  description = "SUSE Repository Mirroring Server Name"
}

variable "disk_size" {
  default     = "25769803776"
  description = "Disk size (in bytes)"
}

variable "dns_domain" {
  type        = "string"
  default     = "caasp.local"
  description = "Name of DNS Domain"
}

variable "network_mode" {
  type        = "string"
  default     = "route"
  description = "Network mode used by the cluster"
}

variable "master_memory" {
  default     = 2048
  description = "Amount of RAM for a master"
}

variable "master_vcpu" {
  default     = 2
  description = "Amount of virtual CPUs for a master"
}

variable "worker_memory" {
  default     = 2048
  description = "Amount of RAM for a worker"
}

variable "worker_vcpu" {
  default     = 2
  description = "Amount of virtual CPUs for a worker"
}

variable "region" {
  description = "Name of the region where the nodes are running"
  default     = "lab"
}
