package main

import (
	"flag"
	"fmt"
	"os"
	"jam/bazaar"
)

// generated by import vendors
import (
	_"jam/vendors/client/zygote"
)

func main() {
	flag.Parse()
	if !flag.Parsed() { os.Exit(1) }

	bazaar.Connect()

	e := bazaar.Create()
	if e != nil {fmt.Println(e.Error())}
}

