package main

import "fmt"

func main()  {

	fruits := []string {"Teh Botol","Sprite","Cococola","Fanta"}
	fmt.Println(fruits)
	var fruit = fruits[1:3]
	fmt.Println(fruit)

	var fruitapp = append(fruit, "Juice")
	fmt.Println(fruitapp)
	
}
