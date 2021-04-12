package main

import "fmt"
import "net/http"
import "strings"
import "errors"

type Webhook struct {
    url string;
    httpMethod string;
}

type Registry map[string][]Webhook

var activeWebhooks Registry


func main() {

    /* create a map*/
    activeWebhooks = make(Registry)

    // fmt.Println(sayHi("Marco"))
    webhook := Webhook{
        "https://webhook.site/57663b0a-12b8-4f6d-a875-c38d30803561",
        "POST",
    }
    registerWebhook(activeWebhooks, "test", webhook)

    webhook2 := Webhook{
        "https://webhook.site/57663b0a-12b8-4f6d-a875-c38d30803561",
        "GET",
    }
    registerWebhook(activeWebhooks, "test", webhook2)

    err := triggerWebhook(activeWebhooks, "test", "this is a test")
    if err != nil {
        fmt.Println(err)
    }
}

func triggerWebhook(registry Registry, eventId string, body string) (error) {

    eventWebhooks, ok := registry[eventId]
    if (ok) {
        for _, hook := range eventWebhooks {
            fmt.Println(hook)
            sendWebhook(hook, body)
        }
        return nil
    } else {
        //error: no webhooks are registered at this id
        return errors.New("no webhooks are registered for the eventId " + eventId)
    }
}


func sendWebhook(webhook Webhook, body string) {
    client := &http.Client{
        // CheckRedirect: redirectPolicyFunc,
    }


    // https://golang.org/pkg/net/http/#Client.Post
    req, err := http.NewRequest(webhook.httpMethod,webhook.url, strings.NewReader(body))//"application/json",
    if err != nil {
        // handle error
        fmt.Println(err)
    }
    _, err = client.Do(req)
    if err != nil {
        // handle error
        fmt.Println(err)
    }
    // http.NewRequest("POST", url, strings.NewReader(form.Encode()))
}

func registerWebhook(registry Registry, eventId string, webhook Webhook) {
    values, ok := registry[eventId]   
   /* if ok is true, entry is present otherwise entry is absent*/
   if (ok) {
       if (webhook != Webhook{}) {
            registry[eventId] = append(values, webhook)
       }
   } else {
        // before the loop
        output := []Webhook{}
        output = append(output, webhook)
        registry[eventId] = output
   }
}

func findIndexInList(list []Webhook, webhook Webhook) (int) {
    for i, n := range list {
        if webhook == n {
            return i
        }
    }
    return -1
}

func findIndexInRegistry(registry Registry, webhook Webhook) (string, int) {
    for key, _ := range registry {
        index := findIndexInList(registry[key], webhook)

        if index != -1 {
            return key, index
        }
    }
    return "", -1
}