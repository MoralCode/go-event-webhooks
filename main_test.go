package main

import "testing"

func TestRegisterWebhook(t *testing.T) {


    testWebhook2 := Webhook{
        "https://example.com/2",
        "POST",
    }

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
    t.Run("Duplicate Webhook", func(t *testing.T) {
    
        registerWebhook(activeWebhooks, "test", testWebhook2)

        if len(activeWebhooks) != 1 {
            t.Errorf("incorrect modifications to event identifers are present")
        }

        if len(activeWebhooks["test"]) != 2 {
            t.Errorf("incorrect number of webhooks present for event id \"test\"")
        }
    })
}


func TestFindIndexInList(t *testing.T) {

    testWebhook3 := Webhook{
        "https://example.com/test",
        "POST",
    }

    t.Run("Find in Empty List", func(t *testing.T) {
        list := []Webhook{}

        result := findIndexInList(list, testWebhook3)
        expected := -1

        if result != expected {
            t.Errorf("incorrect index from searching in an empty list")
        }
    })

    t.Run("Find in list", func(t *testing.T) {
        list2 := []Webhook{testWebhook3}

        result := findIndexInList(list2, testWebhook3)
        expected := 0

        if result != expected {
            t.Errorf("incorrect index from searching in list")

        }
    })

}


func TestFindIndexInRegistry(t *testing.T) {
    testWebhook := Webhook{
        "https://example.com",
        "POST",
    }

    testWebhook2 := Webhook{
        "https://example.com/2",
        "POST",
    }

    testWebhook3 := Webhook{
        "https://example.com/test",
        "POST",
    }

    testWebhook4 := Webhook{
        "https://example.com/test4",
        "POST",
    }
    /* create a map*/
    activeWebhooks := make(Registry)

    activeWebhooks["test"] = []Webhook{testWebhook, testWebhook2}
    activeWebhooks["test2"] = []Webhook{testWebhook3}


    t.Run("Find Unique Webhook", func(t *testing.T) {
        result_key, result_index := findIndexInRegistry(activeWebhooks, testWebhook4)
        expected_index := -1
        expected_key := ""

        if result_key != expected_key {
            t.Errorf("incorrect key from searching in registry")
        }

        if result_index != expected_index {
            t.Errorf("incorrect index from searching in registry")
        }
    })

    t.Run("Find Webhook in registry", func(t *testing.T) {
        result_key, result_index := findIndexInRegistry(activeWebhooks, testWebhook3)
        expected_index := 0
        expected_key := "test2"

        if result_key != expected_key {
            t.Errorf("incorrect key from searching in registry")
        }

        if result_index != expected_index {
            t.Errorf("incorrect index from searching in registry")
        }
    })

}