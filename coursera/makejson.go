// Write a program which prompts the user to first enter a name, and then enter an address. 
// Your program should create a map and add the name and address to the map using the keys “name” and “address”, respectively. 
// Your program should use Marshal() to create a JSON object from the map, and then your program should print the JSON object.

package main

import (
	"fmt"
	"encoding/json"
	"bufio"
	"os"
)

func main(){
	data_map := make(map[string]string)
	consoleReader := bufio.NewReader(os.Stdin)
	fmt.Printf("enter name : ")               
	name, name_read_error := consoleReader.ReadString('\n')
	fmt.Printf("enter address : ")               
	address, address_read_error := consoleReader.ReadString('\n')
	if name_read_error == nil && address_read_error == nil  {
		data_map[name] = address
		json_data, error := json.Marshal(data_map)
		if(error == nil){
			fmt.Println("Marshaled json data in bytes",json_data)
			temp_data_map := make(map[string]string)
			json.Unmarshal([]byte(json_data),&temp_data_map)
			fmt.Println("Unmarshaled json data",temp_data_map)
		}
	}
}