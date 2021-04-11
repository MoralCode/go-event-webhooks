package main

import "testing"

func TestSayHi(t *testing.T) {
    expected := "Hi Marco"
    greeting := sayHi("Marco")
    if greeting != expected {
        t.Errorf("Greeting was incorrect, got: '%s', want: '%s'", greeting, expected)
    }
}