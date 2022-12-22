//Write a program which prompts the user to enter a floating point number 
//and prints the integer which is a truncated version of the floating point number 
//that was entered. Truncation is the process of removing the digits to the right of the 
//decimal place.
package main

import "fmt"

func main() {
	var input_floating_point_number float64
	var output_truncated_number int
    fmt.Printf("enter a floating point number : ")
	fmt.Scan(&input_floating_point_number)
	output_truncated_number = int(input_floating_point_number)
	fmt.Println("truncated version of the floating point number : ",output_truncated_number)
}