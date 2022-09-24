package main

import "fmt"

func main() {
	var intPtr *int
	fmt.Println(intPtr)
	intPtr2 := 10
	intPtr = &intPtr2
	fmt.Println(intPtr)
	fmt.Println(&intPtr)
	fmt.Println(*intPtr)
}
