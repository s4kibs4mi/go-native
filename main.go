package main

// #cgo amd64 386 CFLAGS: -DX86=1
// #cgo CFLAGS: -I.
// #cgo LDFLAGS: -L. -lhello
// #include "hello.h"
import "C"

import "fmt"

func isPrime(x C.int) bool {
	var r C.int = C.isPrime(x)
	if r == 1 {
		return true
	}
	return false
}

func main() {
	ok := isPrime(C.int(14))
	if ok {
		fmt.Printf("Prime\n")
	} else {
		fmt.Printf("Not prime\n")
	}
}
