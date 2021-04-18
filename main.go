package main

import "fmt"
import "net/http"
import "strings"
import "errors"

import "github.com/MoralCode/go-event-webhooks/models"
import "github.com/MoralCode/go-event-webhooks/registries"


var activeWebhooks registries.MapRegistry

type HTTPClient interface {
	Do(req *http.Request) (*http.Response, error)
}


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

func configureClient() {
    return &http.Client{
        // CheckRedirect: redirectPolicyFunc,
    }
}

func TriggerEvent(registry registries.Registry, eventId string, body string) (error) {

    eventWebhooks := registry.GetHooksForEvent(eventId)
    if (eventWebhooks != nil) {
        client := configureClient()
        for _, hook := range eventWebhooks {
            fmt.Println(hook)
            SendWebhook(client, hook, body)
        }
        return nil
    } else {
        //error: no webhooks are registered at this id
        return errors.New("no webhooks are registered for the eventId " + eventId)
    }
}

func SendWebhook(client HTTPClient, webhook models.Webhook, body string) (*http.Response, error) {

    // https://golang.org/pkg/net/http/#Client.Post
    req, err := http.NewRequest(webhook.HttpMethod, webhook.Url, strings.NewReader(body))//"application/json",
    if err != nil {
        // handle error
        fmt.Println(err)
        return &http.Response{}, err
    }
    response, err := client.Do(req)
    if err != nil {
        // handle error
        fmt.Println(err)
        return &http.Response{}, err
    } else {
        defer response.Body.Close()
    }

    return response, nil
    // http.NewRequest("POST", url, strings.NewReader(form.Encode()))
}
