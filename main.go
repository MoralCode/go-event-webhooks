package main

import "fmt"
import "net/http"
import "io"

type Envelope struct {
    Message string
}

func main() {
    resp, err := http.Get("https://webhook.site/57663b0a-12b8-4f6d-a875-c38d30803561")
    if err != nil {
        // handle error
        fmt.Println("Error")
    }
    defer resp.Body.Close()
    body, err := io.ReadAll(resp.Body)
    if err != nil {
        // handle error
        fmt.Println("Error")
    }
    fmt.Println(body)
}