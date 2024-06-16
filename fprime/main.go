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
	fmt.Println(Fprime(num))
}

func Fprime(num int) (factors []int) {
	var prefactors []int
	for i := 1; i <= num; i++ {
		if num%i == 0 {
			prefactors = append(prefactors, i)
		}
	}
	for _, facts := range prefactors {
		if findPrime(facts) {
			factors = append(factors, facts)
		}
	}
	return
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
