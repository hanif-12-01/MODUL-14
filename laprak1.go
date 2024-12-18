package main

import "fmt"

func main() {
	var bilangan, j, counter int
	fmt.Scan(&bilangan)
	for j = 0; j <= bilangan; j++ {
		if j%2 != 0 {
			counter++
		}
	}
	fmt.Printf("TERDAPAT %d BILANGAN GANJIL", counter)
}
