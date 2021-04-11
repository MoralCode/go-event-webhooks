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
    client := &http.Client{
        // CheckRedirect: redirectPolicyFunc,
    }


    // https://golang.org/pkg/net/http/#Client.Post
    req, err := http.NewRequest(webhook.httpMethod,webhook.url, strings.NewReader(body))//"application/json",
    if err != nil {
        // handle error
        fmt.Println("Error")
    }
    _, err := client.Do(req)
    if err != nil {
        // handle error
        fmt.Println("Error")
    }
    // http.NewRequest("POST", url, strings.NewReader(form.Encode()))
}