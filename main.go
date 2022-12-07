package main

import (
	"fmt"
	"time"

	"github.com/leomirandadev/advance-golang-theory/fibonacci"
)

func main() {
	now := time.Now()
	fibonacci.RecursionWithoutMemo(30)
	fmt.Println(time.Since(now))
}

// go build -gcflags '-m -m' main.go
