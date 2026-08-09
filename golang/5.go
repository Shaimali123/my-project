package main
import "fmt"
func main() {
	intslice := []int{1,2,3,4,5}
	fmt.Printf("slice :%v\n ", intslice)
	last := intslice[len(intslice)-1]
	fmt.Printf("last element :%v\n ", last)
	first := intslice[:0]
	fmt.Printf("first element :%v\n ", first)
	remove := intslice[:len(intslice)-1]
	fmt.Printf("slice after removing last element :%v\n ", remove)
}