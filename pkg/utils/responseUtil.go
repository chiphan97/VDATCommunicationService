package utils

import (
	"bytes"
	"encoding/json"
)

func ResponseWithByte(v interface{}) []byte {
	reqBodyBytes := new(bytes.Buffer)
	_ = json.NewEncoder(reqBodyBytes).Encode(v)
	return reqBodyBytes.Bytes()
}
