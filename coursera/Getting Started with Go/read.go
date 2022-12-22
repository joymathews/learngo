// Write a program which reads information from a file and represents it in a slice of structs. 
// Assume that there is a text file which contains a series of names. 
// Each line of the text file has a first name and a last name, in that order, separated by a single space on the line. 

// Your program will define a name struct which has two fields, fname for the first name, and lname for the last name. 
// Each field will be a string of size 20 (characters).

// Your program should prompt the user for the name of the text file. 
// Your program will successively read each line of the text file and create a struct which contains the first and 
// last names found in the file. Each struct created will be added to a slice, and after all lines have been 
// read from the file, your program will have a slice containing one struct for each line in the file. 
// After reading all lines from the file, your program should iterate through your slice of structs and print 
// the first and last names found in each struct.

package main

import (
	"fmt"
	"os"
	"bufio"
	"io/ioutil"
	"strings"
)
type NameStruct struct{
	fname[20] byte
	lname[20] byte
}

func main(){
	
	consoleReader := bufio.NewReader(os.Stdin)
	fmt.Printf("enter file name : ")
	file_name, file_name_read_error := consoleReader.ReadString('\n')
	if file_name_read_error == nil{
		file_name = strings.TrimSpace(file_name)
		file_data, file_read_error := ioutil.ReadFile(file_name)
		if file_read_error == nil{
			words_in_file := strings.Fields(string(file_data))
			no_of_words_in_file := len(words_in_file)
			if no_of_words_in_file%2 == 0{
				word_index := 0 
				no_of_elements_in_struct := no_of_words_in_file/ 2
				struct_slice := make([]NameStruct,no_of_elements_in_struct)
				for struct_index := 0;  struct_index < no_of_elements_in_struct ; struct_index++{
					var new_struct NameStruct
					copy(new_struct.fname[:],words_in_file[word_index])
					word_index = word_index + 1
					copy(new_struct.lname[:], words_in_file[word_index] )
					word_index = word_index + 1
					struct_slice[struct_index] = new_struct
				}
				for index, ele_struct := range struct_slice {
					fmt.Println("Struct element : ",(index+1))
					fmt.Println("fname : ",string(ele_struct.fname[:]))
					fmt.Println("lname : ",string(ele_struct.lname[:]))
				}
			}else{
				fmt.Println("Invalid file format")
			}			
		}else{
			fmt.Println("Error in reading file content",file_read_error)
		}
	}else{
		fmt.Println("Error in reading file name",file_name_read_error)
	}
}
	
