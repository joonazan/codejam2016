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
	var n int
	fmt.Scan(&n)

	if n == 0 {
		fmt.Println("INSOMNIA")
		return
	}

	set := make(map[rune]bool)

	sum := n
	for ; len(set) < 10; sum += n {
		for _, c := range fmt.Sprint(sum) {
			set[c] = true
		}
	}

	fmt.Println(sum - n)
}
