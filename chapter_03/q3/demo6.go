package main

import (
    "flag"
    // "github.com/lovelock/core-go/chapter_03/q3/lib"
    // brother "github.com/lovelock/core-go/chapter_03/q3/lib/brother"
    deeper "github.com/lovelock/core-go/chapter_03/q3/lib/brother/deeper"
    // "github.com/lovelock/core-go/chapter_03/q3/lib/internal"
)

var name string

func init() {
    flag.StringVar(&name, "name", "everyone", "The greeting object.")
}

func main() {
    flag.Parse()
    deeper.Hello(name)
}
