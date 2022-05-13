package definitions

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetTypeNameByFieldName(t *testing.T) {

	tt := []struct {
		description   string
		input         string
		expected      string
		expectedError error
	}{
		{
			description:   "test that `TransferRate` gives `UInt32`",
			input:         "TransferRate",
			expected:      "UInt32",
			expectedError: nil,
		},
		{
			description: "test that invalid value gives an error",
			input:       "yurt",
			expected:    "",
			expectedError: &NotFoundError{
				Instance: "FieldName",
				Input:    "yurt",
			},
		},
	}

	for _, test := range tt {

		t.Run(test.description, func(t *testing.T) {
			got, err := definitions.GetTypeNameByFieldName(test.input)
			if test.expectedError != nil {
				assert.Error(t, err, test.expectedError.Error())
				assert.Zero(t, got)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, test.expected, got)
			}
		})

	}

}

func TestGetTypeCodeByTypeName(t *testing.T) {

	tt := []struct {
		description   string
		input         string
		expected      int
		expectedError error
	}{
		{
			description:   "test that `Done` gives correct code",
			input:         "Done",
			expected:      -1,
			expectedError: nil,
		},
		{
			description:   "test that `Hash128` gives correct code",
			input:         "Hash128",
			expected:      4,
			expectedError: nil,
		},
		{
			description: "test that incorrect value gives an error",
			input:       "yurt",
			expected:    0,
			expectedError: &NotFoundError{
				Instance: "TypeName",
				Input:    "yurt",
			},
		},
	}

	for _, test := range tt {

		t.Run(test.description, func(t *testing.T) {
			got, err := definitions.GetTypeCodeByTypeName(test.input)
			if test.expectedError != nil {
				assert.Error(t, err, test.expectedError.Error())
				assert.Zero(t, got)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, test.expected, got)
			}
		})

	}

}

func TestGetTypeCodeByFieldName(t *testing.T) {
	tt := []struct {
		description   string
		input         string
		expected      int
		expectedError error
	}{
		{
			description:   "test that `TransferRate` gives 2",
			input:         "TransferRate",
			expected:      2,
			expectedError: nil,
		},
		{
			description:   "test that `OwnerNode` gives 3",
			input:         "OwnerNode",
			expected:      3,
			expectedError: nil,
		},
		{
			description: "test that non-existent value gives error",
			input:       "yurt",
			expected:    0,
			expectedError: &NotFoundError{
				Instance: "FieldName",
				Input:    "yurt",
			},
		},
	}

	for _, test := range tt {

		t.Run(test.description, func(t *testing.T) {
			got, err := definitions.GetTypeCodeByFieldName(test.input)
			if test.expectedError != nil {
				assert.Error(t, err, test.expectedError.Error())
				assert.Zero(t, got)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, test.expected, got)
			}
		})
	}
}

func TestGetFieldCodeByFieldName(t *testing.T) {

	tt := []struct {
		description   string
		input         string
		expected      int
		expectedError error
	}{
		{
			description:   "correct FieldCode",
			input:         "TransferRate",
			expected:      11,
			expectedError: nil,
		},
		{
			description: "Invalid FieldName",
			input:       "yurt",
			expected:    0,
			expectedError: &NotFoundError{
				Instance: "FieldName",
				Input:    "yurt",
			},
		},
	}

	for _, test := range tt {

		t.Run(test.description, func(t *testing.T) {
			got, err := definitions.GetFieldCodeByFieldName(test.input)
			if test.expectedError != nil {
				assert.Error(t, err, test.expectedError.Error())
				assert.Zero(t, got)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, test.expected, got)
			}
		})
	}
}

func TestGetFieldHeaderByFieldName(t *testing.T) {
	tt := []struct {
		description   string
		input         string
		expected      *fieldHeader
		expectedError error
	}{
		{
			description: "correct FieldHeader",
			input:       "TransferRate",
			expected: &fieldHeader{
				TypeCode:  2,
				FieldCode: 11,
			},
			expectedError: nil,
		},
		{
			description: "Invalid FieldName",
			input:       "yurt",
			expected:    nil,
			expectedError: &NotFoundError{
				Instance: "FieldName",
				Input:    "yurt",
			},
		},
	}

	for _, test := range tt {

		t.Run(test.description, func(t *testing.T) {
			got, err := definitions.GetFieldHeaderByFieldName(test.input)
			if test.expectedError != nil {
				assert.Error(t, err, test.expectedError.Error())
				assert.Nil(t, got)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, test.expected, got)
			}
		})
	}
}

