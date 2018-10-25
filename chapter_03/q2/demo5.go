package main

import (
    "flag"
    "github.com/lovelock/core-go/chapter_03/q2/lib"
)

var name string

func init() {
    flag.StringVar(&name, "name", "everyone", "The greeting object.")
}

func main() {
    flag.Parse()
    lib5.Hello(name)
}
