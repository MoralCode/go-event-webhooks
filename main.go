package main

import "fmt"
import "net/http"
import "strings"
import "errors"

import "github.com/MoralCode/go-event-webhooks/models"
import "github.com/MoralCode/go-event-webhooks/registries"


var activeWebhooks registries.MapRegistry


func main() {

    /* create a map*/
    activeWebhooks = registries.CreateMapRegistry()

    // fmt.Println(sayHi("Marco"))
    webhook := models.Webhook{
        "https://webhook.site/57663b0a-12b8-4f6d-a875-c38d30803561",
        "POST",
    }
    activeWebhooks.AddToEvent(webhook, "test")

    webhook2 := models.Webhook{
        "https://webhook.site/57663b0a-12b8-4f6d-a875-c38d30803561",
        "GET",
    }
    activeWebhooks.AddToEvent(webhook2, "test")

    err := TriggerEvent(activeWebhooks, "test", "this is a test")
    if err != nil {
        fmt.Println(err)
    }
}

func TriggerEvent(registry registries.Registry, eventId string, body string) (error) {

    eventWebhooks := registry.GetHooksForEvent(eventId)
    if (eventWebhooks != nil) {
        for _, hook := range eventWebhooks {
            fmt.Println(hook)
            SendWebhook(hook, body)
        }
        return nil
    } else {
        //error: no webhooks are registered at this id
        return errors.New("no webhooks are registered for the eventId " + eventId)
    }
}


func SendWebhook(webhook models.Webhook, body string) {
    client := &http.Client{
        // CheckRedirect: redirectPolicyFunc,
    }


    // https://golang.org/pkg/net/http/#Client.Post
    req, err := http.NewRequest(webhook.HttpMethod, webhook.Url, strings.NewReader(body))//"application/json",
    if err != nil {
        // handle error
        fmt.Println(err)
    }
    defer resp.Body.Close()
    _, err = client.Do(req)
    if err != nil {
        // handle error
        fmt.Println(err)
    }
    // http.NewRequest("POST", url, strings.NewReader(form.Encode()))
}
