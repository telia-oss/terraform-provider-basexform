package internal

import (
	"math/big"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
	"github.com/pkg/errors"
)

func ResourceEncode() *schema.Resource {
	return &schema.Resource{
		Create: resourceEncodeCreate,
		Read:   schema.Noop,
		Delete: schema.RemoveFromState,
		Schema: map[string]*schema.Schema{
			"input": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "The string to be encoded.",
			},
			"base": {
				Type:         schema.TypeInt,
				Required:     true,
				ForceNew:     true,
				Description:  "The numerical base for encoding the string. Valid values range from 2 to 62.",
				ValidateFunc: validation.IntBetween(2, 62),
			},
			"encoded": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The encoded string in the specified base.",
			},
		},
	}
}

func resourceEncodeCreate(d *schema.ResourceData, m interface{}) error {
	input := d.Get("input").(string)
	base := d.Get("base").(int)

	// Validate base range
	if base < 2 || base > 62 {
		return errors.New("Invalid base: must be between 2 and 62")
	}

	// Handle leading zeros
	leadingZeros := len(input) - len(strings.TrimLeft(input, "0"))
	bigInt := new(big.Int)

	// Validate input string
	if _, ok := bigInt.SetString(strings.TrimLeft(input, "0"), 10); !ok {
		return errors.New("Invalid input string: must be a valid base-10 integer")
	}

	// Encode
	encoded := bigInt.Text(base)

	// Store the result
	d.SetId(input)
	d.Set("encoded", strings.Repeat("0", leadingZeros)+encoded)

	return nil
}
