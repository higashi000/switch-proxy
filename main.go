package main

import (
	"flag"
	"fmt"

	"github.com/higashi000/switch-proxy/setproxy"
)

func main() {
	flag.Parse()

	args := flag.Args()

	if len(args) < 2 {
		fmt.Println("Please add command line argument: `search address` `proxy address`")
		return
	}

	setproxy.GitProxy(args[1])
}
