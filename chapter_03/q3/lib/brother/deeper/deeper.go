package deeper

import (
    "os"
    in "github.com/lovelock/core-go/chapter_03/q3/lib/internal"
)

func Hello(name string) {
    in.Hello(os.Stdout, name)
}

