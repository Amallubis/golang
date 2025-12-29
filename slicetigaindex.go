package main

import "fmt"

func main()  {

	var fruits = []string{"Apple","Grape","Banana"}
	var afruits = fruits[0:2]
	var bfruits = fruits[0:2:2]

	fmt.Println(fruits)
	fmt.Println(len(fruits))
	fmt.Println(cap(fruits))

	fmt.Println(afruits)
	fmt.Println(len(afruits))
	fmt.Println(cap(afruits))

	fmt.Println(bfruits)
	fmt.Println(len(bfruits))
	fmt.Println(cap(bfruits))
	
}
