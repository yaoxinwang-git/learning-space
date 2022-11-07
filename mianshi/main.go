package main

import "fmt"

var c = make(chan int)
var a int

func f() {
	a = 10
	b := <-c
	fmt.Printf("b = %d\n", b)
}
func main() {
	//go f()
	//c <- 1000
	//print(a)

	fmt.Println((1 << 31) - 1)
}
