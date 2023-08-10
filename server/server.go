package server

import (
	"encoding/json"
	"net/http"
)

var (
	_ http.Handler = (*Server)(nil)
	_ Handler      = (HandlerFunc)(nil)
)

type (
	HandlerFunc func(json.RawMessage) (json.RawMessage, error)

	Handler interface {
		Dispatch(json.RawMessage) (json.RawMessage, error)
	}

	Server struct {
		handlers map[string]Handler
	}
)

func New(handlers map[string]Handler) *Server {
	return &Server{handlers: handlers}
}

func (f HandlerFunc) Dispatch(data json.RawMessage) (json.RawMessage, error) {
	return f(data)
}

func (s Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	var (
		req Request
		res Response
	)
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		res = Response{
			single: true,
			params: []ResponseParams{{Error: &ErrParseError}},
		}
	} else {
		defer r.Body.Close()

		res = Response{
			single: req.single,
			params: make([]ResponseParams, len(req.params)),
		}
		for i := range req.params {
			res.params[i].ID = req.params[i].ID
			if req.params[i].JsonRPC != "2.0" {
				res.params[i].Error = &ErrInvalidRequest
			} else if handler, ok := s.handlers[req.params[i].Method]; !ok {
				res.params[i].Error = &ErrInvalidMethod
			} else if data, err := handler.Dispatch(req.params[i].Params); err != nil {
				e := ErrInternalError
				e.Data = err
				res.params[i].Error = &e
			} else {
				res.params[i].Result = data
			}
		}
	}

	json.NewEncoder(w).Encode(res)
}
