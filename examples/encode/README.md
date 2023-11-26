# Example: Encoding Random AWS Account IDs

This example demonstrates how to use the `basexform` provider to encode a list of randomly generated AWS account IDs.

## Usage

First, the `random_string` resource generates a set of 10 random 12-digit strings, simulating AWS account IDs. Then, `basexform_encode` resource encodes each of these strings into base 62.

The output displays the original account ID, the encoded value, and the difference in length between them.
