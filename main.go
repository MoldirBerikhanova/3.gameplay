package main

import (
	"fmt"
	"math/rand"
)

func main() {
	//counter9
	// while loop

	// elements := []string("камень", "ножницы", "бумага")
	var userChoice int
	var compChoice int

	for {
		fmt.Println("сделайте ваш выбор:")
		fmt.Println("1.камень")
		fmt.Println("2.ножницы")
		fmt.Println("3.бумага")
		fmt.Scan(&userChoice) //1
		compChoice = rand.Intn(3) + 1
		// fmt.Println(userChoice , compChoice)// отображает выбор пользователя  и компьютера

		// 2 random.3
		if userChoice == 1 && compChoice == 2 {
			fmt.Println("Вы выиграли")
		} else if userChoice == 1 && compChoice == 1 {
			fmt.Println("Ничья")
		} else if userChoice == 2 && compChoice == 1 {
			fmt.Println("Вы выиграли")
		} else if userChoice == 2 && compChoice == 2 {
			fmt.Println("Ничья")
		} else if userChoice == 3 && compChoice == 2 {
			fmt.Println("Вы проиграли")
		} else if userChoice == 3 && compChoice == 3 {
			fmt.Println("Ничья")

			//you win
			//userwin++
		}
		//random comp choice [1-3]
		//if else
	}
}

// func generateComputerChoice() string {
// 	source := rand.NewSource(time.Now().UnixNano())
// 	random := rand.New(source)

// 	index := random.Intn(3) 

// 	return choices[index]
// }
