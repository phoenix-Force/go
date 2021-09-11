package main

import (  
    "fmt"
)

func application(a int, b int)(int,int,int int) { //add,sub,mul, div
	add := a + b
	sub := a-b
	mul := a*b
	div := a/b 
	return add,sub,mul,div
}
func main() {
	add, sub, _ , div := application(5,2);
}