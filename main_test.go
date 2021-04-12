package main

import "testing"

func TestRegisterWebhook(t *testing.T) {

    /* create a map*/
    activeWebhooks := make(Registry)

    if len(activeWebhooks) != 0{
        t.Errorf("Registry does not begin in an empty state")
    }


    t.Run("CreatesCategory", func(t *testing.T) {
        testWebhook := Webhook{
            "https://example.com",
            "POST",
        }

        registerWebhook(activeWebhooks, "test", testWebhook)

        if len(activeWebhooks) != 1 {
            t.Errorf("Webhook event was not added correctly")
        }

    })
    t.Run("Add To Existing Category", func(t *testing.T) {
        testWebhook2 := Webhook{
            "https://example.com/2",
            "POST",
        }

        registerWebhook(activeWebhooks, "test", testWebhook2)

        if len(activeWebhooks) != 1 {
            t.Errorf("incorrect modifications to event identifers are present")
        }

        if len(activeWebhooks["test"]) != 2 {
            t.Errorf("incorrect number of webhooks present for event id \"test\"")
        }
    })
    t.Run("Empty Webhook", func(t *testing.T) {
        emptyWebhook := Webhook{}

        registerWebhook(activeWebhooks, "test", emptyWebhook)

        if len(activeWebhooks) != 1 {
            t.Errorf("incorrect modifications to event identifers are present")
        }

        if len(activeWebhooks["test"]) != 2 {
            t.Errorf("incorrect number of webhooks present for event id \"test\"")
        }
    })
}