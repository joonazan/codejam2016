package main

import "fmt"

func main() {
	fmt.Println("Case #1:")
	tapaus()
}

func tapaus() {
	a := alkuluvut(33333333)

	fmt.Println(a[:100])

	jaollinen := func(x int) int {
		for _, alkuluku := range a {
			if x%alkuluku == 0 && x != alkuluku {
				return alkuluku
			}
		}
		return 0
	}

	var jakajat [9]int

	for i := 0; i < (1 << 14); i++ {
		luku := 1<<15 + i<<1 + 1
		for b := 2; b <= 10; b++ {
			if jakaja := jaollinen(toBase(luku, b)); jakaja == 0 {
				goto fail
			} else {
				jakajat[b-2] = jakaja
			}
		}
		fmt.Printf("%b ", luku)
		for _, j := range jakajat {
			fmt.Print(j, " ")
		}
		fmt.Println()
	fail:
	}
}

func toBase(luku, b int) (summa int) {
	a := 1
	for i := 0; i < 16; i++ {
		summa += a * (luku & 1)
		luku >>= 1
		a *= b
	}
	return
}

func alkuluvut(n int) (luvut []int) {
	eiAlkuluku := make([]bool, n+1)
	eiAlkuluku[0] = true
	eiAlkuluku[1] = true

	for i := 2; i <= n; i++ {
		if eiAlkuluku[i] {
			continue
		}
		for j := i * 2; j <= n; j += i {
			eiAlkuluku[j] = true
		}
	}

	for i, ei := range eiAlkuluku {
		if !ei {
			luvut = append(luvut, i)
		}
	}

	return
}
