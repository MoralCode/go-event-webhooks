package registries


import "testing"
import "github.com/MoralCode/go-event-webhooks/models"
import "github.com/go-test/deep"


func TestMakeJsonRegistry(t *testing.T) {

    t.Run("Create New Registry", func(t *testing.T) {
        /* create a map*/
        jsonRegistry := CreateNewJSONRegistry("")
        expected := JSONRegistry{MapRegistry{}, ""}

        if diff := deep.Equal(jsonRegistry, expected); diff != nil {
            t.Error(diff)
        }

        if len(jsonRegistry.Registry) != 0 {
            t.Errorf("Registry does not begin in an empty state")
        }
    })

    t.Run("Create Registry from JSON String", func(t *testing.T) {
        
        testWebhook := models.Webhook{
            "https://example.com",
            "POST",
        }

        testWebhook2 := models.Webhook{
            "https://example.com/2",
            "POST",
        }
        mapreg := make(MapRegistry)
        mapreg["test"] = models.Webhooks{testWebhook, testWebhook2}

        data := "{\"test\":[{\"url\":\"https://example.com\",\"httpMethod\":\"POST\"},{\"url\":\"https://example.com/2\",\"httpMethod\":\"POST\"}]}"

        jsonRegistry := CreateJSONRegistryFromJSONData(data)
        expected := JSONRegistry{mapreg, ""}

        if diff := deep.Equal(jsonRegistry, expected); diff != nil {
            t.Error(diff)
        }

    })
}

func TestJsonRegisterWebhook(t *testing.T) {


    testWebhook2 := models.Webhook{
        "https://example.com/2",
        "POST",
    }

    /* create a map*/
    activeWebhooks := CreateNewJSONRegistry("testfile.json")

    t.Run("CreatesCategory", func(t *testing.T) {
        testWebhook := models.Webhook{
            "https://example.com",
            "POST",
        }

        activeWebhooks.AddToEvent(testWebhook, "test")

        if len(activeWebhooks.Registry) != 1 {
            t.Errorf("Webhook event was not added correctly")
        }

    })
    t.Run("Add To Existing Category", func(t *testing.T) {

        activeWebhooks.AddToEvent(testWebhook2, "test")

        if len(activeWebhooks.Registry) != 1 {
            t.Errorf("incorrect modifications to event identifers are present")
        }

        if len(activeWebhooks.GetHooksForEvent("test")) != 2 {
            t.Errorf("incorrect number of webhooks present for event id \"test\"")
        }
    })
    t.Run("Empty Webhook", func(t *testing.T) {
        emptyWebhook := models.Webhook{}

        activeWebhooks.AddToEvent(emptyWebhook, "test")

        if len(activeWebhooks.Registry) != 1 {
            t.Errorf("incorrect modifications to event identifers are present")
        }

        if len(activeWebhooks.GetHooksForEvent("test")) != 2 {
            t.Errorf("incorrect number of webhooks present for event id \"test\"")
        }
    })
    t.Run("Duplicate Webhook", func(t *testing.T) {
    
        activeWebhooks.AddToEvent(testWebhook2, "test")

        if len(activeWebhooks.Registry) != 1 {
            t.Errorf("incorrect modifications to event identifers are present")
        }

        if len(activeWebhooks.GetHooksForEvent("test")) != 2 {
            t.Errorf("incorrect number of webhooks present for event id \"test\"")
        }
    })
}


func TestJsonFindIndexInList(t *testing.T) {

    testWebhook3 := models.Webhook{
        "https://example.com/test",
        "POST",
    }

    t.Run("Find in Empty List", func(t *testing.T) {
        list := models.Webhooks{}

        result := list.FindIndexOf(testWebhook3)
        expected := -1

        if result != expected {
            t.Errorf("incorrect index from searching in an empty list")
        }
    })

    t.Run("Find in list", func(t *testing.T) {
        list2 := models.Webhooks{testWebhook3}

        result := list2.FindIndexOf(testWebhook3)
        expected := 0

        if result != expected {
            t.Errorf("incorrect index from searching in list")

        }
    })

}


