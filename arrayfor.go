package main

import "fmt"

func main()  {

	var fruits = [4] string {"Nanas","Durian","Jambu","Salak"}

	for i := 0; i < len(fruits); i ++{
		fmt.Printf("elemen %d : %s\n", i, fruits[i])
	}
	
}
