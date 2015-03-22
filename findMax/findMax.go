package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	var array [10]int
	rand.Seed( time.Now().UTC().UnixNano())
	
	for i:= 0; i < len(array); i++{
		array[i] = rand.Intn(100)
		fmt.Printf("%d \n", array[i])
	}
	
	var max int = array[0]
	
	for i:= 1; i < len(array); i++{
		if(array[i] > max){
			max = array[i]
		}
	}
	fmt.Printf("Found max is ")
	fmt.Printf("%d \n",max)
}