package main

import "fmt"

type Envelope struct {
    Message string
}

func main() {
    fmt.Println(sayHi("Marco"))
}

func sayHi(person string) string {
    return fmt.Sprintf("Hi %s", person)
}