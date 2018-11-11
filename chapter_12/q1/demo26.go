package main

import (
	"fmt"
	"github.com/pkg/errors"
	"github.com/gpmgo/gopm/modules/log"
)

type Printer func(contents string) (n int, err error)

func printToStd(contents string) (bytesNum int, err error) {
	return fmt.Println(contents)
}


type operate func(x, y int) int

func calculate(x, y int, op operate) (int, error){
	if op == nil {
		return 0, errors.New("invalid operation")
	}

	return op(x, y), nil
}

func main() {
	var p Printer
	p = printToStd
	p("something")

	//var op = func (x, y int) int {
	//	return x + y
	//}

	op := *new(operate)
	z, err := calculate(20, 30, op)
	if err != nil {
		log.Fatal("%s", err)
	}
	fmt.Println(z)
}