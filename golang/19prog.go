package main
import ("fmt"
        "sort")
func main() {
	fmt.Println("Integers reverse Sort")
	num:= []int{50,40,60,9,80,}
	sort.Sort(sort.Reverse(sort.IntSlice(num)))
	fmt.Println(num)
	fmt.Println()
	fmt.Println("Strings reverse Sort")
	text:= []string{"Japan","UK","Austrialia","Russia",}

	sort.Sort(sort.Reverse(sort.StringSlice(text)))
	fmt.Println(text)
}