func TestGetFieldNameByFieldHeader(t *testing.T) {
	tt := []struct {
		description   string
		input         fieldHeader
		expected      string
		expectedError error
	}{
		{
			description: "correct fieldName",
			input: fieldHeader{
				TypeCode:  1,
				FieldCode: 1,
			},
			expected:      "LedgerEntryType",
			expectedError: nil,
		},
		{
			description: "correct fieldName 2",
			input: fieldHeader{
				TypeCode:  5,
				FieldCode: 21,
			},
			expected:      "Digest",
			expectedError: nil,
		},
		{
			description: "invalid fieldHeader",
			input: fieldHeader{
				TypeCode:  0000000000000111,
				FieldCode: 000000000000111,
			},
			expected: "",
			expectedError: &NotFoundErrorFieldHeader{
				Instance: "FieldHeader",
				Input: fieldHeader{
					TypeCode:  0000000000000111,
					FieldCode: 000000000000111,
				},
			},
		},
	}

	for _, test := range tt {
		t.Run(test.description, func(t *testing.T) {
			got, err := definitions.GetFieldNameByFieldHeader(test.input)
			if test.expectedError != nil {
				assert.Error(t, err, test.expectedError.Error())
				assert.Zero(t, got)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, test.expected, got)
			}
		})
	}
}

func TestGetFieldInfoByFieldName(t *testing.T) {
	tt := []struct {
		description   string
		input         string
		expected      *fieldInfo
		expectedError error
	}{
		{
			description: "correct FieldInfo",
			input:       "TransferRate",
			expected: &fieldInfo{
				Nth:            11,
				IsVLEncoded:    false,
				IsSerialized:   true,
				IsSigningField: true,
				Type:           "UInt32",
			},
			expectedError: nil,
		},
		{
			description: "invalid FieldInfo",
			input:       "yurt",
			expected:    nil,
			expectedError: &NotFoundError{
				Instance: "FieldName",
				Input:    "yurt",
			},
		},
	}

	for _, test := range tt {
		t.Run(test.description, func(t *testing.T) {
			got, err := definitions.GetFieldInfoByFieldName(test.input)
			if test.expectedError != nil {
				assert.Error(t, err, test.expectedError.Error())
				assert.Nil(t, got)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, test.expected, got)
			}

		})
	}
}

func TestGetFieldInstanceByFieldName(t *testing.T) {
	tt := []struct {
		description   string
		input         string
		expected      *fieldInstance
		expectedError error
	}{
		{
			description: "correct FieldInstance",
			input:       "TransferRate",
			expected: &fieldInstance{
				FieldName: "TransferRate",
				fieldInfo: &fieldInfo{
					Nth:            11,
					IsVLEncoded:    false,
					IsSerialized:   true,
					IsSigningField: true,
					Type:           "UInt32",
				},
				FieldHeader: &fieldHeader{
					TypeCode:  2,
					FieldCode: 11,
				},
			},
		},
		{
			description: "invalid FieldName",
			input:       "yurt",
			expected:    nil,
			expectedError: &NotFoundError{
				Instance: "FieldName",
				Input:    "yurt",
			},
		},
	}

	for _, test := range tt {

		t.Run(test.description, func(t *testing.T) {
			got, err := definitions.GetFieldInstanceByFieldName(test.input)
			if test.expectedError != nil {
				assert.Error(t, err, test.expectedError.Error())
				assert.Nil(t, got)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, test.expected, got)
			}
		})
	}
}

