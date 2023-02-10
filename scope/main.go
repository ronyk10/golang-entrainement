package main

import "fmt"

var x int // variable globale, utilisable partout

func main() {
	x = 5
	fmt.Println(x) // 5
	f()
	ShowY()
}

func f() {
	x := 10 //var locale
	fmt.Println(x)
}
