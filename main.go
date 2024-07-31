package main

import (
	inner "firstmodule/innerg"
	"firstmodule/outer"
	"fmt"
	"os"
)

func main() {
	fmt.Println("Raushan Raj")
	fmt.Printf("Anurag")
	os.Exit(0)
	inner.Greeting()
	outer.HelloGreeting()
}
