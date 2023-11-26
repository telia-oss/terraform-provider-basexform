package internal_test

import (
	"testing"

	"github.com/rickardl/terraform-provider-basexform/internal"

	"github.com/stretchr/testify/assert"
)

func TestResourceDecode(t *testing.T) {
	testCases := []struct {
		name     string
		input    string
		base     int
		expected string
	}{
		{
			name:     "Test case 1",
			input:    "1kpqzg2c",
			base:     36,
			expected: "123456789012",
		},
		// Add more test cases here
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			resource := internal.ResourceDecode()
			d := resource.TestResourceData()
			d.Set("input", tc.input)
			d.Set("base", tc.base)

			err := resource.Create(d, nil)
			assert.NoError(t, err)
			assert.Equal(t, tc.expected, d.Get("decoded"))
		})
	}
}
