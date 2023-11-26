# Installation Guide for Terraform Provider BaseXform

## Prerequisites

- Terraform v0.13 or higher.
- Go v1.16 or higher (if building from source).

## Installation Steps

1. **Clone the Repository:**
   ```bash
   git clone https://github.com/rickardl/terraform-provider-basexform.git
   ```

2. **Build the Provider:**
   Navigate to the cloned directory and build the provider binary:
   ```bash
   cd terraform-provider-basexform
   go build
   ```

3. **Install the Provider:**
   Place the compiled binary in the appropriate Terraform plugins directory, typically located at `~/.terraform.d/plugins`.

4. **Use the Provider:**
   Reference the provider in your Terraform configuration as described in the `examples` directory.

For more advanced usage and configuration, refer to the specific documentation of each resource in the `docs/resources` directory.
