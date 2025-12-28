package main

import "fmt"

func main()  {
	var fruits = [] string {"Apple","Manggo","banana","Melon"}
	fmt.Println(fruits[2])
	
	var minuman = []string {"Juice Melon","Juice Manggo","Juice Apple","Juice Pockat"}
	var newMinuman = minuman[0:2]
	fmt.Println(newMinuman)
}
