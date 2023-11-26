package internal

import (
	"math/big"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
	"github.com/pkg/errors"
)

func ResourceDecode() *schema.Resource {
	return &schema.Resource{
		Create: resourceDecodeCreate,
		Read:   schema.Noop,
		Delete: schema.RemoveFromState,
		Schema: map[string]*schema.Schema{
			"input": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "The string to be decoded, which is encoded in the specified base.",
			},
			"base": {
				Type:        schema.TypeInt,
				Required:    true,
				ForceNew:    true,
				Description: "The numerical base of the encoded string. Valid values range from 2 to 62.",

				ValidateFunc: validation.IntBetween(2, 62),
			},
			"decoded": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The decoded string, converted back to its original format.",
			},
		},
	}
}

func resourceDecodeCreate(d *schema.ResourceData, m interface{}) error {
	input := d.Get("input").(string)
	base := d.Get("base").(int)

	// Validate base range
	if base < 2 || base > 62 {
		return errors.New("Invalid base: must be between 2 and 62")
	}

	// Handle leading zeros
	leadingZeros := len(input) - len(strings.TrimLeft(input, "0"))
	bigInt := new(big.Int)

	// Validate and set the input string
	if _, ok := bigInt.SetString(strings.TrimLeft(input, "0"), base); !ok {
		return errors.New("Invalid input string: must be a valid number in the specified base")
	}

	// Decode
	decoded := bigInt.Text(10)

	// Store the result
	d.SetId(input)
	d.Set("decoded", strings.Repeat("0", leadingZeros)+decoded)

	return nil
}
