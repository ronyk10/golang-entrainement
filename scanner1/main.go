package main

import (
	"bufio" //pour le scanner
	"fmt"
	"os"
	"strconv" //atoi (string to int)
)

//Learning about basic scanner

// func main() {
// 	scanner := bufio.NewScanner(os.Stdin) // creating a new scanner
// 	fmt.Print("Write something :")
// 	scanner.Scan()              // scanner starting
// 	userWrite := scanner.Text() // scanner result in stock
// 	fmt.Println("You writed :", userWrite, ". Reload if you want to write something again !")
// }

// Convert a string in int with atoi
func main() {
	scanner := bufio.NewScanner(os.Stdin) // creating a new scanner
	fmt.Print("Please enter a number  :")
	scanner.Scan() // scanner starting

	nbr, _ := strconv.Atoi(scanner.Text()) //convert a string in int
	fmt.Printf("res : %d\n", (nbr + 6))
}
