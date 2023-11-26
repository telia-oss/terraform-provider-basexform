# Resource: basexform_encode

## Description

The `basexform_encode` resource is used to encode a string into a specified numerical base. It is useful for transforming numeric strings into more compact formats.

## Example Usage

```hcl
resource "basexform_encode" "example" {
  input = "123456789012"
  base  = 36
}

output "encoded_value" {
  value = basexform_encode.example.encoded
}
```

## Argument Reference

- `input` - (Required) The string to be encoded.
- `base` - (Required) The numerical base to use for encoding.

## Attributes Reference

- `encoded` - The resulting encoded string.
