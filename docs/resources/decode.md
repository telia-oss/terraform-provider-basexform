# Resource: basexform_decode

## Description

The `basexform_decode` resource decodes a previously encoded string back to its original format. It's complementary to the `basexform_encode` resource.

## Example Usage

```hcl
resource "basexform_decode" "example" {
  input = basexform_encode.example.encoded
  base  = 36
}

output "decoded_value" {
  value = basexform_decode.example.decoded
}
```

## Argument Reference

- `input` - (Required) The encoded string to be decoded.
- `base` - (Required) The numerical base in which the string was encoded.

## Attributes Reference

- `decoded` - The resulting decoded string.
