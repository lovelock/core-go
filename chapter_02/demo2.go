package main

import (
    "flag"
    "fmt"
    "os"
)

var name string
var cmdLine = flag.NewFlagSet("", flag.PanicOnError)

func init() {
    cmdLine.StringVar(&name, "name", "everyone", "The greeting object.") // 地址；名字；默认值；描述；把值存在name中
    // flag.StringVar(&name, "name", "everyone", "The greeting object.") // 地址；名字；默认值；描述；把值存在name中
    // flag.String("name", "everyone", "The greeting object") // 返回一个存放了值的地址
    // flag.CommandLine = flag.NewFlagSet("", flag.PanicOnError)
}

func main() {
    // flag.Usage = func() {
        // fmt.Fprintf(os.Stderr, "Usage of %s:\n", "question")
        // flag.PrintDefaults()
    // }
    // flag.Parse() // 真正解析命令参数
    cmdLine.Parse(os.Args[1:])

    fmt.Printf("Hello %s!\n", name)
}