func TestGetTransactionTypeCodeByTransactionTypeName(t *testing.T) {
	tt := []struct {
		description   string
		input         string
		expected      int
		expectedError error
	}{
		{
			description:   "correct TransactionTypeCode",
			input:         "EscrowCreate",
			expected:      1,
			expectedError: nil,
		},
		{
			description: "invalid TransactionTypeName",
			input:       "yurt",
			expected:    0,
			expectedError: &NotFoundError{
				Instance: "TransactionTypeName",
				Input:    "yurt",
			},
		},
	}

	for _, test := range tt {
		t.Run(test.description, func(t *testing.T) {
			got, err := definitions.GetTransactionTypeCodeByTransactionTypeName(test.input)
			if test.expectedError != nil {
				assert.Error(t, err, test.expectedError.Error())
				assert.Zero(t, got)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, test.expected, got)
			}
		})
	}
}

func TestGetTransactionTypeNameByTransactionTypeCode(t *testing.T) {
	tt := []struct {
		description   string
		input         int
		expected      string
		expectedError error
	}{
		{
			description:   "correct TypeName",
			input:         1,
			expected:      "EscrowCreate",
			expectedError: nil,
		},
		{
			description: "invalid TransactionTypeCode",
			input:       999999999999,
			expected:    "",
			expectedError: &NotFoundErrorInt{
				Instance: "TransactionTypeCode",
				Input:    999999999999,
			},
		},
	}

	for _, test := range tt {
		t.Run(test.description, func(t *testing.T) {
			got, err := definitions.GetTransactionTypeNameByTransactionTypeCode(test.input)
			if test.expectedError != nil {
				assert.Error(t, err, test.expectedError.Error())
				assert.Zero(t, got)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, test.expected, got)
			}
		})
	}
}

func TestGetTransactionResultNameByTransactionResultTypeCode(t *testing.T) {
	tt := []struct {
		description   string
		input         int
		expected      string
		expectedError error
	}{
		{
			description:   "correct TransactionResultName",
			input:         100,
			expected:      "tecCLAIM",
			expectedError: nil,
		},
		{
			description: "invalid txResultTypeCode",
			input:       999999999999999,
			expected:    "",
			expectedError: &NotFoundErrorInt{
				Instance: "TransactionResultTypeCode",
				Input:    999999999999999,
			},
		},
	}

	for _, test := range tt {
		t.Run(test.description, func(t *testing.T) {
			got, err := definitions.GetTransactionResultNameByTransactionResultTypeCode(test.input)
			if test.expectedError != nil {
				assert.Error(t, err, test.expectedError.Error())
				assert.Zero(t, got)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, test.expected, got)
			}
		})
	}
}

func TestGetTransactionResultTypeCodeByTransactionResultName(t *testing.T) {
	tt := []struct {
		description   string
		input         string
		expected      int
		expectedError error
	}{
		{
			description:   "correct TransactionResultTypeCode",
			input:         "tecCLAIM",
			expected:      100,
			expectedError: nil,
		},
		{
			description: "invalid TransactionResultName",
			input:       "yurt",
			expected:    0,
			expectedError: &NotFoundError{
				Instance: "TransactionResultName",
				Input:    "yurt",
			},
		},
	}

	for _, test := range tt {
		t.Run(test.description, func(t *testing.T) {
			got, err := definitions.GetTransactionResultTypeCodeByTransactionResultName(test.input)
			if test.expectedError != nil {
				assert.Error(t, err, test.expectedError.Error())
				assert.Zero(t, got)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, test.expected, got)
			}
		})
	}
}

func TestGetLedgerEntryTypeCodeByLedgerEntryTypeName(t *testing.T) {
	tt := []struct {
		description   string
		input         string
		expected      int
		expectedError error
	}{
		{
			description:   "correct LedgerEntryTypeCode",
			input:         "Any",
			expected:      -3,
			expectedError: nil,
		},
		{
			description: "invalid LedgerEntryTypeName",
			input:       "yurt",
			expected:    0,
			expectedError: &NotFoundError{
				Instance: "LedgerEntryTypeName",
				Input:    "yurt",
			},
		},
	}

	for _, test := range tt {
		t.Run(test.description, func(t *testing.T) {
			got, err := definitions.GetLedgerEntryTypeCodeByLedgerEntryTypeName(test.input)
			if test.expectedError != nil {
				assert.Error(t, err, test.expectedError.Error())
				assert.Zero(t, got)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, test.expected, got)
			}
		})
	}

}

