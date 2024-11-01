package main

import "fmt"

func QuadA(x, y int) {
	if x <= 0 || y <= 0 {
		fmt.Print("geçersiz giris!!")
	}

	for i := 1; i <= y; i++ {
		for j := 1; j <= x; j++ {
			if i == 1 && (j == 1 || j == x) || i == y && (j == 1 || j == x) {
				fmt.Print("o")
			} else if i == 1 || i == y {
				fmt.Print("-")
			} else if j == 1 || j == x {
				fmt.Print("|")
			} else {
				fmt.Print(" ")
			}

		}
		fmt.Println()
	}
}
func main() {
	QuadA(5, 3)
	QuadA(1, 1)
	QuadA(5, 1)
	QuadA(1, 5)
}
