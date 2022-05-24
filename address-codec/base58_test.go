package addresscodec

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEncodeBase58(t *testing.T) {
	tt := []struct {
		description    string
		input          []byte
		expectedOutput string
	}{
		{
			description:    "successful encode with BTC alphabet - 1",
			input:          []byte("bc1q7cd9snht9vgcnxwzrna2zfsgxt9kg8kp655zn4"),
			expectedOutput: "2eEoz8raxrdYefMdpunGH5uhBCmSsygY4pr2eT4Se85WXg15jGsqEfx7rj",
		},
		{
			description:    "successful encode with BTC alphabet - 2",
			input:          []byte("bc1qe6ewupn2f8zreegylpfuxl5vz0j5syema285ma"),
			expectedOutput: "2eEoz8yN8PymBEsFDk6xV3GNs9567GGNhGkTcf56oYJ5VKV9QM2WoeYSn4",
		},
		{
			description:    "successful encode with BTC alphabet - 3",
			input:          []byte("bc1q5ljuq7edlyswpztwt4z3aqcj9mgcw60tz3c5du"),
			expectedOutput: "2eEoz8rJ78w4PmPzQrNfF8RVgz4oCNvj9mkBEPpWACTamm4WBcWnhR5r44",
		},
		// {
		// 	description:    "successful encode with XRP alphabet - 1",
		// 	input:          []byte("rDTXLQ7ZKZVKz33zJbHjgVShjsBnqMBhmN"),
		// 	expectedOutput: "s2Fku4vaPpFiqqXdAD3V5rYrSx5a9h9qvUJW3423akZSCeD",
		// },
		// {
		// 	description:    "successful encode with XRP alphabet - 2",
		// 	input:          []byte("rJrpjzcxwQxokkqPxm62o5rtNfe2XimrTr"),
		// 	expectedOutput: "s2i2Jk6bF44eDSXnnMjxeVhnYZ3qmbteqesuhS6Tz7CSd9j",
		// },
		// {
		// 	description:    "successful encode with XRP alphabet - 3",
		// 	input:          []byte("rUxb5vn9fGYRV3KZcnu3JLM4q5DTnNSavf"),
		// 	expectedOutput: "s2uiNSCBQnQfsVtnX49adC9QqtWNP8upC16t7GFLrmbR7tm",
		// },

	}

	for _, tc := range tt {
		t.Run(tc.description, func(t *testing.T) {
			assert.Equal(t, tc.expectedOutput, EncodeBase58(tc.input))
		})
	}
}

// func TestDecodeBase58(t *testing.T) {
// 	tt := []struct {
// 		description    string
// 		input          string
// 		expectedOutput []byte
// 	}{
// 		// {
// 		// 	description:    "successful decode with BTC alphabet",
// 		// 	input:          "4r2UmqYWb6",
// 		// 	expectedOutput: []byte("example"),
// 		// },
// 		{
// 			description:    "successful decode with XRP alphabet",
// 			input:          "s2Fku4vaPpFiqqXdAD3V5rYrSx5a9h9qvUJW3423akZSCeD",
// 			expectedOutput: []byte("rDTXLQ7ZKZVKz33zJbHjgVShjsBnqMBhmN"),
// 		},
// 	}

// 	for _, tc := range tt {
// 		t.Run(tc.description, func(t *testing.T) {
// 			dec, _ := DecodeBase58(tc.input)
// 			assert.Equal(t, tc.expectedOutput, dec)
// 		})
// 	}
// }
