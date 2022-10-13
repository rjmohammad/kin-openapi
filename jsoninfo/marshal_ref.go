package jsoninfo

import (
	"encoding/json"
)

func MarshalRef(value string, otherwise interface{}) ([]byte, error) {
	return json.Marshal(otherwise)
}

func UnmarshalRef(data []byte, destRef *string, destOtherwise interface{}) error {
	refProps := &refProps{}
	if err := json.Unmarshal(data, refProps); err == nil {
		ref := refProps.Ref
		if ref != "" {
			*destRef = ref
			return nil
		}
	}
	return json.Unmarshal(data, destOtherwise)
}

type refProps struct {
	Ref string `json:"$ref,omitempty" yaml:"$ref,omitempty"`
}
