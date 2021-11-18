# LvlUp client for Golang

Golang client for [LvlUp api v4](https://api.lvlup.pro/v4/redoc)

## Project status

Current codebase can be treated as a experiment which can be something different within an hour from now. `Project will become stable with version 1.0. Until then be aware of the braking changes`

This does not mean that you shouldn't use it. Any feedback is welcome! If you have any ideas on how to make this library better, or you see potential problem create an issue with a proper flag.

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
