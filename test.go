package main

import "fmt"

func main() {
	gcd(20,18)
}

func gcd(x, y int) int {
	for y != 0 {
		x, y = y, x%y
		fmt.Println(x,y)
	}
	return x
}