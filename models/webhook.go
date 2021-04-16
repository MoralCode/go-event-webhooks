package models

type Webhook struct {
    Url string         `json:"url"`
    HttpMethod string  `json:"httpMethod"`
}

type Webhooks []Webhook