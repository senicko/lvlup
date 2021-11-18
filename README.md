# LvlUp client for Golang

Golang client for [LvlUp api v4](https://api.lvlup.pro/v4/redoc)

## Quick Start

```go
package main

import (
  "net/http"
  "fmt"
  
  "github.com/senicko/lvlup"
)

func main() {
  httpClient := &http.Client{}
  client := lvlup.NewLvlCLient("<api_key>", httpClient)

  result, err := client.CreatePayment(
    "24.99",
    lvlup.WithRedirectUrl("<redirect_url>"),
    lvlup.WithWebhookUrl("<webhook_url>"),
  )

  if err != nil {
    panic(err)
  }

  fmt.Println(result)
}
```
