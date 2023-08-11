package jsonrpc

import "net/http"

func ExampleServer() {
	type MathParams struct {
		A int64
		B int64
	}

	srv := NewServer(map[string]Handler{
		"sum": TypedHandlerFunc(func(params MathParams) (int64, error) {
			return params.A + params.B, nil
		}),
	})

	srv.HandleFunc("mul", TypedHandlerFunc(func(params MathParams) (int64, error) {
		return params.A * params.B, nil
	}))

	http.ListenAndServe(":8080", srv)
}

func ExampleClient() {
	type MathParams struct {
		A int64
		B int64
	}
	client := NewClient("127.0.0.1:8080", *http.DefaultClient)

	TypedCall[MathParams, int64](client, "sum", MathParams{A: 0, B: 0})
}
