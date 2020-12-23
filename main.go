package main

import (
	"flag"
	"fmt"
	"log"
	"strings"

	"github.com/higashi000/switch-proxy/loadresolv"
	"github.com/higashi000/switch-proxy/setproxy"
)

func main() {
	flag.Parse()

	args := flag.Args()

	if len(args) < 2 {
		fmt.Println("Please add command line argument: `search address` `proxy address`")
		return
	}

	search, err := loadresolv.LoadResolv("/etc/resolv.conf")
	if err != nil {
		log.Println(err)
		return
	}

	if strings.Index(search, args[0]) != -1 {
		setproxy.SetProxy(args[1])
	}
}
