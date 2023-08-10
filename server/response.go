package server

import (
	"encoding/json"
)

var (
	_ json.Marshaler = (*Response)(nil)
)

type (
	ResponseParams struct {
		JsonRPC string          `json:"jsonrpc"`
		ID      json.RawMessage `json:"id"`
		Error   *Error          `json:"error,omitempty"`
		Result  json.RawMessage `json:"result,omitempty"`
	}

	Response struct {
		single bool
		params []ResponseParams
	}
)

func (res Response) MarshalJSON() ([]byte, error) {
	if res.single {
		return json.Marshal(res.params[0])
	}
	return json.Marshal(res.params)
}
