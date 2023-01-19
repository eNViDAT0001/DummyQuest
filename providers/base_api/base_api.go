package base_api

import (
	"DummyQuest/providers/base_api/convert"
	"fmt"
	"net/http"
)

type BaseAPI struct {
	Header  map[string]string
	Url     string
	Query   string
	Payload *interface{}
	Client  *http.Client
}

func (api *BaseAPI) Request(method string) (*http.Response, error) {
	if len(api.Query) > 0 {
		api.Url = fmt.Sprintf("%s?%s", api.Url, api.Query)
	}

	body, err := convert.MapToBufferBody(api.Payload)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest(method, api.Url, body)
	if err != nil {
		return nil, err
	}

	for k, v := range api.Header {
		req.Header.Set(k, v)
	}

	resp, err := api.Client.Do(req)
	if err != nil {
		return nil, err
	}

	return resp, nil
}
