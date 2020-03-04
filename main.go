package main

import (
	"flag"
	"fmt"
	"plugin"
)

func Double(a int) int {
	return 2 * a
}

func Triple(a int) int {
	return 3 * a
}

func main() {
	flag.Parse()
	p, err := plugin.Open(flag.Arg(0))
	if err != nil {
		panic(err)
	}
	fp, err := p.Lookup("AddDouble")
	if err != nil {
		panic(err)
	}
	f, ok := fp.(func(int, int) int)
	if !ok {
		panic("type mismatch")
	}
	fmt.Println(f(2, 2))
}
