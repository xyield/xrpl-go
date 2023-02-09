package test

import (
	"encoding/json"
	"fmt"
	"reflect"
)

func SerializeAndDeserialize(s interface{}, d string) error {
	j, err := json.MarshalIndent(s, "", "\t")
	if err != nil {
		return err
	}
	if string(j) != d {
		fmt.Println(string(j))
		fmt.Println(d)
		return fmt.Errorf("json encoding does not match expected string")
	}
	decode := reflect.New(reflect.TypeOf(s))
	err = json.Unmarshal(j, decode.Interface())
	if err != nil {
		return err
	}
	if !reflect.DeepEqual(s, decode.Elem().Interface()) {
		fmt.Printf("%+v\n", s)
		fmt.Printf("%+v\n", decode.Elem().Interface())
		return fmt.Errorf("json decoding does not match expected struct")
	}
	return nil
}
