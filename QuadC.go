package main

import "fmt"

func QuadC(x, y int) {
	if x <= 0 || y <= 0 {
		fmt.Print("geçersiz giris!!")
	}

	for i := 1; i <= y; i++ {
		for j := 1; j <= x; j++ {
			if i == 1 && j == 1 { // Sol üst köşe
				fmt.Print("A")
			} else if i == 1 && j == x { // Sağ üst köşe
				fmt.Print("A")
			} else if i == y && j == 1 { // Sol alt köşe
				fmt.Print("C")
			} else if i == y && j == x { // Sağ alt köşe
				fmt.Print("C")
			} else if i == 1 || i == y { // Üst ve alt kenar
				fmt.Print("B")
			} else if j == 1 || j == x { // Sol ve sağ kenar
				fmt.Print("B")
			} else {
				fmt.Print(" ")
			}
		}
		fmt.Println()
	}
}

func main() {
	QuadC(5, 3)
	fmt.Println()
	QuadC(5, 1)
	fmt.Println()
	QuadC(1, 1)
	fmt.Println()
	QuadC(1, 5)
}
