package main
import "fmt"

func main()  {

	var point = 5

	switch point {
	case 8:
		fmt.Println("Perfect")
	case 7, 6, 5, 4:
		fmt.Println("awesome")
	default: 
	  fmt.Println("nod bad")
		fmt.Println("You can be better")
		
	}
	
}
