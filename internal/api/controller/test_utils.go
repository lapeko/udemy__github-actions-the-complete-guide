package controller

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
)

func MustMarshal(data interface{}) []byte {
	jsonData, err := json.Marshal(data)
	if err != nil {
		panic(err)
	}
	return jsonData
}

func BuildBody(data interface{}) io.ReadCloser {
	return io.NopCloser(bytes.NewReader(MustMarshal(data)))
}

func BuildRequest(data interface{}) *http.Request {
	return &http.Request{
		Body:   BuildBody(data),
		Header: http.Header{"Content-type": []string{"application/json"}},
	}
}
