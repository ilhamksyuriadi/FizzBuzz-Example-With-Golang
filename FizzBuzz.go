package main
import (
	"fmt"
	"math"
)

func main(){
	var start, end, i int
	fmt.Print("First number: ")
	fmt.Scanln(&start)
	fmt.Print("Last number: ")
	fmt.Scanln(&end)
	i = start
	if start <= end {
		for i <= end {
			if math.Mod(float64(i),3) == 0 && math.Mod(float64(i),5) == 0 {
				fmt.Println("FizzBuzz")
			} else if math.Mod(float64(i),3) == 0 {
				fmt.Println("Fizz")
			} else if math.Mod(float64(i),5) == 0 {
				fmt.Println("Buzz")
			} else {
				fmt.Println(i)
			}
			i += 1
			if i == end+1 {
				break
			}
		}
	}else {
		fmt.Println("The start number must be smaller than end number")
	}
}