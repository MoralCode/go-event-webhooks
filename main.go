package main

import "fmt"
import "net/http"
import "strings"
import "errors"

import "github.com/MoralCode/go-event-webhooks/models"

type Registry map[string][]models.Webhook

var activeWebhooks Registry


func main() {

    /* create a map*/
    activeWebhooks = make(Registry)

    // fmt.Println(sayHi("Marco"))
    webhook := models.Webhook{
        "https://webhook.site/57663b0a-12b8-4f6d-a875-c38d30803561",
        "POST",
    }
    registerWebhook(activeWebhooks, "test", webhook)

    webhook2 := models.Webhook{
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


func sendWebhook(webhook models.Webhook, body string) {
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

func registerWebhook(registry Registry, eventId string, webhook models.Webhook) {
    values, ok := registry[eventId]   
   /* if ok is true, entry is present otherwise entry is absent*/
   if (ok) {
       if (webhook != models.Webhook{} && findIndexInList(values, webhook) == -1) {
            registry[eventId] = append(values, webhook)
       }
   } else {
        // before the loop
        output := []models.Webhook{}
        output = append(output, webhook)
        registry[eventId] = output
   }
}

func deregisterWebhook(registry Registry, eventId string, webhook models.Webhook) (error) {

    index := findIndexInList(registry[eventId], webhook)

    if index == -1 {
        return errors.New("provided webhook is not present in the registry for the given event ID")
    }

    newlist, err := remove(registry[eventId], index)
    if err != nil {
        return err
    }

    registry[eventId] = newlist
    return nil
}

// inspired by: https://stackoverflow.com/a/37335777/
func remove(list []models.Webhook, index int) ([]models.Webhook, error) {
    if index < 0 {
        return list, errors.New("negative indices are not allowed")
    }
    
    if (index >= len(list)) {
        return list, errors.New("provided index exceeds the bounds of the provided list ")
    }

    if (index == len(list)-1) {
        //if the item to remove is the last one, just return all but the last item
        return list[:len(list)-1], nil
    }

    //otherwise replace the item to be removed with the last item

    list[index] = list[len(list)-1]
    // We do not need to put s[i] at the end, as it will be discarded anyway
    return list[:len(list)-1], nil
}

func findIndexInList(list []models.Webhook, webhook models.Webhook) (int) {
    for i, n := range list {
        if webhook == n {
            return i
        }
    }
    return -1
}

func findIndexInRegistry(registry Registry, webhook models.Webhook) (string, int) {
    for key, _ := range registry {
        index := findIndexInList(registry[key], webhook)

        if index != -1 {
            return key, index
        }
    }
    return "", -1
}