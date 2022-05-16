package main

import (
    "fmt"
    "log"

    "example.com/greetings"
)

func main() {
    // Set properties of the predefined Logger, including
    // the log entry prefix and a flag to disable printing
    // the time, source file, and line number.
    log.SetPrefix("greetings: ")
    log.SetFlags(0)

    fmt.Println("Hello Sir, how was your last night :)")
    // A slice of names.
    names := []string{"Govardhan", "Abhishek", "Dheeraj", "Hemant", "Pragati", "Manoji", "Shaun Ji"}

    // Request a greeting message.
    //message, err := greetings.Hello("Govardhan")
    message, err := greetings.Hellos(names)
    // If an error was returned, print it to the console and
    // exit the program.
    if err != nil {
        log.Fatal(err)
    }

    // If no error was returned, print the returned message
    // to the console.
    for _, element := range message {
        fmt.Println(element)
    }
}