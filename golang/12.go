package main 
import "fmt" 
var area int 
func main() { 
var l, b int 
fmt.Println("enter length of rectangle :") 
fmt.Scanln(&l) 
fmt.Println("enter the breadth of rectangle :") 
fmt.Scanln(&b) 
area = l * b 
fmt.Println("area of rectangle :", area) 
fmt.Println("enter length of square :") 
fmt.Scanln(&l) 
area = l * l 
fmt.Println("area of square :", area) 
}