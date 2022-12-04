package types

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsPathStep(t *testing.T) {

	tt := []struct {
		description string
		input       map[string]any
		expected    bool
	}{
		{
			description: "represents valid path step",
			input: map[string]any{
				"account":  "rHb9CJAWyB4rj91VRWn96DkukG4bwdtyTh",
				"currency": "USD",
				"issuer":   "rHb9CJAWyB4rj91VRWn96DkukG4bwdtyTh",
			},
			expected: true,
		},
		{
			description: "represents valid path step",
			input: map[string]any{
				"account":  "rHb9CJAWyB4rj91VRWn96DkukG4bwdtyTh",
				"currency": "USD",
			},
			expected: true,
		},
		{
			description: "represents invalid path step",
			input:       map[string]any{},
			expected:    false,
		},
	}

	for _, tc := range tt {
		t.Run(tc.description, func(t *testing.T) {
			assert.Equal(t, tc.expected, isPathStep(tc.input))
		})
	}
}

func TestNewPathStep(t *testing.T) {

	tt := []struct {
		description string
		input       map[string]any
		expected    []byte
	}{
		{
			description: "created valid path step",
			input: map[string]any{
				"account":  "rDTXLQ7ZKZVKz33zJbHjgVShjsBnqMBhmN",
				"currency": "USD",
				"issuer":   "r3Y6vCE8XqfZmYBRngy22uFYkmz3y9eCRA",
			},
			expected: []byte{0x31, 0x88, 0xa5, 0xa5, 0x7c, 0x82, 0x9f, 0x40, 0xf2, 0x5e, 0xa8, 0x33, 0x85, 0xbb, 0xde, 0x6c, 0x3d, 0x8b, 0x4c, 0xa0, 0x82, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x55, 0x53, 0x44, 0x0, 0x0, 0x0, 0x0, 0x0, 0x52, 0xc7, 0xf0, 0x1a, 0xd1, 0x3b, 0x3c, 0xa9, 0xc1, 0xd1, 0x33, 0xfa, 0x8f, 0x34, 0x82, 0xd2, 0xef, 0x8, 0xfa, 0x7d},
		},
	}

	for _, tc := range tt {
		t.Run(tc.description, func(t *testing.T) {
			assert.Equal(t, tc.expected, newPathStep(tc.input))
		})
	}
}

func TestNewPath(t *testing.T) {

	tt := []struct {
		description string
		input       []any
		expected    []byte
	}{
		{
			description: "created valid path",
			input: []any{
				map[string]any{
					"account":  "rDTXLQ7ZKZVKz33zJbHjgVShjsBnqMBhmN",
					"currency": "USD",
					"issuer":   "r3Y6vCE8XqfZmYBRngy22uFYkmz3y9eCRA",
				},
				map[string]any{
					"account":  "rDTXLQ7ZKZVKz33zJbHjgVShjsBnqMBhmN",
					"currency": "USD",
					"issuer":   "r3Y6vCE8XqfZmYBRngy22uFYkmz3y9eCRA",
				},
			},
			expected: []byte{0x31, 0x88, 0xa5, 0xa5, 0x7c, 0x82, 0x9f, 0x40, 0xf2, 0x5e, 0xa8, 0x33, 0x85, 0xbb, 0xde, 0x6c, 0x3d, 0x8b, 0x4c, 0xa0, 0x82, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x55, 0x53, 0x44, 0x0, 0x0, 0x0, 0x0, 0x0, 0x52, 0xc7, 0xf0, 0x1a, 0xd1, 0x3b, 0x3c, 0xa9, 0xc1, 0xd1, 0x33, 0xfa, 0x8f, 0x34, 0x82, 0xd2, 0xef, 0x8, 0xfa, 0x7d, 0x31, 0x88, 0xa5, 0xa5, 0x7c, 0x82, 0x9f, 0x40, 0xf2, 0x5e, 0xa8, 0x33, 0x85, 0xbb, 0xde, 0x6c, 0x3d, 0x8b, 0x4c, 0xa0, 0x82, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x55, 0x53, 0x44, 0x0, 0x0, 0x0, 0x0, 0x0, 0x52, 0xc7, 0xf0, 0x1a, 0xd1, 0x3b, 0x3c, 0xa9, 0xc1, 0xd1, 0x33, 0xfa, 0x8f, 0x34, 0x82, 0xd2, 0xef, 0x8, 0xfa, 0x7d},
		},
	}

	for _, tc := range tt {
		t.Run(tc.description, func(t *testing.T) {
			assert.Equal(t, tc.expected, newPath(tc.input))
		})
	}

}

