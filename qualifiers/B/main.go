package main

import "fmt"

func main() {
	var t int
	fmt.Scan(&t)

	for i := 0; i < t; i++ {
		fmt.Printf("Case #%d: ", i+1)
		tapaus()
	}
}

func tapaus() {
	var s string
	fmt.Scan(&s)
	n := len(s)

	goal := uint(0)
	for _, c := range s {
		goal <<= 1
		if c == '+' {
			goal += 1
		}
	}
	goal = flip(goal, n)

	visited := make([]bool, 1024)
	current := []uint{0}
	next := []uint{}

	for epoch := 0; ; epoch++ {
		for _, c := range current {
			if c == goal {
				fmt.Println(epoch)
				return
			}

			if visited[c] {
				continue
			}
			visited[c] = true

			for i := 1; i <= n; i++ {
				next = append(next, flip(c, i))
			}
		}

		current = next
		next = []uint{}
	}
}

func flip(stack uint, take int) (out uint) {
	out = stack

	// zero the bits of the flipped part
	out &^= (1 << uint(take)) - 1

	for i := 0; i <= take; i++ {
		out |= (1 &^ stack) << uint(take-i-1)
		stack >>= 1
	}

	return
}
