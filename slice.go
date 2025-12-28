package main

import "fmt"

func main()  {
	var fruits = [] string {"Apple","Manggo","banana","Melon"}
	fmt.Println(fruits[2])
	
	var minuman = []string {"Juice Melon","Juice Manggo","Juice Apple","Juice Pockat"}
	var newMinuman = minuman[0:3]
	var newMinuman2 = minuman[1:4]
	fmt.Println(newMinuman)
	fmt.Println(len(newMinuman))
	fmt.Println(cap(newMinuman))
	fmt.Println(cap(newMinuman2))
}
