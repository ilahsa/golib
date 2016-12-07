package main

import (
	"fmt"
)

func main() {
	fmt.Println("start")
	
	a:=new(int)
	defer func(){
		fmt.Println(*a)
	}()
	for i:=0;i<100;i++{
		*a++
	}
	// var a string
	// fmt.Scan(&a)
}
