# jsonrpc
JSON-RPC 2.0 server and client pure implementation

# Server usage

```golang
package main

import (
	"net/http"

	"github.com/n0cte/jsonrpc/server"
)

type SumParams struct {
	A int64
	B int64
}

func main() {
	srv := server.New(map[string]server.Handler{
		"sum": server.TypedHandlerFunc(func(params SumParams) (int64, error) {
			return params.A + params.B, nil
		}),
	})
	http.ListenAndServe(":8080", srv)
}
```