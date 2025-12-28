package main
import "fmt"

func main()  {

	var chicken map[string]int
	chicken = map[string]int{}

	chicken["januari"] = 50
	chicken["febuari"] = 30

	fmt.Println("Januari", chicken["januari"])
	fmt.Println("Febuari", chicken["febuari"])
	fmt.Println("Mei", chicken["mei"])
	
}
