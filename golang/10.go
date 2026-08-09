package main
import "fmt"
func main() {
	var n int
	fmt.Println("enter a number")
	fmt.Scan(&n)
	i:=1
	for{
		if i>10{
			break
		}
		fmt.Printf("%d x %d = %d\n",n,i,n*i)
		i++
	}
}