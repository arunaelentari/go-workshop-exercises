package main

import "fmt"

// Expected output:
//
// $ go run loops.go
// 0
// 1
// 2
// 3
// 5
// 6
// 7
// 9

func main() {
	// LOOP GOES HERE
	for i := 0; i < 10; i++ {
		if i != 4 && i != 8 {
			fmt.Println(i)
		}
	}
}
