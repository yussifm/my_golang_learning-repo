package main

import (
	"fmt"
)

/*
 multiline comments

*/
const favColor string = "blue"
func main() {

	// comments
	// var name string = `hello`;
	// var age int = 20;
	// x:=5;
	// x +=1;


	// fmt.Println("Hello Go!");
	// fmt.Println(name);
	// fmt.Println(age);
	// fmt.Println(x);
	// fmt.Println("1 + 1 =", 1 + 1)
	// fmt.Scanln(&name)
	// name = strings.TrimSpace(name)
	// fmt.Println(name);

	// i := 0
	// for i < 10 {
	// 	fmt.Println("counting numbers", i)
	// 	if i%2 == 0 {
	// 		println(i, " is Even")
	// 	} else if i%2 != 0 {
	// 		println(i, " is Odd")
	// 	}

	// 	i = i + 1

	// }

	// Guessing Colors

	var guess string
	for {
		fmt.Println("Guess my Color")

		if _, err := fmt.Scanln(&guess); err != nil {
			fmt.Printf("Error: %v\n", err)
			return
		}
		if favColor == guess {
			fmt.Printf("%v is my favorite color\n", guess)
			return
		}

		fmt.Printf("sorry, %v is not my favorite color, guess again\n", guess)
	}
}