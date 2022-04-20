package definitions

import (
	"reflect"

	"github.com/ugorji/go/codec"
)

// AnyJson type is used when deserialising an unknown fields into a Golang struct.
type AnyJson map[string]interface{}

// UnmarshalJSON unmarshals a response body into the AnyJson type.
func (a *AnyJson) UnmarshalJSON(data []byte) error {

	var jh codec.JsonHandle

	jh.SignedInteger = true
	jh.MapType = reflect.TypeOf(map[string]interface{}{})

	err := codec.NewDecoderBytes(data, &jh).Decode(a)

	return err
}
