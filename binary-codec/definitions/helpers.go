package definitions

import (
	"fmt"
	"sort"
)

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

func (d *Definitions) GetTypeCodeByTypeName(n string) (int, error) {
	typeCode, ok := d.Types[n]

	if !ok {
		return 0, &NotFoundError{
			Instance: "TypeName",
			Input:    n,
		}
	}
	return typeCode, nil
}

func (d *Definitions) GetTypeCodeByFieldName(n string) (int, error) {
	typeName, err := d.GetTypeNameByFieldName(n)

	if err != nil {
		return 0, err
	}

	return d.Types[typeName], nil
}

func (d *Definitions) GetFieldCodeByFieldName(n string) (int, error) {

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

func (d *Definitions) GetTransactionTypeCodeByTransactionTypeName(n string) (int, error) {
	txTypeCode, ok := d.TransactionTypes[n]

	if !ok {
		return 0, &NotFoundError{
			Instance: "TransactionTypeName",
			Input:    n,
		}
	}
	return txTypeCode, nil
}

func (d *Definitions) GetTransactionTypeNameByTransactionTypeCode(c int) (string, error) {

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

func (d *Definitions) GetTransactionResultNameByTransactionResultTypeCode(c int) (string, error) {

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

func (d *Definitions) GetTransactionResultTypeCodeByTransactionResultName(n string) (int, error) {

	txResultTypeCode, ok := d.TransactionResults[n]

	if !ok {
		return 0, &NotFoundError{
			Instance: "TransactionResultName",
			Input:    n,
		}
	}
	return txResultTypeCode, nil
}

func (d *Definitions) GetLedgerEntryTypeCodeByLedgerEntryTypeName(n string) (int, error) {

	ledgerEntryTypeCode, ok := d.LedgerEntryTypes[n]

	if !ok {
		return 0, &NotFoundError{
			Instance: "LedgerEntryTypeName",
			Input:    n,
		}
	}
	return ledgerEntryTypeCode, nil
}

func (d *Definitions) GetLedgerEntryTypeNameByLedgerEntryTypeCode(c int) (string, error) {

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

func (d *Definitions) BinaryGetNameByCode(c int, vmap map[string]int) (string, error) {

	k, tc := definitions.SortMapByValue(vmap)
	i := definitions.BinarySearch(tc, 0, len(tc)-1, c)

	// NEED TO ADD ERROR HANDLING

	fmt.Printf("Found Name: `%v` from Code: %v ", k[i], c)

	return k[i], nil
}

func (d *Definitions) SortMapByValue(vmap map[string]int) (sortedKeys []string, sortedValues []int) {

	keys := make([]string, 0, len(vmap))

	for key := range vmap {
		keys = append(keys, key)
	}

	// fmt.Println("KEYS (Before Sorting):", keys)

	sort.SliceStable(keys, func(i, j int) bool {
		return vmap[keys[i]] < vmap[keys[j]]
	})

	// fmt.Println("KEYS (After Sorting)", keys)

	values := make([]int, 0, len(vmap))

	for _, value := range vmap {
		values = append(values, value)
	}

	// fmt.Println("VALUES (Before Sorting:)", values)

	sort.Slice(values, func(i, j int) bool {
		return values[i] < values[j]
	})
	// fmt.Println("VALUES (After Sorting:)", values)

	return keys, values
}

func (d *Definitions) BinarySearch(numbers []int, leftBound, rightBound, numberToFind int) int {
	for leftBound <= rightBound {
		midPoint := leftBound + (rightBound-leftBound)/2

		if numbers[midPoint] == numberToFind {
			return midPoint
		}

		if numbers[midPoint] > numberToFind {
			rightBound = midPoint - 1
		} else {
			leftBound = midPoint + 1
		}
	}

	return -1
}
