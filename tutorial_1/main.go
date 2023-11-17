package main

import "fmt"

func main() {
	fmt.Println("Hello, World!")
	// var intNum int16 = 32767 + 1
	var intNum int16 = 32767
	intNum += 1 // Overflow Error but not doesn't crash
	fmt.Println(intNum)
}
