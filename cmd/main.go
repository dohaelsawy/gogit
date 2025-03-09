package main

import (
	"flag"
	cli "github.com/dohaelsawy/gogit/internal"
)

var (
	_init bool
)

func main() {

	flag.BoolVar(&_init, "init", true, "initial command to print hello world")

	if _init {
		cli.CustomInit()
	}

}
