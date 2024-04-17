package person

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
