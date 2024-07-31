package main

import "fmt"

var z int

func main() {
	Greeting()
	var x int
	var y int
	fmt.Printf("Enter Three Number \n")
	fmt.Scan(&x)
	fmt.Scan(&y)
	fmt.Scanf("%d", &z)
	sum := sum(x, y)
	fmt.Println("Sum of Two Number =", sum)
	fmt.Println(x, "        ", y)
	fmt.Println("**************************")
	x, y = swapNumber(x, y)
	fmt.Println(x, "        ", y)
	fmt.Println(z)

}
func sum(a, b int) int {
	return (a + b)
}
func Greeting() {
	fmt.Println("Hii College Over")
}
func swapNumber(a, b int) (int, int) {
	return b, a
}
