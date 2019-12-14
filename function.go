package main

import "fmt"

func main() {
	res := plus(1,2)
	fmt.Println(res)

	res = plusplus(1,2,4)
	fmt.Println(res)
}

func plus(a int,b int) int  {
	return a + b
}

func plusplus(a,b,c int) int  {
	return a + b + c
}