provider "basexform" {
}

terraform {
  required_providers {
    basexform = {
      source  = "github.com/rickardl/basexform"
      version = "0.1"
    }
  }
}

# Example usage of the basexform_decode resource

locals {
  value = "123456789012"
}

resource "basexform_encode" "example_encode" {
  input = local.value
  base  = 36
}

resource "basexform_decode" "example_decode" {
  input = basexform_encode.example_encode.encoded
  base  = 36
}

output "decoded_value" {
  value = basexform_decode.example_decode.decoded
}

output "encoded_value" {
  value = basexform_encode.example_encode.encoded
}

output "value" {
  value = local.value

}

output "difference" {
  value = abs(length(local.value) - length(basexform_encode.example_encode.encoded))

}
