package definitions

import (
	"errors"
	"io/ioutil"
	"log"
	"path"
	"runtime"

	"github.com/ugorji/go/codec"
)

var definitions Definitions

type TypeNotFoundError struct{}

func (tnf *TypeNotFoundError) Error() string {
	return "Type not found"
}

type Definitions struct {
	Types              map[string]int   `json:"TYPES"`
	LedgerEntryTypes   map[string]int   `json:"LEDGER_ENTRY_TYPES"`
	Fields             fieldInstanceMap `json:"FIELDS"`
	TransactionResults map[string]int   `json:"TRANSACTION_RESULTS"`
	TransactionTypes   map[string]int   `json:"TRANSACTION_TYPES"`
}

type fieldInstanceMap map[string]*fieldInstance

func (fi *fieldInstanceMap) CodecEncodeSelf(e *codec.Encoder) {}

func (fi *fieldInstanceMap) CodecDecodeSelf(d *codec.Decoder) {
	var x [][]interface{}
	d.MustDecode(&x)
	y := convertToFieldInstanceMap(x)
	*fi = y
}

func (d *Definitions) GetTypeNameByFieldName(n string) (string, error) {

	fieldName, ok := d.Fields[n]

	if !ok {
		return "", &TypeNotFoundError{}
	}

	typeName := fieldName.Type

	return typeName, nil
}

func (d *Definitions) GetTypeCodeByTypeName(n string) (int, error) {
	typeCode, ok := d.Types[n]

	if !ok {
		return 0, &TypeNotFoundError{}
	}
	return typeCode, nil
}

func (d *Definitions) GetTypeCodeByFieldName(n string) (int, error) {
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

func (d *Definitions) GetFieldCodeByFieldName(n string) (int, error) {

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

func (d *Definitions) GetTransactionTypeCodeByTransactionTypeName(n string) (int, error) {
	txTypeCode, ok := d.TransactionTypes[n]

	if !ok {
		return 0, &TypeNotFoundError{}
	}

	return txTypeCode, nil
}

func (d *Definitions) GetTransactionTypeNameByTransactionTypeCode(c int) (string, error) {
	return "EscrowCreate", nil
}

func (d *Definitions) GetTransactionResultNameByTransactionResultTypeCode(c int) (string, error) {
	return "", nil
}

func (d *Definitions) GetTransactionResultTypeCodeByTransactionResultName(n string) (int, error) {
	return 0, nil
}

func (d *Definitions) GetLedgerEntryTypeCodeByLedgerEntryTypeName(n string) (int, error) {
	return 0, nil
}

func (d *Definitions) GetLedgerEntryTypeNameByLedgerEntryTypeCode(c int) (string, error) {
	return "", nil
}

func loadDefinitions() error {

	_, f, _, _ := runtime.Caller(0)
	wd := path.Dir(f)
	docBytes, err := ioutil.ReadFile(wd + "/definitions.json")
	if err != nil {
		return err
	}

	var jh codec.JsonHandle

	jh.MapKeyAsString = true
	jh.SignedInteger = true

	dec := codec.NewDecoderBytes(docBytes, &jh)

	err = dec.Decode(&definitions)
	if err != nil {
		return err
	}

	addFieldHeaders()

	return nil
}

func convertToFieldInstanceMap(m [][]interface{}) map[string]*fieldInstance {
	nm := make(map[string]*fieldInstance, len(m))

	for _, j := range m {
		k := j[0].(string)
		fi, _ := castFieldInfo(j[1])
		nm[k] = &fieldInstance{
			FieldName: k,
			fieldInfo: fi,
		}
	}
	return nm
}

func castFieldInfo(v interface{}) (fieldInfo, error) {
	if fi, ok := v.(map[string]interface{}); ok {
		return fieldInfo{
			Nth:            int(fi["nth"].(int64)),
			IsVLEncoded:    fi["isVLEncoded"].(bool),
			IsSerialized:   fi["isSerialized"].(bool),
			IsSigningField: fi["isSigningField"].(bool),
			Type:           fi["type"].(string),
		}, nil
	}
	return fieldInfo{}, errors.New("unable to cast to field info")
}

func addFieldHeaders() {
	for k, _ := range definitions.Fields {
		t, _ := definitions.GetTypeCodeByTypeName(definitions.Fields[k].Type)
		if fi, ok := definitions.Fields[k]; ok {
			fi.FieldHeader = fieldHeader{
				TypeCode:  byte(t),
				FieldCode: byte(definitions.Fields[k].Nth),
			}
		}
	}
}
