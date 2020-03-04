package main

import (
	"flag"
	"fmt"
	"plugin"
)

func main() {
	flag.Parse()
	p, err := plugin.Open(flag.Arg(0))
	if err != nil {
		panic(err)
	}
	fp, err := p.Lookup("Add")
	if err != nil {
		panic(err)
	}
	f, ok := fp.(func(int, int) int)
	if !ok {
		panic("type mismatch")
	}
	fmt.Println(f(2, 2))
}
