package utils

import (
	"bytes"
	"encoding/json"
	"net/http"
)

type Error struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
}
type ResponseBool struct {
	Result bool `json:"result"`
}

func ResponseWithByte(v interface{}) []byte {
	reqBodyBytes := new(bytes.Buffer)
	_ = json.NewEncoder(reqBodyBytes).Encode(v)
	return reqBodyBytes.Bytes()
}

func ResponseWithJSON(response http.ResponseWriter, statusCode int, data interface{}) {
	result, _ := json.Marshal(data)
	response.Header().Set("Content-Type", "application/json")
	response.WriteHeader(statusCode)
	_, _ = response.Write(result)
}

func ResponseErr(w http.ResponseWriter, statusCode int) {
	jData, err := json.Marshal(Error{
		Status:  statusCode,
		Message: http.StatusText(statusCode),
	})
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(jData)
}

func ResponseOk(w http.ResponseWriter, data interface{}) {
	if data == nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	jData, err := json.Marshal(data)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(jData)
}

func ResponseWithErrByte(statusCode int) []byte {
	jData, err := json.Marshal(Error{
		Status:  statusCode,
		Message: http.StatusText(statusCode),
	})
	if err != nil {
		return nil
	}
	return jData
}
func ResponseOkByte(data interface{}) []byte {
	if data == nil {
		return ResponseWithErrByte(http.StatusInternalServerError)
	}

	jData, err := json.Marshal(data)
	if err != nil {
		return ResponseWithErrByte(http.StatusInternalServerError)
		return nil
	}
	return jData

}
