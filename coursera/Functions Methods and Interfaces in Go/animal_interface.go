package main

import (
	"fmt"
	"bufio"
	"strings"
	"os"
)

type AnimalInterface interface {
	Eat()
	Move()
	Speak()
}

type Cow struct{
	food_eaten string
	locomotion_method string
	spoken_sound string
}

type Bird struct{
	food_eaten string
	locomotion_method string
	spoken_sound string
}

type Snake struct{
	food_eaten string
	locomotion_method string
	spoken_sound string
}

func (c *Cow) Eat() {
	fmt.Println(c.food_eaten)
}

func (c *Cow) Move() {
	fmt.Println(c.locomotion_method)
}
   
func (c *Cow) Speak() {
	fmt.Println(c.spoken_sound)
}

func (b *Bird) Eat() {
	fmt.Println(b.food_eaten)
}

func (b *Bird) Move() {
	fmt.Println(b.locomotion_method)
}
   
func (b *Bird) Speak() {
	fmt.Println(b.spoken_sound)
}

func (s *Snake) Eat() {
	fmt.Println(s.food_eaten)
}

func (s *Snake) Move() {
	fmt.Println(s.locomotion_method)
}
   
func (s *Snake) Speak() {
	fmt.Println(s.spoken_sound)
}

func main(){
	aniaml_name_map := make(map[string]AnimalInterface)
	consoleReader := bufio.NewReader(os.Stdin)
	flag_to_exit := " "
	for flag_to_exit != "x"{
		fmt.Println("Enter input (newanimal animalname cow/bird/snake) or (query animalname eat/move/speak ), Enter 'x' to exit >: ")
		user_input, user_input_error := consoleReader.ReadString('\n')
		if user_input_error == nil{
			user_input = strings.TrimSpace(user_input)
			user_input = strings.ToLower(user_input)
			user_inputs := strings.Fields(user_input)
			user_inputs_length := len(user_inputs)
			if user_inputs_length == 1{
				flag_to_exit = user_inputs[0]
			}else if user_inputs_length == 3{
				if user_inputs[0] == "newanimal"{
					switch user_inputs[2]{
						case "cow":
							aniaml_name_map[user_inputs[1]] = &Cow{"grass", "walk", "moo"}
							fmt.Println("Created it!")
						case "bird":
							aniaml_name_map[user_inputs[1]] = &Bird{"worms", "fly", "peep"}
							fmt.Println("Created it!")
						case "snake":
							aniaml_name_map[user_inputs[1]] = &Snake{"mice", "slither", "hsss"}
							fmt.Println("Created it!")
						default:
							fmt.Println("Animal type not supported")
					}
				}else if user_inputs[0] == "query"{
					animal := aniaml_name_map[user_inputs[1]]
					if animal == nil{
						fmt.Println("No Animal with this name was created")
					}else{
					switch user_inputs[2]{
						case "eat":
							animal.Eat()
						case "move":
							animal.Move()
						case "speak":
							animal.Speak()
						default:
							fmt.Println("Invalid query")
					}
				}
				}
		}
		}
	}
}