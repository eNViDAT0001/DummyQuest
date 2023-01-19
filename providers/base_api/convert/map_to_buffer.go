package convert

import (
	"bytes"
	"encoding/json"
)

func MapToBufferBody(input interface{}) (*bytes.Buffer, error) {
	jsBody, err := json.Marshal(input)
	if err != nil {
		return nil, err
	}
	body := bytes.NewBuffer(jsBody)
	return body, nil
}
