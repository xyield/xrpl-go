package addresscodec

import (
	"encoding/hex"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEncodeAddress(t *testing.T) {
	tt := []struct {
		description string
		input       string
		output      string
	}{
		{
			description: "test 1",
			input:       "ED9434799226374926EDA3B54B1B461B4ABF7237962EAE18528FEA67595397FA32",
			output:      "rDTXLQ7ZKZVKz33zJbHjgVShjsBnqMBhmN",
		},
	}

	for _, tc := range tt {
		t.Run(tc.description, func(t *testing.T) {

			got, _ := EncodeAddress(tc.input)
			assert.Equal(t, tc.output, got)
			assert.Equal(t, pubkeyHex, strings.ToUpper(hex.EncodeToString(pubkey)))
			assert.Equal(t, ED25519PrefixHex, strings.ToUpper(hex.EncodeToString(ED25519Prefix)))
		})
	}
}
