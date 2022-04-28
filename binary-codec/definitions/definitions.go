package definitions

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"log"
	"path"
	"runtime"
)

var definitions Definitions

type TypeNotFoundError struct{}

func (tnf *TypeNotFoundError) Error() string {
	return "Type not found"
}

type Definitions struct {
	Types              map[string]int64          `json:"TYPES"`
	LedgerEntryTypes   map[string]int64          `json:"LEDGER_ENTRY_TYPES"`
	Fields             map[string]*fieldInstance `json:"FIELDS"`
	TransactionResults map[string]int64          `json:"TRANSACTION_RESULTS"`
	TransactionTypes   map[string]int64          `json:"TRANSACTION_TYPES"`
}

func (d *Definitions) GetTypeNameByFieldName(n string) (string, error) {

	fieldName, ok := d.Fields[n]

	if !ok {
		return "", &TypeNotFoundError{}
	}

	typeName := fieldName.Type

	return typeName, nil
}

func (d *Definitions) GetTypeCodeByTypeName(n string) (int64, error) {
	typeCode, ok := d.Types[n]

	if !ok {
		return 0, &TypeNotFoundError{}
	}
	return typeCode, nil
}

func (d *Definitions) GetTypeCodeByFieldName(n string) (int64, error) {
	typeName, err := d.GetTypeNameByFieldName(n)

	if err != nil {
		log.Println("TypeName not found from the FieldName provided.")
		return 0, err
	}

	typeCode, ok := d.Types[typeName]

	if !ok {
		return 0, &TypeNotFoundError{}
	}

	return typeCode, nil
}

func (d *Definitions) GetFieldCodeByFieldName(n string) (int64, error) {

	fieldName, ok := d.Fields[n]

	if !ok {
		return 0, &TypeNotFoundError{}
	}

	return fieldName.Nth, nil
}

func (d *Definitions) GetFieldHeaderByFieldName(n string) (fieldHeader, error) {

	fieldCode, err := d.GetFieldCodeByFieldName(n)

	if err != nil {
		return fieldHeader{}, &TypeNotFoundError{}
	}

	typeCode, _ := d.GetTypeCodeByFieldName(n)

	return fieldHeader{
		TypeCode:  byte(typeCode),
		FieldCode: byte(fieldCode),
	}, nil
}

func (d *Definitions) GetFieldInfoByFieldName(n string) (fieldInfo, error) {

	fieldName, ok := d.Fields[n]

	if !ok {
		return fieldInfo{}, &TypeNotFoundError{}
	}

	return fieldInfo{
		Nth:            fieldName.Nth,
		IsVLEncoded:    fieldName.IsVLEncoded,
		IsSerialized:   fieldName.IsSerialized,
		IsSigningField: fieldName.IsSigningField,
		Type:           fieldName.Type,
	}, nil
}

func (d *Definitions) GetFieldInstanceByFieldName(n string) (fieldInstance, error) {

	fieldHeader, err := d.GetFieldHeaderByFieldName(n)

	if err != nil {
		return fieldInstance{}, &TypeNotFoundError{}
	}

	fieldInfo, _ := d.GetFieldInfoByFieldName(n)

	if err != nil {
		return fieldInstance{}, &TypeNotFoundError{}
	}

	return fieldInstance{
		FieldName:   n,
		fieldInfo:   fieldInfo,
		FieldHeader: fieldHeader,
	}, nil
}

func (d *Definitions) GetTransactionTypeCodeByTransactionTypeName(n string) (int64, error) {
	txTypeCode, ok := d.TransactionTypes[n]

	if !ok {
		return 0, &TypeNotFoundError{}
	}

	return txTypeCode, nil
}

func (d *Definitions) GetTransactionTypeNameByTransactionTypeCode(c int64) (string, error) {
	return "EscrowCreate", nil
}

func (d *Definitions) GetTransactionResultNameByTransactionResultTypeCode(c int64) (string, error) {
	return "", nil
}

func (d *Definitions) GetTransactionResultTypeCodeByTransactionResultName(n string) (int64, error) {
	return 0, nil
}

func (d *Definitions) GetLedgerEntryTypeCodeByLedgerEntryTypeName(n string) (int64, error) {
	return 0, nil
}

func (d *Definitions) GetLedgerEntryTypeNameByLedgerEntryTypeCode(c int64) (string, error) {
	return "", nil
}

func loadDefinitions() error {

	_, f, _, _ := runtime.Caller(0)
	wd := path.Dir(f)
	docBytes, err := ioutil.ReadFile(wd + "/definitions.json")
	if err != nil {
		return err
	}

	var jsonDoc AnyJson
	err = json.Unmarshal(docBytes, &jsonDoc)
	if err != nil {
		return err
	}

	types := jsonDoc["TYPES"].(map[string]interface{})
	ledgerEntryTypes := jsonDoc["LEDGER_ENTRY_TYPES"].(map[string]interface{})
	transactionResults := jsonDoc["TRANSACTION_RESULTS"].(map[string]interface{})
	transactionTypes := jsonDoc["TRANSACTION_TYPES"].(map[string]interface{})
	fields := jsonDoc["FIELDS"].([]interface{})

	definitions.Types = castMap(types)
	definitions.LedgerEntryTypes = castMap(ledgerEntryTypes)
	definitions.TransactionResults = castMap(transactionResults)
	definitions.TransactionTypes = castMap(transactionTypes)
	definitions.Fields = convertToFieldInstanceMap(fields)
	addFieldHeaders(definitions.Types, definitions.Fields)

	return nil
}

func castMap(m map[string]interface{}) map[string]int64 {
	nm := make(map[string]int64)
	for k, v := range m {
		nm[k] = v.(int64)
	}
	return nm
}

func convertToFieldInstanceMap(m []interface{}) map[string]*fieldInstance {
	nm := make(map[string]*fieldInstance, len(m))

	for _, j := range m {
		if v, ok := j.([]interface{}); ok {
			k := v[0].(string)
			fi, _ := castFieldInfo(v[1])
			nm[k] = &fieldInstance{
				FieldName: k,
				fieldInfo: fi,
			}
		}
	}
	return nm
}

func castFieldInfo(v interface{}) (fieldInfo, error) {
	if m, ok := v.(map[string]interface{}); ok {
		return fieldInfo{
			Nth:            m["nth"].(int64),
			IsVLEncoded:    m["isVLEncoded"].(bool),
			IsSerialized:   m["isSerialized"].(bool),
			IsSigningField: m["isSigningField"].(bool),
			Type:           m["type"].(string),
		}, nil
	}
	return fieldInfo{}, errors.New("unable to cast to field info")
}

func addFieldHeaders(typeMap map[string]int64, fieldInstances map[string]*fieldInstance) {
	for k, _ := range fieldInstances {
		t := typeMap[fieldInstances[k].Type]
		if fi, ok := fieldInstances[k]; ok {
			fi.FieldHeader = fieldHeader{
				TypeCode:  byte(t),
				FieldCode: byte(fieldInstances[k].Nth),
			}
		}
	}
}
