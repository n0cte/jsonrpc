# jsonrpc
JSON-RPC 2.0 server and client pure implementation

## Server usage

```golang
package main

import (
  "net/http"

  "github.com/n0cte/jsonrpc"
)

type MathParams struct {
  A int64
  B int64
}

func main() {
  srv := jsonrpc.NewServer(map[string]jsonrpc.Handler{
    "sum": jsonrpc.TypedHandlerFunc(func(params MathParams) (int64, error) {
      return params.A + params.B, nil
    }),
  })
  srv.HandleFunc("mul", jsonrpc.TypedHandlerFunc(func(params MathParams) (int64, error) {
    return params.A * params.B, nil
  }))
  http.ListenAndServe(":8080", srv)
}
```

## Client usage
```golang
package main

import (
  "net/http"

  "github.com/n0cte/jsonrpc"
)

type MathParams struct {
  A int64
  B int64
}

func main() {
  client := jsonrpc.NewClient("127.0.0.1:8080", *http.DefaultClient)
  jsonrpc.TypedCall[MathParams, int64](client, "sum", MathParams{A: 0, B: 0})
}
```