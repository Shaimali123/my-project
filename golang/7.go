package main
import "fmt"
func main() {
	var row int
	var k int=0
	fmt.Println("enter number of rows")
	fmt.Scanln(&row)
	for i:=1;i<=row;i++{
		k=0
		for space:=1;space<=row-i;space++{
			fmt.Print(" ")
		}
		for{
			fmt.Print("*")
			k++
			if k==2*i-1{
				break
			}
			fmt.Print(" ")
		}
		fmt.Println()
		}
	}
