package main

import "fmt"

func main() {
	var t int
	fmt.Scan(&t)

	for i := 0; i < t; i++ {
		fmt.Printf("Case #%d:", i+1)
		tapaus()
	}
}

func tapaus() {
	var k, c, s int
	fmt.Scan(&k, &c, &s)

	needed := (k + c - 1) / c

	if needed > s {
		fmt.Println(" IMPOSSIBLE")
		return
	}

	pos := 0
	for i := 0; i < needed; i++ {
		stride := 1
		sum := 0
		for j := 0; j < c; j++ {
			sum += stride * pos
			pos++
			if pos >= k {
				pos = k - 1
			}
			stride *= k
		}
		fmt.Print(" ", sum+1)
	}
	fmt.Println()
}
