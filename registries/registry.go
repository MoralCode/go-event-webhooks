package registries

import "github.com/MoralCode/go-event-webhooks/models"

type Registry interface {
    AddToEvent(webhook models.Webhook, eventId string) (error)
    RemoveFromEvent(webhook models.Webhook, eventId string) (error)
    GetHooksForEvent(eventId string) ([]models.Webhook)
    Find(webhook models.Webhook) (string, int)
    FindInEvent(eventId string, webhook models.Webhook) (int, error)
    ListEvents() ([]string)
}