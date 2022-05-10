package definitions

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

func (d *Definitions) GetTypeCodeByFieldName(n string) (int32, error) {
	typeName, err := d.GetTypeNameByFieldName(n)

	if err != nil {
		return 0, err
	}

	return d.Types[typeName], nil
}

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

func (d *Definitions) GetFieldHeaderByFieldName(n string) (*fieldHeader, error) {

	fi, ok := d.Fields[n]

	if !ok {
		return nil, &NotFoundError{
			Instance: "FieldName",
			Input:    n,
		}
	}

	return fi.FieldHeader, nil
}

func (d *Definitions) GetFieldNameByFieldHeader(fh fieldHeader) (string, error) {

	fim, ok := definitions.FieldIdNameMap[fh]

	if !ok {
		return "", &NotFoundErrorFieldHeader{
			Instance: "FieldHeader",
			Input:    fh,
		}
	}
	return fim, nil
}

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

func (d *Definitions) GetFieldInstanceByFieldName(n string) (*fieldInstance, error) {

	fi, ok := d.Fields[n]

	if !ok {
		return nil, &NotFoundError{
			Instance: "FieldName",
			Input:    n,
		}
	}
	return fi, nil
}

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
