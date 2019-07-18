package main

import "fmt"

func Sum(x int, y int) int {
	return x + y
}

func subtraction(x, y int) int {
	return x - y
}

func main()  {
	sum := Sum(2,3)
	sub := subtraction(10, 5)

	//printing result
	fmt.Println(sum)
	fmt.Println(sub)
}