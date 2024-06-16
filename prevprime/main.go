package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	num, err := strconv.Atoi(os.Args[1])
	if err != nil {
		fmt.Println(err)
		os.Exit(0)
	}
	fmt.Println(prevprime(num))
}

func findPrime(n int) (prime bool) {
	if n < 2 {
		prime = false
	} else if n < 4 {
		prime = true
	} else if n%2 == 0 {
		prime = false
	} else if n == 5 {
		prime = true
	} else if n%5 == 0 && n != 5 {
		prime = false
	} else if n%3 == 0 {
		prime = false
	} else {
		prime = true
	}
	return
}

func prevprime(prime int) (prev int) {
	for i := prime; i >= prime-10; i-- {
		if !findPrime(i) {
			prev += 0
		} else {
			prev += i
			break
		}
	}
	return
}
