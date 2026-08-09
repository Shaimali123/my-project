package main

import "fmt"

func main() {
    fmt.Println	("enter a number")
	var n int
	fmt.Scan(&n)
	if(n%2==0){
		fmt.Println("even")
	}else{
		fmt.Println("odd")
	}
}