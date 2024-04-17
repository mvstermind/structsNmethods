package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Inventory struct {
	Name     string
	Quantity int
	Price    int
	ID       int
}

func (n *Inventory) AddItem() {
	b := bufio.NewReader(os.Stdin)

	fmt.Print("Enter product name: ")
	n.Name, _ = b.ReadString('\n')
	n.Name = strings.TrimSpace(n.Name)

	fmt.Print("Item quantity: ")
	fmt.Scanf("%d\n", &n.Quantity)

	fmt.Print("Price: ")
	fmt.Scanf("%d\n", &n.Price)

	fmt.Print("Item ID: ")
	fmt.Scanf("%d\n", &n.ID)
}

func main() {
	item := Inventory{}
	item.AddItem()

	fmt.Printf("Product name: %v\n", item.Name)
	fmt.Printf("Item Quantity: %d\n", item.Quantity)
	fmt.Printf("Item Price: %d\n", item.Price)
	fmt.Printf("Item ID: %d\n", item.ID)

}
