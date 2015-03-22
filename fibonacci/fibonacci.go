package main

import (
	"fmt"
)

func main() {
	var seed1 int = 0
	var seed2 int = 1
	var n int = 50
	
	fmt.Printf("%d \n",seed1)
	fmt.Printf("%d \n",seed2)
	
	for i := 0; i < n-2; i++{
		var sum int = seed1+seed2
		fmt.Printf("%d \n",sum)
		seed1 = seed2
		seed2 = sum		
	}
}