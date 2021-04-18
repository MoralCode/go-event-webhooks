package main

import "testing"
import "fmt"
import "io/ioutil"
import "github.com/MoralCode/go-event-webhooks/models"
import "github.com/MoralCode/go-event-webhooks/helpers/tests"


func TestSendWebhook(t *testing.T) {
	t.Run("Successful Mirrored Response", func(t *testing.T) {

		client := &helpers.MockClient{}
		
		helpers.SetupMirrorResponse(client)

		testWebhook := models.Webhook{
			"https://example.com",
			"POST",
		}

		bodystr := "test"

		response, err := SendWebhook(client, testWebhook, bodystr)
		if err != nil {
			fmt.Println("an error occurred")
		}

		bytes, err := ioutil.ReadAll(response.Body)
		if err != nil {
			fmt.Println("an error occurred converting body to bytes")
		}
		if string(bytes) != bodystr {
			t.Errorf("Body changed between receiving response and returning it")
		}
	})

	t.Run("Error Response", func(t *testing.T) {
		client := &helpers.MockClient{}
		helpers.SetupServerError(client)

		testWebhook := models.Webhook{
			"https://example.com",
			"POST",
		}

		bodystr := "test"

		_, err := SendWebhook(client, testWebhook, bodystr)
		if err == nil {
			t.Errorf("Server Error was not correctly returned")
		}

	})

}