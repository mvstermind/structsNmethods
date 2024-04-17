package shop

import (
	"bufio"
	"fmt"
	"os"
)

type Shop struct {
	Product  string
	Quantity int
	Adress   string
	Price    int
}

func (s *Shop) Delivery(name string) (string, int, string) {
	if name != "" {
		b := bufio.NewReader(os.Stdin)

		fmt.Print("Type product name: ")
		s.Product, _ = b.ReadString('\n')
		if s.Product == "\n" {
			fmt.Println("You gotta pick a product")
			os.Exit(2)
		}

		fmt.Printf("How many items you wanna get: ")
		fmt.Scanf("%d\n", &s.Quantity)
		if s.Quantity <= 0 {
			fmt.Println("you gotta get somrthing")
			os.Exit(3)
		}

		fmt.Print("Delivery adress: ")
		s.Adress, _ = b.ReadString('\n')
		if s.Adress == "\n" {
			fmt.Println("Adress cannot be blank!")
			os.Exit(3)
		}

		return s.Product, s.Quantity, s.Adress
	} else {

		fmt.Println("Invalid Name")
		os.Exit(1)
		return "", 0, ""
	}
}

func (s *Shop) PriceCustomer(q int) int {
	p := Shop{Price: 32}
	total := p.Price * q
	return total
}