func TestJsonFindIndexInRegistry(t *testing.T) {
    testWebhook := models.Webhook{
        "https://example.com",
        "POST",
    }

    testWebhook2 := models.Webhook{
        "https://example.com/2",
        "POST",
    }

    testWebhook3 := models.Webhook{
        "https://example.com/test",
        "POST",
    }

    testWebhook4 := models.Webhook{
        "https://example.com/test4",
        "POST",
    }
    /* create a map*/
    mapreg := make(MapRegistry)

    mapreg["test"] = models.Webhooks{testWebhook, testWebhook2}
    mapreg["test2"] = models.Webhooks{testWebhook3}

    activeWebhooks := CreateJSONRegistryFromMapRegistry(mapreg)


    t.Run("Find Unique Webhook", func(t *testing.T) {
        result_key, result_index := activeWebhooks.Find(testWebhook4)
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
        result_key, result_index := activeWebhooks.Find(testWebhook3)
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


func TestJsonRemove(t *testing.T) {

   testWebhook := models.Webhook{
        "https://example.com",
        "POST",
    }

    testWebhook2 := models.Webhook{
        "https://example.com/2",
        "POST",
    }
     /* create a map*/
    mapreg := make(MapRegistry)

    mapreg["test"] = models.Webhooks{testWebhook, testWebhook2}

    activeWebhooks := CreateJSONRegistryFromMapRegistry(mapreg)

    t.Run("Remove Negative index", func(t *testing.T) {

        _, err := activeWebhooks.GetHooksForEvent("test").RemoveIndex(-1)
    
        if err == nil {
            t.Errorf("does not handle negative indices")
        }
    })

    t.Run("Remove out of bounds index", func(t *testing.T) {

        _, err := activeWebhooks.GetHooksForEvent("test").RemoveIndex(len(activeWebhooks.GetHooksForEvent("test"))+2)
    
        if err == nil {
            t.Errorf("does not handle out of bounds indices")
        }

    })

    t.Run("Remove last item", func(t *testing.T) {
        result, err := activeWebhooks.GetHooksForEvent("test").RemoveIndex(len(activeWebhooks.GetHooksForEvent("test"))-1)
    
        if (err != nil || result[0] != testWebhook) {
            t.Errorf("does not handle the special case for removing the last item")
        }
    })

    t.Run("Remove another item from the list", func(t *testing.T) {
        result, err := activeWebhooks.GetHooksForEvent("test").RemoveIndex(0)
    
        if (err != nil || result[0] != testWebhook2) {
            t.Errorf("does not correctly remove items in the general case")
        }
    })
}

func TestJsonDeregister(t *testing.T) {

    testWebhook := models.Webhook{
        "https://example.com",
        "POST",
    }

    testWebhook2 := models.Webhook{
        "https://example.com/2",
        "POST",
    }

    testWebhook3 := models.Webhook{
        "https://example.com/test",
        "POST",
    }

    testWebhook4 := models.Webhook{
        "https://example.com/test4",
        "POST",
    }

    /* create a map*/
    mapreg := make(MapRegistry)

    mapreg["test"] = models.Webhooks{testWebhook, testWebhook2}
    mapreg["test2"] = models.Webhooks{testWebhook3, testWebhook4}

    activeWebhooks := CreateJSONRegistryFromMapRegistry(mapreg)


    t.Run("Deregisters Webhook from wrong event", func(t *testing.T) {
        err := activeWebhooks.RemoveFromEvent(testWebhook, "test2")

        if err == nil {
             t.Errorf("does not correctly check if webhook exists in the provided eventId before removing")
        }
    })

    t.Run("Deregisters Webhook", func(t *testing.T) {
        err := activeWebhooks.RemoveFromEvent(testWebhook, "test")

        if (err != nil) {
            t.Error(err)
        }

        if (activeWebhooks.GetHooksForEvent("test")[0] != testWebhook2) {
             t.Errorf("does not correctly remove webhook from registry")
        }
    })
}