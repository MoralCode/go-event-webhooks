package main

import "fmt"
import "net/http"
import "strings"
import "errors"

import "github.com/MoralCode/go-event-webhooks/models"
import "github.com/MoralCode/go-event-webhooks/registries"

type HTTPClient interface {
	Do(req *http.Request) (*http.Response, error)
}

func configureClient() HTTPClient {
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
