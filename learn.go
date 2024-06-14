package main

import "fmt"

func add() {
	x := 2
	y := 2

	return x + y
}

func hex() {	
	fmt.Println("say h3x")
}

func main() {
	hex()
	add()
}