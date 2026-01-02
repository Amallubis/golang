package main
import "fmt"
import "math"

func calculate(d float64)(float64, float64)  {

	var area = math.Pi * math.Pow(d / 2, 2)
	var circumference = math.Pi * d

	return area, circumference
	
}

func main()  {
	var diamater float64 =15
	var area, circumference = calculate(diamater)
	fmt.Printf("luas lingkaran\t\t: %.2f \n", area)
	fmt.Printf("keliling lingkaran\t: %.2f \n", circumference)
	
}
