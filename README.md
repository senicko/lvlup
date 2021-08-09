# LvlUp client for Golang

Golang client for LvlUp api v4

`WARNING:` This package is under development. Not all features are implemented and major changes may occur.

## Quick Start

```go
import (
  "net/http"
  "fmt"
)

func main() {
  httpClient := &http.Client{}
  client := lvlup.NewLvlCLient("<api_key>", httpClient)

  result, err := client.CreatePayment(
    "24.99",
    lvlup.WithRedirectUrl("<redirect_url>"),
    lvlup.WithWebhookUrl("<webhook_url>")
  )

  if err != nil {
    panic(err)
  }

  fmt.Println(result)
}
```
