package server

import (
	"bytes"
	"encoding/json"
)

var (
	_ json.Unmarshaler = (*Request)(nil)
)

type (
	RequestParams struct {
		JsonRPC string          `json:"jsonrpc"`
		ID      json.RawMessage `json:"id"`
		Method  string          `json:"method"`
		Params  json.RawMessage `json:"params,omitempty"`
	}

	Request struct {
		single bool
		params []RequestParams
	}
)

func (req *Request) UnmarshalJSON(data []byte) error {
	data = bytes.TrimSpace(data)
	req.single = !bytes.HasPrefix(data, []byte("["))
	req.params = make([]RequestParams, 0)

	if req.single {
		var prm RequestParams
		if err := json.Unmarshal(data, &prm); err != nil {
			return nil
		}
		req.params = append(req.params, prm)
		return nil
	}

	return json.Unmarshal(data, &req.params)
}
