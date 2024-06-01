package util

import (
	"encoding/json"
)

func MarshalJson(v interface{}) string {
	data, err := json.Marshal(v)
	if err != nil {
		Log("util.MarshalJson Error, v: %v", v)
		return ""
	}
	return string(data)
}

func MarshalJsonRetPtr(v interface{}) *string {
	data, err := json.MarshalIndent(v, "", "  ")
	if err != nil {
		Log("util.MarshalJsonRetPtr Error, v: %v", v)
		return new(string)
	}

	str := string(data)
	return &str
}

func UnmarshalJson[T any](data string) T {
	var x T
	json.Unmarshal([]byte(data), &x)
	return x
}

func UnmarshalJsonRetPrt[T any](data string) *T {
	var x T
	json.Unmarshal([]byte(data), &x)
	return &x
}
