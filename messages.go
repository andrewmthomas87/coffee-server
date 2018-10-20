package main

import (
	"bytes"
	"encoding/json"
)

type authStatus struct {
	Auth bool
}

func formatMessage(messageType string, data interface{}) ([]byte, error) {
	json, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}

	return bytes.Join([][]byte{[]byte(messageType), json}, newline), nil
}
