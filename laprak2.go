package main

import "fmt"

func main() {
	var b int
	fmt.Scan(&b)

	if b == 1 {
		fmt.Println("BUKAN PRIMA")
		return
	}

	isPrime := true
	for i := 2; i*i <= b; i++ {
		if b%i == 0 {
			isPrime = false
			break
		}
	}

	if isPrime {
		fmt.Println("PRIMA")
	} else {
		fmt.Println("BUKAN PRIMA")
	}
}
