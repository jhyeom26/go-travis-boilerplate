package main

import (
	"fmt"
	"os"
	"strconv"

	mymath "github.com/jhyeom26/go-travis-boilerplate/mymath"
)

func main() {
	if len(os.Args) != 3 {
		fmt.Println("Usage:", os.Args[0], "1", "2")
		return
	}
	op1, err := strconv.Atoi(os.Args[1])
	if err != nil {
		return
	}
	op2, err := strconv.Atoi(os.Args[2])
	if err != nil {
		return
	}
	result := mymath.AddInt(op1, op2)

	fmt.Printf("AddInt(%v, %v) : %v\n", op1, op2, result)
}
