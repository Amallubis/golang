package main

import "fmt"

func main()  {

	var fruits = [4]string{"Apple","Mango","benana","Melon"}

	for _, fruit := range fruits {
		fmt.Println("Nama buah", fruit)
	}
	
}
