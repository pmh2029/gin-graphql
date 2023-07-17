package utils

import "github.com/mitchellh/mapstructure"

func MapToStruct(
	input map[string]interface{},
	output interface{},
) error {
	decoder, err := mapstructure.NewDecoder(&mapstructure.DecoderConfig{
		TagName:     "mapstructure",
		ErrorUnused: true,
		ZeroFields:  true,
		Result:      output,
	})
	if err != nil {
		return err
	}

	return decoder.Decode(input)
}

func IsEmptyMap(input map[string]interface{}) bool {
	for _, value := range input {
		if nestedMap, ok := value.(map[string]interface{}); ok {
			if !IsEmptyMap(nestedMap) {
				return false
			}
		} else if value != nil {
			return false
		}
	}
	return true
}
