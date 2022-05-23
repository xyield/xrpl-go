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
		// {
		// 	description:    "successful encode with BTC alphabet - long string",
		// 	input:          []byte("A love paragraph is no doubt an important part of a romantic relationship, so if you have been searching for ideas for writing helpful love letters, you are at the right place. If you are looking for suggestions on how to write the perfect long paragraph for your loved one, start your paragraph with something deeply personal. Check out these long paragraphs for her ideas below."),
		// 	expectedOutput: "D3vsZgoihrQxR1NWjMXsH1XVjsagStQ8YR6GHxjzGgnszJfkoJ2SKCVSsB55yWc3RxYGg3apDW84AkQ49nraGakKfprjozaM7d7btwrHkAFdRYjHgjPVgoTuQS6xDSnhAvck49RFNF5ujdSRayt2opDbVoh9MxLBd5NY6pKrbaijhYbrUqNMq9wLCJCLCpYWBKXGgEFRRmFQShg4ELsu2TZS4T4AgiMpc294qe9Vo6YEuUirMmxGwY1RW6WvDWQHTdC2bp1sLnyDZXJwL3vJErBf6N7bm36G2Dpxe1tgsSEPaRFG1a97QSzKBpgQQ2DU2Ndqs6wegYQJQujtQTKjkUm747HGo1SBLobfxKPsmgV9Y3CZU3NVWdQ5oGxyz4xPbvicU28trznQzsKEEvK5oPCxR7WqTaMGweMyy2STBGvE1q4sNtziZXLts5urT8eFS2fkchRtgScQBqHc8RdNrqmQ52ZQtMk7dVs8C4KyY2NX94Sd8xJrjdn6ZjoXhYSnQ3Y88xu",
		// },
		{
			description:    "successful encode with XRP alphabet",
			input:          []byte("rDTXLQ7ZKZVKz33zJbHjgVShjsBnqMBhmN"),
			expectedOutput: "s2Fku4vaPpFiqqXdAD3V5rYrSx5a9h9qvUJW3423akZSCeD",
		},
	}

	for _, tc := range tt {
		t.Run(tc.description, func(t *testing.T) {
			assert.Equal(t, tc.expectedOutput, EncodeBase58(tc.input))
		})
	}
}

func TestEncodeBase58Original(t *testing.T) {
	tt := []struct {
		description    string
		input          []byte
		expectedOutput string
	}{
		// {
		// 	description:    "successful encode with BTC alphabet - long string",
		// 	input:          []byte("A love paragraph is no doubt an important part of a romantic relationship, so if you have been searching for ideas for writing helpful love letters, you are at the right place. If you are looking for suggestions on how to write the perfect long paragraph for your loved one, start your paragraph with something deeply personal. Check out these long paragraphs for her ideas below."),
		// 	expectedOutput: "D3vsZgoihrQxR1NWjMXsH1XVjsagStQ8YR6GHxjzGgnszJfkoJ2SKCVSsB55yWc3RxYGg3apDW84AkQ49nraGakKfprjozaM7d7btwrHkAFdRYjHgjPVgoTuQS6xDSnhAvck49RFNF5ujdSRayt2opDbVoh9MxLBd5NY6pKrbaijhYbrUqNMq9wLCJCLCpYWBKXGgEFRRmFQShg4ELsu2TZS4T4AgiMpc294qe9Vo6YEuUirMmxGwY1RW6WvDWQHTdC2bp1sLnyDZXJwL3vJErBf6N7bm36G2Dpxe1tgsSEPaRFG1a97QSzKBpgQQ2DU2Ndqs6wegYQJQujtQTKjkUm747HGo1SBLobfxKPsmgV9Y3CZU3NVWdQ5oGxyz4xPbvicU28trznQzsKEEvK5oPCxR7WqTaMGweMyy2STBGvE1q4sNtziZXLts5urT8eFS2fkchRtgScQBqHc8RdNrqmQ52ZQtMk7dVs8C4KyY2NX94Sd8xJrjdn6ZjoXhYSnQ3Y88xu",
		// },
		{
			description:    "successful encode with XRP alphabet",
			input:          []byte("rDTXLQ7ZKZVKz33zJbHjgVShjsBnqMBhmN"),
			expectedOutput: "s2Fku4vaPpFiqqXdAD3V5rYrSx5a9h9qvUJW3423akZSCeD",
		},
	}

	for _, tc := range tt {
		t.Run(tc.description, func(t *testing.T) {
			assert.Equal(t, tc.expectedOutput, EncodeBase58Original(tc.input))
		})
	}
}

func TestDecodeBase58(t *testing.T) {
	tt := []struct {
		description    string
		input          string
		expectedOutput []byte
	}{
		// {
		// 	description:    "successful decode with BTC alphabet",
		// 	input:          "D3vsZgoihrQxR1NWjMXsH1XVjsagStQ8YR6GHxjzGgnszJfkoJ2SKCVSsB55yWc3RxYGg3apDW84AkQ49nraGakKfprjozaM7d7btwrHkAFdRYjHgjPVgoTuQS6xDSnhAvck49RFNF5ujdSRayt2opDbVoh9MxLBd5NY6pKrbaijhYbrUqNMq9wLCJCLCpYWBKXGgEFRRmFQShg4ELsu2TZS4T4AgiMpc294qe9Vo6YEuUirMmxGwY1RW6WvDWQHTdC2bp1sLnyDZXJwL3vJErBf6N7bm36G2Dpxe1tgsSEPaRFG1a97QSzKBpgQQ2DU2Ndqs6wegYQJQujtQTKjkUm747HGo1SBLobfxKPsmgV9Y3CZU3NVWdQ5oGxyz4xPbvicU28trznQzsKEEvK5oPCxR7WqTaMGweMyy2STBGvE1q4sNtziZXLts5urT8eFS2fkchRtgScQBqHc8RdNrqmQ52ZQtMk7dVs8C4KyY2NX94Sd8xJrjdn6ZjoXhYSnQ3Y88xu",
		// 	expectedOutput: []byte("A love paragraph is no doubt an important part of a romantic relationship, so if you have been searching for ideas for writing helpful love letters, you are at the right place. If you are looking for suggestions on how to write the perfect long paragraph for your loved one, start your paragraph with something deeply personal. Check out these long paragraphs for her ideas below."),
		// },
		{
			description:    "successful decode with XRP alphabet",
			input:          "s2Fku4vaPpFiqqXdAD3V5rYrSx5a9h9qvUJW3423akZSCeD",
			expectedOutput: []byte("rDTXLQ7ZKZVKz33zJbHjgVShjsBnqMBhmN"),
		},
	}

	for _, tc := range tt {
		t.Run(tc.description, func(t *testing.T) {
			dec, _ := DecodeBase58(tc.input)
			assert.Equal(t, tc.expectedOutput, dec)
		})
	}
}
