## go-seapi

[![GoDoc](https://godoc.org/github.com/nathan-osman/go-seapi?status.svg)](https://godoc.org/github.com/nathan-osman/go-seapi)
[![MIT License](http://img.shields.io/badge/license-MIT-9370d8.svg?style=flat)](http://opensource.org/licenses/MIT)

This package provides a simple way to access the Stack Exchange API.

### Example

This example fetches a list of recent questions on Stack Overflow and displays their titles:

    package main

    import (
        "log"
        "fmt"

        "github.com/nathan-osman/go-seapi"
    )

    func main() {
        v, err := seapi.NewRequest("/questions").Site("stackoverflow").Do()
        if err != nil {
            log.Fatal(err)
        }
        for _, i := range v.List("items") {
            fmt.Println(i.String("title"))
        }
    }