func TestGetLedgerEntryTypeNameByLedgerEntryTypeCode(t *testing.T) {
	tt := []struct {
		description   string
		input         int
		expected      string
		expectedError error
	}{
		{
			description:   "correct LedgerEntryTypeName",
			input:         -3,
			expected:      "Any",
			expectedError: nil,
		},
		{
			description: "invalid LedgerEntryTypeCode",
			input:       999999999999,
			expected:    "",
			expectedError: &NotFoundErrorInt{
				Instance: "LedgerEntryTypeCode",
				Input:    999999999999,
			},
		},
	}

	for _, test := range tt {
		t.Run(test.description, func(t *testing.T) {
			got, err := definitions.GetLedgerEntryTypeNameByLedgerEntryTypeCode(test.input)
			if test.expectedError != nil {
				assert.Error(t, err, test.expectedError.Error())
				assert.Zero(t, got)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, test.expected, got)
			}
		})
	}
}

func TestSortMapByValue(t *testing.T) {
	tt := []struct {
		description    string
		input          map[string]int
		expectedKeys   []string
		expectedValues []int
	}{
		{
			description: "LedgerEntryTypes sorted correctly",
			input:       definitions.LedgerEntryTypes,
			expectedKeys: []string{
				"Any",
				"Child",
				"Invalid",
				"NFTokenOffer",
				"Check",
				"NegativeUNL",
				"NFTokenPage",
				"SignerList",
				"Ticket",
				"AccountRoot",
				"Contract",
				"DirectoryNode",
				"Amendments",
				"LedgerHashes",
				"Nickname",
				"Offer",
				"DepositPreauth",
				"RippleState",
				"FeeSettings",
				"Escrow",
				"PayChannel",
			},
			expectedValues: []int{-3, -2, -1, 55, 67, 78, 80, 83, 84, 97, 99, 100, 102, 104, 110, 111, 112, 114, 115, 117, 120},
		},
		{
			description: "Types sorted correctly",
			input:       definitions.Types,
			expectedKeys: []string{
				"Unknown",
				"Done",
				"NotPresent",
				"UInt16",
				"UInt32",
				"UInt64",
				"Hash128",
				"Hash256",
				"Amount",
				"Blob",
				"AccountID",
				"STObject",
				"STArray",
				"UInt8",
				"Hash160",
				"PathSet",
				"Vector256",
				"Transaction",
				"LedgerEntry",
				"Validation",
			},
			expectedValues: []int{-2, -1, 0, 1, 2, 3, 4, 5, 6, 7, 8, 14, 15, 16, 17, 18, 19, 10001, 10002, 10003},
		},
		{
			description: "TransactionTypes sorted correctly",
			input:       definitions.TransactionTypes,
			expectedKeys: []string{
				"Invalid",
				"Payment",
				"EscrowCreate",
				"EscrowFinish",
				"AccountSet",
				"EscrowCancel",
				"SetRegularKey",
				"NickNameSet",
				"OfferCreate",
				"OfferCancel",
				"Contract",
				"TicketCreate",
				"TicketCancel",
				"SignerListSet",
				"PaymentChannelCreate",
				"PaymentChannelFund",
				"PaymentChannelClaim",
				"CheckCreate",
				"CheckCash",
				"CheckCancel",
				"DepositPreauth",
				"TrustSet",
				"AccountDelete",
				"NFTokenMint",
				"NFTokenBurn",
				"NFTokenCreateOffer",
				"NFTokenCancelOffer",
				"NFTokenAcceptOffer",
				"EnableAmendment",
				"SetFee",
				"UNLModify",
			},
			expectedValues: []int{-1, 0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 25, 26, 27, 28, 29, 100, 101, 102},
		},
		{
			description: "TransactionResults sorted correctly",
			input:       definitions.TransactionResults,
			expectedKeys: []string{
				"telLOCAL_ERROR",
				"telBAD_DOMAIN",
				"telBAD_PATH_COUNT",
				"telBAD_PUBLIC_KEY",
				"telFAILED_PROCESSING",
				"telINSUF_FEE_P",
				"telNO_DST_PARTIAL",
				"telCAN_NOT_QUEUE",
				"telCAN_NOT_QUEUE_BALANCE",
				"telCAN_NOT_QUEUE_BLOCKS",
				"telCAN_NOT_QUEUE_BLOCKED",
				"telCAN_NOT_QUEUE_FEE",
				"telCAN_NOT_QUEUE_FULL",
				"temMALFORMED",
				"temBAD_AMOUNT",
				"temBAD_CURRENCY",
				"temBAD_EXPIRATION",
				"temBAD_FEE",
				"temBAD_ISSUER",
				"temBAD_LIMIT",
				"temBAD_OFFER",
				"temBAD_PATH",
				"temBAD_PATH_LOOP",
				"temBAD_REGKEY",
				"temBAD_SEND_XRP_LIMIT",
				"temBAD_SEND_XRP_MAX",
				"temBAD_SEND_XRP_NO_DIRECT",
				"temBAD_SEND_XRP_PARTIAL",
				"temBAD_SEND_XRP_PATHS",
				"temBAD_SEQUENCE",
				"temBAD_SIGNATURE",
				"temBAD_SRC_ACCOUNT",
				"temBAD_TRANSFER_RATE",
				"temDST_IS_SRC",
				"temDST_NEEDED",
				"temINVALID",
				"temINVALID_FLAG",
				"temREDUNDANT",
				"temRIPPLE_EMPTY",
				"temDISABLED",
				"temBAD_SIGNER",
				"temBAD_QUORUM",
				"temBAD_WEIGHT",
				"temBAD_TICK_SIZE",
				"temINVALID_ACCOUNT_ID",
				"temCANNOT_PREAUTH_SELF",
				"temUNCERTAIN",
				"temUNKNOWN",
				"temSEQ_AND_TICKET",
				"temBAD_NFTOKEN_TRANSFER_FEE",
				"tefFAILURE", "tefALREADY",
				"tefBAD_ADD_AUTH",
				"tefBAD_AUTH",
				"tefBAD_LEDGER",
				"tefCREATED",
				"tefEXCEPTION",
				"tefINTERNAL",
				"tefNO_AUTH_REQUIRED",
				"tefPAST_SEQ",
				"tefWRONG_PRIOR",
				"tefMASTER_DISABLED",
				"tefMAX_LEDGER",
				"tefBAD_SIGNATURE",
				"tefBAD_QUORUM",
				"tefNOT_MULTI_SIGNING",
				"tefBAD_AUTH_MASTER",
				"tefINVARIANT_FAILED",
				"tefTOO_BIG",
				"tefNO_TICKET",
				"tefNFTOKEN_IS_NOT_TRANSFERABLE",
				"terRETRY",
				"terFUNDS_SPENT",
				"terINSUF_FEE_B",
				"terNO_ACCOUNT",
				"terNO_AUTH",
				"terNO_LINE",
				"terOWNERS",
				"terPRE_SEQ",
				"terLAST",
				"terNO_RIPPLE",
				"terQUEUED",
				"terPRE_TICKET",
				"tesSUCCESS",
				"tecCLAIM",
				"tecPATH_PARTIAL",
				"tecUNFUNDED_ADD",
				"tecUNFUNDED_OFFER",
				"tecUNFUNDED_PAYMENT",
				"tecFAILED_PROCESSING",
				"tecDIR_FULL",
				"tecINSUF_RESERVE_LINE",
				"tecINSUF_RESERVE_OFFER",
				"tecNO_DST", "tecNO_DST_INSUF_XRP",
				"tecNO_LINE_INSUF_RESERVE",
				"tecNO_LINE_REDUNDANT",
				"tecPATH_DRY",
				"tecUNFUNDED",
				"tecNO_ALTERNATIVE_KEY",
				"tecNO_REGULAR_KEY",
				"tecOWNERS",
				"tecNO_ISSUER",
				"tecNO_AUTH",
				"tecNO_LINE",
				"tecINSUFF_FEE",
				"tecFROZEN",
				"tecNO_TARGET",
				"tecNO_PERMISSION",
				"tecNO_ENTRY",
				"tecINSUFFICIENT_RESERVE",
				"tecNEED_MASTER_KEY", "tecDST_TAG_NEEDED",
				"tecINTERNAL",
				"tecOVERSIZE",
				"tecCRYPTOCONDITION_ERROR",
				"tecINVARIANT_FAILED",
				"tecEXPIRED",
				"tecDUPLICATE",
				"tecKILLED",
				"tecHAS_OBLIGATIONS",
				"tecTOO_SOON",
				"tecMAX_SEQUENCE_REACHED",
				"tecNO_SUITABLE_NFTOKEN_PAGE",
				"tecNFTOKEN_BUY_SELL_MISMATCH",
				"tecNFTOKEN_OFFER_TYPE_MISMATCH",
				"tecCANT_ACCEPT_OWN_NFTOKEN_OFFER",
				"tecINSUFFICIENT_FUNDS",
				"tecOBJECT_NOT_FOUND",
				"tecINSUFFICIENT_PAYMENT",
				"tecINCORRECT_ASSET",
				"tecTOO_MANY",
			},
			expectedValues: []int{-399, -398, -397, -396, -395, -394, -393, -392, -391, -390, -389, -388, -387, -299, -298, -297, -296, -295, -294, -293, -292, -291, -290, -289, -288, -287, -286, -285, -284, -283, -282, -281, -280, -279, -278, -277, -276, -275, -274, -273, -272, -271, -270, -269, -268, -267, -266, -265, -264, -263, -199, -198, -197, -196, -195, -194, -193, -192, -191, -190, -189, -188, -187, -186, -185, -184, -183, -182, -181, -180, -179, -99, -98, -97, -96, -95, -94, -93, -92, -91, -90, -89, -88, 0, 100, 101, 102, 103, 104, 105, 121, 122, 123, 124, 125, 126, 127, 128, 129, 130, 131, 132, 133, 134, 135, 136, 137, 138, 139, 140, 141, 142, 143, 144, 145, 146, 147, 148, 149, 150, 151, 152, 154, 155, 156, 157, 158, 159, 160, 161, 162, 163},
		},
	}

	for _, test := range tt {
		t.Run(test.description, func(t *testing.T) {
			keys, values := definitions.SortMapByValue(test.input)

			assert.Equal(t, test.expectedKeys, keys)
			assert.Equal(t, test.expectedValues, values)
		})
	}

}

