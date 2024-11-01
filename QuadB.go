package main

import "fmt"

func QuadB(x, y int) {
	if x <= 0 || y <= 0 {
		fmt.Print("geçersiz giris!!")
	}

	for i := 1; i <= y; i++ {
		for j := 1; j <= x; j++ {
			if i == 1 && j == 1 { // Sol üst köşe
				fmt.Print("/")
			} else if i == 1 && j == x { // Sağ üst köşe
				fmt.Print("\\")
			} else if i == y && j == 1 { // Sol alt köşe
				fmt.Print("\\")
			} else if i == y && j == x { // Sağ alt köşe
				fmt.Print("/")
			} else if i == 1 || i == y { // Üst ve alt kenar
				fmt.Print("*")
			} else if j == 1 || j == x { // Sol ve sağ kenar
				fmt.Print("*")
			} else {
				fmt.Print(" ")
			}
		}
		fmt.Println()
	}
}

func main() {
	QuadB(5, 3)
	fmt.Println()
	QuadB(5, 1)
	fmt.Println()
	QuadB(1, 1)
	fmt.Println()
	QuadB(1, 5)
}
