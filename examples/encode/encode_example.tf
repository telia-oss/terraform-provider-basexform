
provider "basexform" {

}

terraform {
  required_providers {
    basexform = {
      source  = "github.com/rickardl/basexform"
      version = "0.1"
    }
    random = {
      source  = "hashicorp/random"
      version = "3.1.0"
    }
  }
}

resource "random_string" "random_account_ids" {

  count   = 10
  length  = 12
  upper   = false
  lower   = false
  special = false
}


# Example usage of the basexform_encode resource

resource "basexform_encode" "encoded_accounts" {
  for_each = { for id in toset(random_string.random_account_ids[*].result) : id => id }
  input    = each.value
  base     = 62
}

output "account_info" {
  value = {
    for account_id, encoded in basexform_encode.encoded_accounts : account_id => {
      "original"        = account_id
      "encoded"         = encoded.encoded
      "original_length" = length(account_id)
      "encoded_length"  = length(encoded.encoded)
      "length_diff"     = length(account_id) - length(encoded.encoded)
    }
  }
}
