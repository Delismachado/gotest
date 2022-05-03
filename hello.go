package main

import "fmt"

const prefixPortHi = "Hi "

func Hello(name string) string {
	if name == "" {
		name = "Any"
	}
	return prefixPortHi + name

}

func main() {
	fmt.Println(Hello("anyone"))
}
