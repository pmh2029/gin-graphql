package utils

import (
	"encoding/json"
	"reflect"

	"github.com/mitchellh/mapstructure"
)

func StructToStruct(
	input interface{},
	output interface{},
) error {
	decoder, err := mapstructure.NewDecoder(&mapstructure.DecoderConfig{
		TagName:     "mapstructure",
		ErrorUnused: false,
		ZeroFields:  true,
		Result:      output,
	})
	if err != nil {
		return err
	}

	return decoder.Decode(input)
}

func StructToMap(
	input interface{},
) (map[string]interface{}, error) {
	output := make(map[string]interface{})
	jsonData, err := json.Marshal(input)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(jsonData, &output)
	if err != nil {
		return nil, err
	}

	for key := range output {
		output[key] = reflect.Zero(reflect.TypeOf(output[key])).Interface()
	}
	return output, nil
}

func CompareStruct(input, target interface{}) bool {
	t1 := reflect.TypeOf(input)
	t2 := reflect.TypeOf(target)

	if t1.Kind() != reflect.Struct || t2.Kind() != reflect.Struct {
		return false
	}

	fields1 := make(map[string]string)
	fields2 := make(map[string]string)

	for i := 0; i < t1.NumField(); i++ {
		field1 := t1.Field(i)
		fields1[field1.Name] = field1.Type.String()
	}

	for i := 0; i < t2.NumField(); i++ {
		field2 := t2.Field(i)
		fields2[field2.Name] = field2.Type.String()
	}

	return reflect.DeepEqual(fields1, fields2)
}
