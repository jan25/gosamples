package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/jan25/gosamples/custom"
	"github.com/jan25/gosamples/internal"
)

func main() {
	args := os.Args[1:]
	if len(args) < 3 {
		fmt.Println(`usage:
		$ gosample [m|a|p] 3 4`)
		return
	}

	a32, _ := strconv.ParseInt(args[1], 10, 32)
	b32, _ := strconv.ParseInt(args[2], 10, 32)
	a := int(a32)
	b := int(b32)
	switch args[0] {
	case "m":
		fmt.Printf("%d*%d=%d", a, b, internal.Mult(a, b))
	case "a":
		fmt.Printf("%d+%d=%d", a, b, internal.Add(a, b))
	case "p":
		fmt.Printf("%d^%d=%d", a, b, custom.Pow(a, b))
	default:
		fmt.Printf("ERROR: invalid args")
	}
}
