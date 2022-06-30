package definitions

import (
	_ "embed"
	"errors"
	"fmt"

	"github.com/ugorji/go/codec"
)

var definitions *Definitions

//go:embed definitions.json
var docBytes []byte

func Get() *Definitions {
	return definitions
}

type NotFoundError struct {
	Instance string
	Input    string
}

func (e *NotFoundError) Error() string {
	return fmt.Sprintf("%v %v not found", e.Instance, e.Input)
}

type NotFoundErrorInt struct {
	Instance string
	Input    int32
}

func (e *NotFoundErrorInt) Error() string {
	return fmt.Sprintf("%v %v not found", e.Instance, e.Input)
}

type NotFoundErrorFieldHeader struct {
	Instance string
	Input    fieldHeader
}

func (e *NotFoundErrorFieldHeader) Error() string {
	return fmt.Sprintf("%v %v not found", e.Instance, e.Input)
}

type Definitions struct {
	Types              map[string]int32
	LedgerEntryTypes   map[string]int32
	Fields             fieldInstanceMap
	TransactionResults map[string]int32
	TransactionTypes   map[string]int32
	FieldIdNameMap     map[fieldHeader]string
}
type definitionsDoc struct {
	Types              map[string]int32 `json:"TYPES"`
	LedgerEntryTypes   map[string]int32 `json:"LEDGER_ENTRY_TYPES"`
	Fields             fieldInstanceMap `json:"FIELDS"`
	TransactionResults map[string]int32 `json:"TRANSACTION_RESULTS"`
	TransactionTypes   map[string]int32 `json:"TRANSACTION_TYPES"`
}

type fieldInstanceMap map[string]*FieldInstance

func (fi *fieldInstanceMap) CodecEncodeSelf(e *codec.Encoder) {}

func (fi *fieldInstanceMap) CodecDecodeSelf(d *codec.Decoder) {
	var x [][]interface{}
	d.MustDecode(&x)
	y := convertToFieldInstanceMap(x)
	*fi = y
}

// Loads JSON from the definitions file and converts it to a preferred format.
// The definitions file contains information required for the XRP Ledger's
// canonical binary serialization format:
// `Serialization <https://xrpl.org/serialization.html>`_
func loadDefinitions() error {

	var jh codec.JsonHandle

	jh.MapKeyAsString = true
	jh.SignedInteger = true

	dec := codec.NewDecoderBytes(docBytes, &jh)
	var data definitionsDoc
	err := dec.Decode(&data)
	if err != nil {
		return err
	}
	definitions = &Definitions{
		Types:              data.Types,
		Fields:             data.Fields,
		LedgerEntryTypes:   data.LedgerEntryTypes,
		TransactionResults: data.TransactionResults,
		TransactionTypes:   data.TransactionTypes,
	}

	addFieldHeadersAndOrdinals()
	createFieldIdNameMap()

	return nil
}

func convertToFieldInstanceMap(m [][]interface{}) map[string]*FieldInstance {
	nm := make(map[string]*FieldInstance, len(m))

	for _, j := range m {
		k := j[0].(string)
		fi, _ := castFieldInfo(j[1])
		nm[k] = &FieldInstance{
			FieldName: k,
			fieldInfo: &fi,
			Ordinal:   fi.Nth,
		}
	}
	return nm
}

func castFieldInfo(v interface{}) (fieldInfo, error) {
	if fi, ok := v.(map[string]interface{}); ok {
		return fieldInfo{
			Nth:            int32(fi["nth"].(int64)),
			IsVLEncoded:    fi["isVLEncoded"].(bool),
			IsSerialized:   fi["isSerialized"].(bool),
			IsSigningField: fi["isSigningField"].(bool),
			Type:           fi["type"].(string),
		}, nil
	}
	return fieldInfo{}, errors.New("unable to cast to field info")
}

func addFieldHeadersAndOrdinals() {
	for k := range definitions.Fields {
		t, _ := definitions.GetTypeCodeByTypeName(definitions.Fields[k].Type)

		if fi, ok := definitions.Fields[k]; ok {
			fi.FieldHeader = &fieldHeader{
				TypeCode:  t,
				FieldCode: definitions.Fields[k].Nth,
			}
			fi.Ordinal = (t<<16 | definitions.Fields[k].Nth)
		}
	}
}

func createFieldIdNameMap() {
	definitions.FieldIdNameMap = make(map[fieldHeader]string, len(definitions.Fields))
	for k := range definitions.Fields {
		fh, _ := definitions.GetFieldHeaderByFieldName(k)

		definitions.FieldIdNameMap[*fh] = k
	}
}

func convertIntToBytes(i int32) []byte {

	return []byte{3}
}
