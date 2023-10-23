package signing

import (
	"testing"

	"github.com/CreatureDev/xrpl-go/model/transactions"
	"github.com/CreatureDev/xrpl-go/model/transactions/types"
	"github.com/CreatureDev/xrpl-go/test"
)

func TestSignRequest(t *testing.T) {
	s := SignRequest{
		TxJson: &transactions.Payment{
			BaseTx: transactions.BaseTx{
				Account:         "rf1BiGeXwwQoi8Z2ueFYTEXSwuJYfV2Jpn",
				TransactionType: transactions.PaymentTx,
			},
			Destination: "ra5nK24KXen9AHvsdFTKHSANinZseWnPcX",
			Amount: types.IssuedCurrencyAmount{
				Issuer:   "rf1BiGeXwwQoi8Z2ueFYTEXSwuJYfV2Jpn",
				Currency: "USD",
				Value:    "1",
			},
		},
		Secret:     "sabc",
		Offline:    false,
		FeeMultMax: 1000,
	}
	j := `{
	"tx_json": {
		"Account": "rf1BiGeXwwQoi8Z2ueFYTEXSwuJYfV2Jpn",
		"TransactionType": "Payment",
		"Amount": {
			"issuer": "rf1BiGeXwwQoi8Z2ueFYTEXSwuJYfV2Jpn",
			"currency": "USD",
			"value": "1"
		},
		"Destination": "ra5nK24KXen9AHvsdFTKHSANinZseWnPcX"
	},
	"secret": "sabc",
	"fee_mult_max": 1000
}`

	if err := test.SerializeAndDeserialize(t, s, j); err != nil {
		t.Error(err)
	}
}

func TestSignResponse(t *testing.T) {
	s := SignResponse{
		TxBlob: "1200142200040000240000000263D5038D7EA4C680000000000000000000000000005553440000000000B5F762798A53D543A014CAF8B297CFF8F2F937E868400000000000753073008114A3780F5CB5A44D366520FC44055E8ED44D9A2270F3E0107321EDDF4ECB8F34A168143B928D48EFE625501FB8552403BBBD3FC038A5788951D7707440C3DCA3FEDE6D785398EEAB10A46B44047FF1B0863FC4313051FB292C991D1E3A9878FABB301128FE4F86F3D8BE4706D53FA97F5536DBD31AF14CD83A5ACDEB068114D96CB910955AB40A0E987EEE82BB3CEDD4441AAAE1F1",
		TxJson: &transactions.TrustSet{
			BaseTx: transactions.BaseTx{
				Account:         "rEuLyBCvcw4CFmzv8RepSiAoNgF8tTGJQC",
				TransactionType: transactions.TrustSetTx,
				Fee:             types.XRPCurrencyAmount(30000),
				Sequence:        2,
				Flags:           types.SetFlag(262144),
				Signers: []transactions.Signer{
					{
						SignerData: transactions.SignerData{
							Account:       "rLFd1FzHMScFhLsXeaxStzv3UC97QHGAbM",
							SigningPubKey: "EDDF4ECB8F34A168143B928D48EFE625501FB8552403BBBD3FC038A5788951D770",
							TxnSignature:  "C3DCA3FEDE6D785398EEAB10A46B44047FF1B0863FC4313051FB292C991D1E3A9878FABB301128FE4F86F3D8BE4706D53FA97F5536DBD31AF14CD83A5ACDEB06",
						},
					},
				},
			},
			LimitAmount: types.IssuedCurrencyAmount{
				Issuer:   "rHb9CJAWyB4rj91VRWn96DkukG4bwdtyTh",
				Currency: "USD",
				Value:    "100",
			},
		},
	}

	j := `{
	"tx_blob": "1200142200040000240000000263D5038D7EA4C680000000000000000000000000005553440000000000B5F762798A53D543A014CAF8B297CFF8F2F937E868400000000000753073008114A3780F5CB5A44D366520FC44055E8ED44D9A2270F3E0107321EDDF4ECB8F34A168143B928D48EFE625501FB8552403BBBD3FC038A5788951D7707440C3DCA3FEDE6D785398EEAB10A46B44047FF1B0863FC4313051FB292C991D1E3A9878FABB301128FE4F86F3D8BE4706D53FA97F5536DBD31AF14CD83A5ACDEB068114D96CB910955AB40A0E987EEE82BB3CEDD4441AAAE1F1",
	"tx_json": {
		"Account": "rEuLyBCvcw4CFmzv8RepSiAoNgF8tTGJQC",
		"TransactionType": "TrustSet",
		"Fee": "30000",
		"Sequence": 2,
		"Flags": 262144,
		"Signers": [
			{
				"Signer": {
					"Account": "rLFd1FzHMScFhLsXeaxStzv3UC97QHGAbM",
					"TxnSignature": "C3DCA3FEDE6D785398EEAB10A46B44047FF1B0863FC4313051FB292C991D1E3A9878FABB301128FE4F86F3D8BE4706D53FA97F5536DBD31AF14CD83A5ACDEB06",
					"SigningPubKey": "EDDF4ECB8F34A168143B928D48EFE625501FB8552403BBBD3FC038A5788951D770"
				}
			}
		],
		"LimitAmount": {
			"issuer": "rHb9CJAWyB4rj91VRWn96DkukG4bwdtyTh",
			"currency": "USD",
			"value": "100"
		}
	}
}`

	if err := test.SerializeAndDeserialize(t, s, j); err != nil {
		t.Error(err)
	}
}