func TestNewPathSet(t *testing.T) {
	tt := []struct {
		description string
		input       []any
		expected    []byte
	}{
		{
			description: "created valid path set with multiple paths",
			input: []any{
				[]any{
					map[string]any{
						"account":  "rDTXLQ7ZKZVKz33zJbHjgVShjsBnqMBhmN",
						"currency": "USD",
						"issuer":   "r3Y6vCE8XqfZmYBRngy22uFYkmz3y9eCRA",
					},
					map[string]any{
						"account":  "rDTXLQ7ZKZVKz33zJbHjgVShjsBnqMBhmN",
						"currency": "USD",
						"issuer":   "r3Y6vCE8XqfZmYBRngy22uFYkmz3y9eCRA",
					},
				},
				[]any{
					map[string]any{
						"account":  "rDTXLQ7ZKZVKz33zJbHjgVShjsBnqMBhmN",
						"currency": "USD",
						"issuer":   "r3Y6vCE8XqfZmYBRngy22uFYkmz3y9eCRA",
					},
					map[string]any{
						"account":  "rDTXLQ7ZKZVKz33zJbHjgVShjsBnqMBhmN",
						"currency": "USD",
						"issuer":   "r3Y6vCE8XqfZmYBRngy22uFYkmz3y9eCRA",
					},
				},
				[]any{
					map[string]any{
						"account":  "rDTXLQ7ZKZVKz33zJbHjgVShjsBnqMBhmN",
						"currency": "USD",
						"issuer":   "r3Y6vCE8XqfZmYBRngy22uFYkmz3y9eCRA",
					},
					map[string]any{
						"account":  "rDTXLQ7ZKZVKz33zJbHjgVShjsBnqMBhmN",
						"currency": "USD",
						"issuer":   "r3Y6vCE8XqfZmYBRngy22uFYkmz3y9eCRA",
					},
				},
			},
			expected: []byte{0x31, 0x88, 0xa5, 0xa5, 0x7c, 0x82, 0x9f, 0x40, 0xf2, 0x5e, 0xa8, 0x33, 0x85, 0xbb, 0xde, 0x6c, 0x3d, 0x8b, 0x4c, 0xa0, 0x82, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x55, 0x53, 0x44, 0x0, 0x0, 0x0, 0x0, 0x0, 0x52, 0xc7, 0xf0, 0x1a, 0xd1, 0x3b, 0x3c, 0xa9, 0xc1, 0xd1, 0x33, 0xfa, 0x8f, 0x34, 0x82, 0xd2, 0xef, 0x8, 0xfa, 0x7d, 0x31, 0x88, 0xa5, 0xa5, 0x7c, 0x82, 0x9f, 0x40, 0xf2, 0x5e, 0xa8, 0x33, 0x85, 0xbb, 0xde, 0x6c, 0x3d, 0x8b, 0x4c, 0xa0, 0x82, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x55, 0x53, 0x44, 0x0, 0x0, 0x0, 0x0, 0x0, 0x52, 0xc7, 0xf0, 0x1a, 0xd1, 0x3b, 0x3c, 0xa9, 0xc1, 0xd1, 0x33, 0xfa, 0x8f, 0x34, 0x82, 0xd2, 0xef, 0x8, 0xfa, 0x7d, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0xff, 0x31, 0x88, 0xa5, 0xa5, 0x7c, 0x82, 0x9f, 0x40, 0xf2, 0x5e, 0xa8, 0x33, 0x85, 0xbb, 0xde, 0x6c, 0x3d, 0x8b, 0x4c, 0xa0, 0x82, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x55, 0x53, 0x44, 0x0, 0x0, 0x0, 0x0, 0x0, 0x52, 0xc7, 0xf0, 0x1a, 0xd1, 0x3b, 0x3c, 0xa9, 0xc1, 0xd1, 0x33, 0xfa, 0x8f, 0x34, 0x82, 0xd2, 0xef, 0x8, 0xfa, 0x7d, 0x31, 0x88, 0xa5, 0xa5, 0x7c, 0x82, 0x9f, 0x40, 0xf2, 0x5e, 0xa8, 0x33, 0x85, 0xbb, 0xde, 0x6c, 0x3d, 0x8b, 0x4c, 0xa0, 0x82, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x55, 0x53, 0x44, 0x0, 0x0, 0x0, 0x0, 0x0, 0x52, 0xc7, 0xf0, 0x1a, 0xd1, 0x3b, 0x3c, 0xa9, 0xc1, 0xd1, 0x33, 0xfa, 0x8f, 0x34, 0x82, 0xd2, 0xef, 0x8, 0xfa, 0x7d, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0xff, 0x31, 0x88, 0xa5, 0xa5, 0x7c, 0x82, 0x9f, 0x40, 0xf2, 0x5e, 0xa8, 0x33, 0x85, 0xbb, 0xde, 0x6c, 0x3d, 0x8b, 0x4c, 0xa0, 0x82, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x55, 0x53, 0x44, 0x0, 0x0, 0x0, 0x0, 0x0, 0x52, 0xc7, 0xf0, 0x1a, 0xd1, 0x3b, 0x3c, 0xa9, 0xc1, 0xd1, 0x33, 0xfa, 0x8f, 0x34, 0x82, 0xd2, 0xef, 0x8, 0xfa, 0x7d, 0x31, 0x88, 0xa5, 0xa5, 0x7c, 0x82, 0x9f, 0x40, 0xf2, 0x5e, 0xa8, 0x33, 0x85, 0xbb, 0xde, 0x6c, 0x3d, 0x8b, 0x4c, 0xa0, 0x82, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x55, 0x53, 0x44, 0x0, 0x0, 0x0, 0x0, 0x0, 0x52, 0xc7, 0xf0, 0x1a, 0xd1, 0x3b, 0x3c, 0xa9, 0xc1, 0xd1, 0x33, 0xfa, 0x8f, 0x34, 0x82, 0xd2, 0xef, 0x8, 0xfa, 0x7d, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0},
		},
	}

	for _, tc := range tt {
		t.Run(tc.description, func(t *testing.T) {
			assert.Equal(t, tc.expected, newPathSet(tc.input))
		})
	}
}