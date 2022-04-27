package definitions

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLoadDefinitions(t *testing.T) {

	err := loadDefinitions()
	assert.NoError(t, err)
	assert.Equal(t, int64(-1), definitions.Types["Done"])
	assert.Equal(t, int64(4), definitions.Types["Hash128"])
	assert.Equal(t, int64(-3), definitions.LedgerEntryTypes["Any"])
	assert.Equal(t, int64(-399), definitions.TransactionResults["telLOCAL_ERROR"])
	assert.Equal(t, int64(1), definitions.TransactionTypes["EscrowCreate"])
	assert.Equal(t, fieldInfo{Nth: int64(0), IsVLEncoded: false, IsSerialized: false, IsSigningField: false, Type: "Unknown"}, definitions.Fields["Generic"].FieldInfo)
	assert.Equal(t, fieldInfo{Nth: int64(28), IsVLEncoded: false, IsSerialized: true, IsSigningField: true, Type: "Hash256"}, definitions.Fields["NFTokenBuyOffer"].FieldInfo)
	assert.Equal(t, fieldInfo{Nth: int64(16), IsVLEncoded: false, IsSerialized: true, IsSigningField: true, Type: "UInt8"}, definitions.Fields["TickSize"].FieldInfo)
	assert.Equal(t, fieldHeader{TypeCode: 2, FieldCode: 4}, definitions.Fields["Sequence"].FieldHeader)
	assert.Equal(t, fieldHeader{TypeCode: 18, FieldCode: 1}, definitions.Fields["Paths"].FieldHeader)
	assert.Equal(t, fieldHeader{TypeCode: 2, FieldCode: 33}, definitions.Fields["SetFlag"].FieldHeader)
	assert.Equal(t, fieldHeader{TypeCode: 16, FieldCode: 16}, definitions.Fields["TickSize"].FieldHeader)
	assert.Equal(t, definitions.Types["Done"], int64(-1))
	assert.Equal(t, "UInt32", definitions.Fields["TransferRate"].FieldInfo.Type)

}

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
			description:   "test that non-existent value gives an error",
			input:         "yurt",
			expected:      "",
			expectedError: &TypeNotFoundError{},
		},
	}

	for _, test := range tt {

		t.Run(test.description, func(t *testing.T) {
			got, err := definitions.GetTypeNameByFieldName(test.input)
			if test.expectedError != nil {
				assert.Error(t, test.expectedError, err.Error())
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
		expected      int64
		expectedError error
	}{
		{
			description:   "test that `Done` gives correct code",
			input:         "Done",
			expected:      int64(-1),
			expectedError: nil,
		},
		{
			description:   "test that `Hash128` gives correct code",
			input:         "Hash128",
			expected:      int64(4),
			expectedError: nil,
		},
		{
			description:   "test that incorrect value gives an error",
			input:         "yurt",
			expected:      0,
			expectedError: &TypeNotFoundError{},
		},
	}

	for _, test := range tt {

		t.Run(test.description, func(t *testing.T) {
			got, err := definitions.GetTypeCodeByTypeName(test.input)
			if test.expectedError != nil {
				assert.Error(t, test.expectedError, err.Error())
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
		expected      int64
		expectedError error
	}{
		{
			description:   "test that `TransferRate` gives `int64(2)`",
			input:         "TransferRate",
			expected:      int64(2),
			expectedError: nil,
		},
		{
			description:   "test that `OwnerNode` gives `int64(3)`",
			input:         "OwnerNode",
			expected:      int64(3),
			expectedError: nil,
		},
		{
			description:   "test that non-existent value gives error",
			input:         "yurt",
			expected:      0,
			expectedError: &TypeNotFoundError{},
		},
	}

	for _, test := range tt {

		t.Run(test.description, func(t *testing.T) {
			got, err := definitions.GetTypeCodeByFieldName(test.input)
			if test.expectedError != nil {
				assert.Error(t, test.expectedError, err.Error())
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
		expected      int64
		expectedError error
	}{
		{
			description:   "correct FieldCode",
			input:         "TransferRate",
			expected:      int64(11),
			expectedError: nil,
		},
		{
			description:   "non-existent FieldName",
			input:         "yurt",
			expected:      0,
			expectedError: &TypeNotFoundError{},
		},
	}

	for _, test := range tt {

		t.Run(test.description, func(t *testing.T) {
			got, err := definitions.GetFieldCodeByFieldName(test.input)
			if test.expectedError != nil {
				assert.Error(t, test.expectedError, err.Error())
				assert.Zero(t, got)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, test.expected, got)
			}
		})
	}
}
