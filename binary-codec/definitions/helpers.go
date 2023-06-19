package definitions

// Returns the serialization data type for the given field name.
func (d *Definitions) GetTypeNameByFieldName(n string) (string, error) {

	fi, ok := d.Fields[n]

	if !ok {
		return "", &NotFoundError{
			Instance: "FieldName",
			Input:    n,
		}
	}

	return fi.Type, nil
}

// Returns the type code associated with the given type name.
func (d *Definitions) GetTypeCodeByTypeName(n string) (int32, error) {
	typeCode, ok := d.Types[n]

	if !ok {
		return 0, &NotFoundError{
			Instance: "TypeName",
			Input:    n,
		}
	}
	return typeCode, nil
}

// Returns the type code associated with the given field name.
func (d *Definitions) GetTypeCodeByFieldName(n string) (int32, error) {
	typeName, err := d.GetTypeNameByFieldName(n)

	if err != nil {
		return 0, err
	}

	return d.Types[typeName], nil
}

// Returns the field code associated with the given field name.
func (d *Definitions) GetFieldCodeByFieldName(n string) (int32, error) {

	fi, ok := d.Fields[n]

	if !ok {
		return 0, &NotFoundError{
			Instance: "FieldName",
			Input:    n,
		}
	}

	return fi.Nth, nil
}

// Returns the field header struct associated with the given field name.
func (d *Definitions) GetFieldHeaderByFieldName(n string) (*FieldHeader, error) {

	fi, ok := d.Fields[n]

	if !ok {
		return nil, &NotFoundError{
			Instance: "FieldName",
			Input:    n,
		}
	}

	return fi.FieldHeader, nil
}

// Returns the field name associated with the given field header struct.
func (d *Definitions) GetFieldNameByFieldHeader(fh FieldHeader) (string, error) {

	fim, ok := definitions.FieldIdNameMap[fh]

	if !ok {
		return "", &NotFoundErrorFieldHeader{
			Instance: "FieldHeader",
			Input:    fh,
		}
	}
	return fim, nil
}

// Returns the field info struct associated with the given field name.
func (d *Definitions) GetFieldInfoByFieldName(n string) (*fieldInfo, error) {

	fi, ok := d.Fields[n]

	if !ok {
		return nil, &NotFoundError{
			Instance: "FieldName",
			Input:    n,
		}
	}

	return fi.fieldInfo, nil
}

// Returns the field instance struct associated with the given field name.
func (d *Definitions) GetFieldInstanceByFieldName(n string) (*FieldInstance, error) {

	fi, ok := d.Fields[n]

	if !ok {
		return nil, &NotFoundError{
			Instance: "FieldName",
			Input:    n,
		}
	}
	return fi, nil
}

// Returns the transaction type code associated with the transaction type name.
func (d *Definitions) GetTransactionTypeCodeByTransactionTypeName(n string) (int32, error) {
	txTypeCode, ok := d.TransactionTypes[n]

	if !ok {
		return 0, &NotFoundError{
			Instance: "TransactionTypeName",
			Input:    n,
		}
	}
	return txTypeCode, nil
}

// Returns the transaction type name associated with the transaction type code.
func (d *Definitions) GetTransactionTypeNameByTransactionTypeCode(c int32) (string, error) {

	for txTypeName, code := range d.TransactionTypes {
		if code == c {
			return txTypeName, nil
		}
	}
	return "", &NotFoundErrorInt{
		Instance: "TransactionTypeCode",
		Input:    c,
	}
}

// Returns the transaction result name associated with the transaction result type code.
func (d *Definitions) GetTransactionResultNameByTransactionResultTypeCode(c int32) (string, error) {

	for txResultName, code := range d.TransactionResults {
		if code == c {
			return txResultName, nil
		}
	}

	return "", &NotFoundErrorInt{
		Instance: "TransactionResultTypeCode",
		Input:    c,
	}
}

// Returns the transaction result type code associated with the transaction result name.
func (d *Definitions) GetTransactionResultTypeCodeByTransactionResultName(n string) (int32, error) {

	txResultTypeCode, ok := d.TransactionResults[n]

	if !ok {
		return 0, &NotFoundError{
			Instance: "TransactionResultName",
			Input:    n,
		}
	}
	return txResultTypeCode, nil
}

// Returns the ledger entry type code associated with the ledger entry type name.
func (d *Definitions) GetLedgerEntryTypeCodeByLedgerEntryTypeName(n string) (int32, error) {

	ledgerEntryTypeCode, ok := d.LedgerEntryTypes[n]

	if !ok {
		return 0, &NotFoundError{
			Instance: "LedgerEntryTypeName",
			Input:    n,
		}
	}
	return ledgerEntryTypeCode, nil
}

// Returns the ledger entry type name associated with the ledger entry type code.
func (d *Definitions) GetLedgerEntryTypeNameByLedgerEntryTypeCode(c int32) (string, error) {

	for ledgerEntryTypeName, code := range d.LedgerEntryTypes {

		if code == c {
			return ledgerEntryTypeName, nil
		}
	}

	return "", &NotFoundErrorInt{
		Instance: "LedgerEntryTypeCode",
		Input:    c,
	}
}
