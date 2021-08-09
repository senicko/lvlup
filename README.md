# LvlUp client for Golang

Golang client for LvlUp api v4

`WARNING:` This package is under development. Not all features are implemented and breaking changes may occur.

## Project status

Current codebase can be treated as a experiment which can be something different within an hour from now. `Project will become stable with version 1.0. Until then be aware of the braking changes`

This does not mean that you shouldn't use it. Any feedback is welcome! If you have any ideas on how to make this library better, or you see potential issues with sourcecode create an issue with a proper flag.

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
