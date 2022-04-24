package util

import (
	"encoding/json"
)

func StructToByte(value interface{}) ([]byte, error) {
	return json.Marshal(value)
}

func ByteToStruct(data []byte, dataStruct interface{}) error {
	return json.Unmarshal(data, &dataStruct)
}
