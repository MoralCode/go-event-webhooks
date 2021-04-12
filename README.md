# go-event-webhooks

a go module to simplify the process of adding webhook-sending functionality to a project.


 
 webhook tester 
https://webhook.site/#!/57663b0a-12b8-4f6d-a875-c38d30803561
https://marcofranssen.nl/start-on-your-first-golang-project/

## run program
`go run main.go`


## tests
`go test ./...`

run tests and gen coverage file
`go test -coverage -coverprofile=coverage.out ./...
`

##  build executable:
`go build .`


## How does it work?

The go-event-webhooks module allows apps to easily create "events" identified by a string. These event identifiers allow webhooks to be associated with that event, allowing the application to easily trigger webhooks with an event ID and webhook body.  

## TODO
- [ ] write unit tests
- [ ] find a way 
- [ ] Support custom headers
- [ ] support custom cookies
- [ ] add easy setup methods to handle sending webhooks in common formats like discord and github