// Write a program which allows the user to get information about a predefined set of 
// animals. Three animals are predefined, cow, bird, and snake. 
// Each animal can eat, move, and speak. The user can issue a request to find out one of 
// three things about an animal:
//  1) the food that it eats, 2) its method of locomotion, and 
//  3) the sound it makes when it speaks. 
//  The following table contains the three animals and their associated data which should 
//  be hard-coded into your program.
//  Animal	Food eaten	Locomotion method	Spoken sound
//  cow		grass		walk				moo
//  bird	worms		fly					peep
//  snake	mice		slither				hsss 
//  Your program should present the user with a prompt, “>”, to indicate that the user can
//   type a request. Your program accepts one request at a time from the user, prints out 
//   the answer to the request, and prints out a new prompt. 
//   Your program should continue in this loop forever. Every request from the user must 
//   be a single line containing 2 strings. 
//   The first string is the name of an animal, either “cow”, “bird”, or “snake”. 
//   The second string is the name of the information requested about the animal, 
//   either “eat”, “move”, or “speak”. Your program should process each request by 
//   printing out the requested data.

// You will need a data structure to hold the information about each animal. 
// Make a type called Animal which is a struct containing three fields:food, 
// locomotion, and noise, all of which are strings. 
// Make three methods called Eat(), Move(), and Speak(). 
// The receiver type of all of your methods should be your Animal type. 
// The Eat() method should print the animal’s food, the Move() method should print the
//  animal’s locomotion, and the Speak() method should print the animal’s spoken sound. 
//  Your program should call the appropriate method when the user makes a request.
package main

import (
	"fmt"
	"bufio"
	"strings"
	"os"
)

type Animal struct{
	food_eaten string
	locomotion_method string
	spoken_sound string
}

func (A *Animal) InitAnimal(food_eaten, locomotion_method, spoken_sound string){
	A.food_eaten		= food_eaten
	A.locomotion_method	= locomotion_method
	A.spoken_sound		= spoken_sound
}

func (A *Animal) Eat() string{
	return A.food_eaten
}

func (A *Animal) Move() string{
	return A.locomotion_method
}
   
func (A *Animal) Speak() string{
	return A.spoken_sound
}

func PrintInfoRequested(animal_instance *Animal, info_requested string){
	switch info_requested{
	case "eat":
		fmt.Println(animal_instance.Eat())
	case "move":
		fmt.Println(animal_instance.Move())
	case "speak":
		fmt.Println(animal_instance.Speak())
	}
}

func main(){
	var cow Animal
	cow.InitAnimal("grass", "walk", "moo")
	var bird Animal
	bird.InitAnimal("worms", "fly", "peep")
	var snake Animal
	snake.InitAnimal("mice", "slither", "hsss")
	consoleReader := bufio.NewReader(os.Stdin)
	flag_to_exit := " "
	for flag_to_exit != "x"{
		fmt.Println("Enter input example(cow eat), Enter 'x' to exit >: ")
		user_input, user_input_error := consoleReader.ReadString('\n')
		if user_input_error == nil{
			user_input = strings.TrimSpace(user_input)
			user_input = strings.ToLower(user_input)
			user_inputs := strings.Fields(user_input)
			user_inputs_length := len(user_inputs)
			if user_inputs_length == 1{
				flag_to_exit = user_inputs[0]
			}else if user_inputs_length == 2{
				switch user_inputs[0]{
				case "cow":
					PrintInfoRequested(&cow,user_inputs[1])
				case "bird":
					PrintInfoRequested(&bird,user_inputs[1])
				case "snake":
					PrintInfoRequested(&snake,user_inputs[1])
			}
		}
		}
	}
}