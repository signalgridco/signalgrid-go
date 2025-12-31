# Signalgrid Go Client

Official Go client for the Signalgrid notification API.

## Install

```bash
go get github.com/signalgridco/signalgrid-go
```
## Example
```go
package main

import (
	"fmt"
	"log"

	"github.com/signalgridco/signalgrid-go"
)

func main() {
	client, err := signalgrid.NewClient("YOUR_CLIENT_KEY")
	if err != nil {
		log.Fatal(err)
	}

	resp, err := client.Send(signalgrid.Message{
		Channel:  "CHANNEL_TOKEN",
		Type:     "info",
		Title:    "Hello from Go",
		Body:     "Sent from a Go app using go get.",
		Critical: false,
	})

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(resp)
}
```
## Response
```
{
  "ruuid":"09c2e7c68390458f9927c7564529d3599b906840",
  "text":"OK",
  "code":"200",
  "node":"dp02",
  "time":"380ms",
  "remaining":"99915"
}
```
