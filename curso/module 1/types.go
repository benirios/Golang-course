package main

import "fmt"

func main() {
	//bool (t/f)
	fmt.Printf("type: %T - value %v\n", true, true)
	//string (text)
	fmt.Printf("type: %T - value %v\n", "string", "string")
	//int (without decimal)
	fmt.Printf("type: %T - value %v\n", 1, 1)
	//float (64/32) (decimal)
	fmt.Printf("type: %T - value %v\n", 0.25, 0.25)

}
