package main

import (
	"fmt"
	"os"

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

	a := int(args[1])
	b := int(args[2])
	switch args[0] {
	case 'm':
		fmt.Println("%d*%d=%d", a, b, mathlib.Mult(a, b))
	case 'a':
		fmt.Println("%d+%d=%d", a, b, mathlib.Add(a, b))
	case 'p':
		fmt.Println("%d^%d=%d", a, b, mymath.Pow(a, b))
	default:
		fmt.Println("ERROR: invalid args")
	}
}