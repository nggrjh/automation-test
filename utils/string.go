package utils

import "encoding/json"

func ToByte(v any) []byte {
	b, _ := json.Marshal(v)
	return b
}

func ToString(v any) string {
	return string(ToByte(v))
}
