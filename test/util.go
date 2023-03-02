package test

import (
	"encoding/json"
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

func SerializeAndDeserialize(t *testing.T, s interface{}, d string) error {
	j, err := json.MarshalIndent(s, "", "\t")
	if err != nil {
		return err
	}
	assert.Equal(t, d, string(j), "json encoding does not match expected string")
	decode := reflect.New(reflect.TypeOf(s))
	err = json.Unmarshal(j, decode.Interface())
	if err != nil {
		return err
	}
	assert.Equal(t, s, decode.Elem().Interface(), "json decoding does not match expected struct")
	return nil
}
