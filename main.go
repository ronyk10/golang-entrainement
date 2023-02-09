package main

import "fmt"

func main() {
	var list [3]int

	list[0] = 5
	list[1] = 10
	list[2] = 15

	fmt.Println(list)
	fmt.Println(list[0])
	fmt.Println(list[1])
	fmt.Println(list[2])

	newList := [2]int{40, 50}

	fmt.Println(newList)
	fmt.Println(newList[0])
	fmt.Println(newList[1])

	bigList := [...]int{1, 2, 3, 456, 789, 123, 456, 789, 456, 014, 852, 963}

	for pos, value := range bigList {
		fmt.Printf("Position %d est égale à  %d. \n", pos, value)
	}
}
