# Examples for Terraform Provider BaseXform

This directory contains examples of how to use the `terraform-provider-basexform`.

## Structure

- `encode/`: Contains examples demonstrating how to encode strings into different bases.
- `decode/`: Contains examples demonstrating how to decode strings from different bases back to their original format.

## Usage

To use these examples, navigate into the respective directory and run the following Terraform commands:

```bash
terraform init
terraform plan
terraform apply
```

## Examples

1. **Encoding Random AWS Account IDs (`encode/` directory):**
   Shows how to generate random strings simulating AWS account IDs and encode them using a high numerical base.

2. **Encode and Decode a Fixed String (`decode/` directory):**
   Demonstrates encoding a fixed string into a specified base and then decoding it back to its original format.
