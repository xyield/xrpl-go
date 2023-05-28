package test

import (
	"encoding/json"
	"fmt"
	"reflect"
	"testing"

	"github.com/stretchr/testify/require"
)

func Deserialize(s interface{}, d string) error {
	decode := reflect.New(reflect.TypeOf(s))
	err := json.Unmarshal([]byte(d), decode.Interface())
	if err != nil {
		return err
	}
	if !reflect.DeepEqual(s, decode.Elem().Interface()) {
		return fmt.Errorf("json decoding does not match expected struct")
	}
	return nil

}

func SerializeAndDeserialize(t *testing.T, s interface{}, d string) error {
	j, err := json.MarshalIndent(s, "", "\t")
	if err != nil {
		return err
	}
	require.Equal(t, d, string(j), "json encoding does not match expected string")
	decode := reflect.New(reflect.TypeOf(s))
	err = json.Unmarshal(j, decode.Interface())
	if err != nil {
		return err
	}
	require.Equal(t, s, decode.Elem().Interface(), "json decoding does not match expected struct")
	return nil
}
