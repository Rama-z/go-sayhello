package sayHello

import "fmt"

func SayHello(name string) string {
	fmt.Println("This is new feature in sayHello v1.1.0")
	return "Hello World!" + name
}