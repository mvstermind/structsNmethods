package main

import (
	"bufio"
	"fmt"
	"os"
)

type Person struct {
	Name string
	Age  int
}

func (p *Person) Client() (bool, string) {
	fmt.Print("Enter your name: ")
	b := bufio.NewReader(os.Stdin)
	name, _ := b.ReadString('\n')

	if name != "" {
		fmt.Print("Your age: ")
		fmt.Scanf("%d\n", &p.Age)
		switch {
		case p.Age >= 18:
			return true, name
		case p.Age < 18:
			return false, ""
		}
	} else {
		return false, ""
	}
	return true, name
}

type Shop struct {
	Product  string
	Quantity int
	Adress   string
	Price    int
}

func (s *Shop) delivery(name string) (string, int, string) {
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

func (s *Shop) price(q int) int {
	p := Shop{Price: 32}
	total := p.Price * q
	return total
}

func main() {
	customer := Person{}
	boo, name := customer.Client()
	if boo == true {
		shop := Shop{}
		product, quantity, adress := shop.delivery(name)
		fmt.Printf("Thanks for your order\n")
		fmt.Printf("Product: %v\n", product)
		fmt.Printf("Quantity: %v\n", quantity)
		fmt.Printf("Your adress: %v\n", adress)
		fmt.Printf("Total Price: %v\n", shop.price(quantity))
	} else {
		fmt.Println("sorry you can't order here")
	}
}
