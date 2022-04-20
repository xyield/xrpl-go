package definitions

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"path"
	"runtime"
)

var definitions Definitions

type Definitions struct {
	Types              map[string]int64         `json:"TYPES"`
	LedgerEntryTypes   map[string]int64         `json:"LEDGER_ENTRY_TYPES"`
	Fields             map[string]fieldInstance `json:"FIELDS"`
	TransactionResults map[string]int64         `json:"TRANSACTION_RESULTS"`
	TransactionTypes   map[string]int64         `json:"TRANSACTION_TYPES"`
}

func LoadDefinitions() error {

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

	return nil
}

func castMap(m map[string]interface{}) map[string]int64 {
	nm := make(map[string]int64)
	for k, v := range m {
		nm[k] = v.(int64)
	}
	return nm
}

func convertToFieldInstanceMap(m []interface{}) map[string]fieldInstance {
	nm := make(map[string]fieldInstance, len(m))

	for _, j := range m {
		if v, ok := j.([]interface{}); ok {
			k := v[0].(string)
			fi, _ := castFieldInfo(v[1])
			nm[k] = fieldInstance{
				FieldName: k,
				FieldInfo: fi,
			}
			fmt.Printf("%+v\n", nm[k])
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