func TestBinarySearch(t *testing.T) {
	tt := []struct {
		description string
		inputCode   int
		inputMap    map[string]int
		expected    string
	}{
		{
			description: "successfully found `NFTokenOffer` (LedgerEntryTypes)",
			inputCode:   55,
			inputMap:    definitions.LedgerEntryTypes,
			expected:    "NFTokenOffer",
		},
		{
			description: "successfully found `Amount` (Types)",
			inputCode:   6,
			inputMap:    definitions.Types,
			expected:    "Amount",
		},
		{
			description: "successfully found `Contract` (TransactionTypes)",
			inputCode:   9,
			inputMap:    definitions.TransactionTypes,
			expected:    "Contract",
		},
		{
			description: "successfully found `tecINSUFFICIENT_PAYMENT` (TransactionResults)",
			inputCode:   161,
			inputMap:    definitions.TransactionResults,
			expected:    "tecINSUFFICIENT_PAYMENT",
		},
		// {
		// 	description: "non-existent type code ",
		// 	inputCode:   6666,
		// 	inputMap:    definitions.LedgerEntryTypes,
		// 	expected:    "",
		// },
	}

	for _, test := range tt {
		t.Run(test.description, func(t *testing.T) {
			keys, err := definitions.BinaryGetNameByCode(test.inputCode, test.inputMap)
			assert.NoError(t, err)
			assert.Equal(t, test.expected, keys)
			// NEED TO ADD ERROR HANDLING
		})
	}
}
