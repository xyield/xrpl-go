package definitions

import (
	"errors"
	"fmt"
	"io/ioutil"
	"path"
	"runtime"

	"github.com/ugorji/go/codec"
)

var definitions *Definitions

func Get() *Definitions {
	return definitions
}

type NotFoundError struct {
	Instance string
	Input    string
}

func (tnf *NotFoundError) Error() string {
	return fmt.Sprintf("%v %v not found", tnf.Instance, tnf.Input)
}

type NotFoundErrorInt struct {
	Instance string
	Input    int
}

func (tnf *NotFoundErrorInt) Error() string {
	return fmt.Sprintf("%v %v not found", tnf.Instance, tnf.Input)
}

type NotFoundErrorFieldHeader struct {
	Instance string
	Input    fieldHeader
}

func (tnf *NotFoundErrorFieldHeader) Error() string {
	return fmt.Sprintf("%v %v not found", tnf.Instance, tnf.Input)
}

type Definitions struct {
	Types              map[string]int
	LedgerEntryTypes   map[string]int
	Fields             fieldInstanceMap
	TransactionResults map[string]int
	TransactionTypes   map[string]int
	FieldIdNameMap     map[fieldHeader]string
}
type definitionsDoc struct {
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
	var data definitionsDoc
	err = dec.Decode(&data)
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

	addFieldHeaders()
	createFieldIdNameMap()

	return nil
}

func convertToFieldInstanceMap(m [][]interface{}) map[string]*fieldInstance {
	nm := make(map[string]*fieldInstance, len(m))

	for _, j := range m {
		k := j[0].(string)
		fi, _ := castFieldInfo(j[1])
		nm[k] = &fieldInstance{
			FieldName: k,
			fieldInfo: &fi,
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
			fi.FieldHeader = &fieldHeader{
				TypeCode:  byte(t),
				FieldCode: byte(definitions.Fields[k].Nth),
			}
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
