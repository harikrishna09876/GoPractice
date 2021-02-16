package main

//this is package declaration
import "fmt"

var s = "hi John"

/* this is main function
   this is the entry point
   there will be only one main function
*/
func main() {

	fmt.Println("hello Go and rock...", s)
	Come()
}

func hello() {

	fmt.Println(s)
}
