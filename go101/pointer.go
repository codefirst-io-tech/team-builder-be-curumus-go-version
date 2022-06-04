package main

import "fmt"

func main() {
	name := "abidino"
	pointerAddress := &name
	nameValue := *pointerAddress
	fmt.Print(nameValue)
}
