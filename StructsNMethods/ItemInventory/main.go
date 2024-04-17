package main

import (
	"fmt"
	"inventory/person"
	"inventory/shop"
)

func main() {
	customer := person.Person{}
	boo, name := customer.Client()
	if boo == true {
		shop := shop.Shop{}
		product, quantity, adress := shop.Delivery(name)
		fmt.Printf("Thanks for your order\n")
		fmt.Printf("Product: %v\n", product)
		fmt.Printf("Quantity: %v\n", quantity)
		fmt.Printf("Your adress: %v\n", adress)
		fmt.Printf("Total Price: %v\n", shop.PriceCustomer(quantity))
	} else {
		fmt.Println("sorry you can't order here")
	}
}
