# Terraform Provider BaseXform

## Overview

`terraform-provider-basexform` is a Terraform provider that allows for encoding and decoding strings in arbitrary numerical bases, with special handling for leading zeros. This provider is particularly useful for transforming numeric strings to more compact formats in various base systems.

## Features

- **Encoding**: Convert strings to a specified base.
- **Decoding**: Convert strings back from a specified base to their original format.
- **Leading Zero Preservation**: Maintains leading zeros in the original string.

## Usage

To use this provider, add the following to your Terraform configuration:

```hcl
resource "basexform_encode" "example" {
  input = "123456789012"
  base  = 36 # Base can be any integer from 2 to 62
}

output "encoded_value" {
  value = basexform_encode.example.encoded
}
```

## Recommendations

The choice of the base in encoding significantly impacts the length of the resulting encoded string. A higher base leads to a more compact representation of the original string, as it uses a wider range of characters to represent the same amount of data. Here's an expanded table including expected length reductions and typical uses:

| Base | Character Set                                                | Expected Length Reduction | Typical Use                                  |
|------|--------------------------------------------------------------|---------------------------|----------------------------------------------|
| 2    | 0, 1                                                         | Very Low                  | Binary data, low-level computing             |
| 8    | 0-7                                                          | Low                       | Octal number system in computing             |
| 10   | 0-9                                                          | None                      | Standard decimal system                      |
| 16   | 0-9, A-F                                                     | Moderate                  | Hexadecimal, widely used in computing        |
| 32   | A-Z, 2-7                                                     | High                      | Data encoding like Base32                    |
| 36   | 0-9, A-Z                                                     | Higher                    | Compact alphanumeric data without casing     |
| 62   | 0-9, A-Z, a-z                                                | Highest                   | Highly compact data, full case sensitivity   |

**Note**: The 'Expected Length Reduction' is relative to the original string's length. Higher bases generally offer more significant compression, making them suitable for scenarios where space efficiency is important. It's important to note, however, that the extent of reduction depends on the input string's nature. For instance, a string with many repeated characters might not see as significant a reduction

### Examples

Example usage of the `basexform_encode` and `basexform_decode` resources:

```hcl
resource "basexform_encode" "example" {
  input = "123456789012"
  base  = 36
}

output "encoded_value" {
  value = basexform_encode.example.encoded
}

resource "basexform_decode" "example" {
  input = basexform_encode.example.encoded
  base  = 36
}

output "decoded_value" {
  value = basexform_decode.example.decoded
}
```

See more examples in the [examples](examples) directory.

### Installation

Clone the repository and build the provider:

```bash
git clone https://github.com/rickardl/terraform-provider-basexform.git
cd terraform-provider-basexform
go build
```

### Building the Provider

To build the provider, run `go build` in the root directory.

### Testing

Run `go test` in the root directory to execute the provider's tests.

### Installing the Provider

#### Using Your Locally Built Terraform Provider BaseXform

**Option 1: Default Plugins Directory**
Place the provider binary in `~/.terraform.d/plugins` (macOS/Linux) or `C:\Users\[username]\.terraform.d\plugins` (Windows).

```bash
chmod +x terraform-provider-basexform
# Depending on your OS, the following command may differ
cp terraform-provider-basexform ~/.terraform.d/plugins/github.com/rickardl/basexform/0.1.0/darwin_amd64/terraform-provider-basexform
```

**Option 2: Custom Plugins Directory**
Set the `TERRAFORM_PLUGIN_DIRS` environment variable to your custom directory:

```bash
export TERRAFORM_PLUGIN_DIRS="/path/to/plugins"
```

## Detailed Configuration Options

- **`input`**: A string representing the value you wish to encode or decode. For encoding, this should be a numeric string (e.g., "12345"). It's essential to ensure the string is valid in base10 format if it's going to be encoded.
- **`base`**: An integer defining the numerical base to which the string will be converted. The valid range for this parameter is 2 to 62, encompassing binary (base 2), octal (base 8), decimal (base 10), hexadecimal (base 16), and other alphanumeric systems (up to base 62). The choice of base depends on your specific encoding needs.

## Contributing

Contributions to this provider are welcome! Please submit pull requests or issues to the GitHub repository.

## License

Distributed under the [MIT License](LICENSE).
