package main

import (
	"fmt"
	"jam/bazaar"
)

// generated by import vendors
import (
	_"jam/vendors/client/zygote"
)

func main() {
	e := bazaar.Destroy()
	if e != nil { fmt.Println(e.Error()) }
}

