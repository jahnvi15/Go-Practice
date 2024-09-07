package main

import "fmt"
import "rsc.io/quote"

func main() {
    fmt.Println("Hello, World!")
	fmt.Println(quote.Go())
	fmt.Println("go" + " string concat")
	fmt.Println("1+1 = ", 1+1)
	fmt.Println("7.0/3.0", 7.0/3.0)
	fmt.Println("7/3", 7/3)
	fmt.Println(true && false)
	fmt.Println(true || false)
	fmt.Println(!true)
	
}