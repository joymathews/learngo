// Write a program which prompts the user to enter integers and stores the integers in a sorted slice.
// The program should be written as a loop. Before entering the loop, the program should create an empty 
// integer slice of size (length) 3. During each pass through the loop, the program prompts the user to 
// enter an integer to be added to the slice. The program adds the integer to the slice, sorts the slice, 
// and prints the contents of the slice in sorted order. The slice must grow in size to accommodate any 
// number of integers which the user decides to enter. The program should only quit (exiting the loop) 
// when the user enters the character ‘X’ instead of an integer.

package main

import (
    "fmt"
	"strings"
    "strconv"
)

func my_sort(input_slice []int){
	for i:=0;i<len(input_slice); i++{
		for j:=i+1;j<len(input_slice);j++{
			if input_slice[i] < input_slice[j]{
				temp := input_slice[j]
				input_slice[j] = input_slice[i]
				input_slice[i] = temp
			}
		}
	}
}

func main(){

	input_number_in_string := "z"
	integer_slice := make([]int,3)
	index := 0
	for input_number_in_string != "x"{
		fmt.Printf("enter a number : ")
		fmt.Scan(&input_number_in_string)
		input_number, error := strconv.Atoi(input_number_in_string)
		if(error == nil){
			if index == len(integer_slice) {
				integer_slice = append(integer_slice,input_number)
			}else{
				integer_slice[index] = input_number
			}
			index = index + 1
			my_sort(integer_slice)
			fmt.Println(integer_slice)
		}
		input_number_in_string = strings.ToLower(input_number_in_string)
	}
}