package main

import "fmt"
import "net/http"
import "strings"

type Webhook struct {
    url string;
    httpMethod string;
}

func main() {
    // fmt.Println(sayHi("Marco"))
    webhook := Webhook{
        "https://webhook.site/57663b0a-12b8-4f6d-a875-c38d30803561",
        "POST",
    }

    sendWebhook(webhook, "testttt")
}


func sendWebhook(webhook Webhook, body string) {

    // https://golang.org/pkg/net/http/#Client.Post
    _, err := http.Post(webhook.url,  "application/json", strings.NewReader(body))
    if err != nil {
        // handle error
        fmt.Println("Error")
    }
    // http.NewRequest("POST", url, strings.NewReader(form.Encode()))
}