# go-event-webhooks

a go module to simplify the process of adding webhook-sending functionality to a project.


## Installation
Run `go get github.com/MoralCode/go-event-webhooks` and add an `import "go-event-webhooks"` to your file

## Usage

Below is an example of basic usage

```go
import (
  "github.com/MoralCode/go-event-webhooks/models"
  "github.com/MoralCode/go-event-webhooks/registries"
  )

// Create a registry to store your webhooks:
registry := registries.CreateNewJSONRegistry("registry.json")

// Create a webhook to add to your registry
myWebhook, err := models.CreateWebhook("http://my.website", "GET")
if err != nil {
  //an error occured creating the webhook
}

//Add your webhook to the registry for a particular event id
registry.AddToEvent(myWebhook, "file-uploaded")

```

When an event happens in your program that you would like to send a webhook for, use the `TriggerEvent` method to send a request to all webhooks registered to that event id.
```go
err := TriggerEvent(registry, "test", "this is a test")
if err != nil {
    fmt.Println(err)
}
```

## Key Concepts

The two key concepts used by this library are "Registry" and "Event". These are outlined below.

An **Event** is something that happens within your program that you would like to be able to send webhooks for. Examples may include things like new content being created. Events are represented as a string identifier.

A **Registry** refers to a data structure that stores Webhooks and associates, or "registers" them to an event identifier. This allows your app to send webhooks for a given event simply by using the event id and providing a webhook request body.

## Registry Types
Since registries are just simple data storage, it's possible to write a registry for any type of data storage backend.

### MapRegistry
This registry stores webhooks using a map structure where each event id stores a list of webhooks. This is a simple in-memory storage solution and is the most basic type of registry.

### JSONRegistry
The JSONRegistry builds on top of a MapRegistry to allow a registry to be stored and re-created from a JSON file.

### Custom Registries
Of course, since Registries are just structures that implement the `Registry` interface, you can also create your own registries. This allows for the use of storage backends that are not yet included in this library, such as databases and s3 storage as well as to other file formats like XML or CSV. 

If you create a registry for a common data storage backend that is not currently implemented here, feel free to send in a Pull Request!

## Development

If you have a main function defined in `main.go`, run `go run main.go` to run it.

You can build this program into an executable by running `go build .`

To run the unit tests, use the command `go test ./...`. to run tests and generate a coverage report, use `go test -coverprofile=coverage.out ./...`. To open this report in a browser HTML run `go tool cover -html=coverage.out`.
