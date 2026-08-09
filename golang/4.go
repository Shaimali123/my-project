package main
import (
	"fmt"
	"strings"
)
func main() {
	str:="the cat is hot"
	i:=strings.Index(str,"cat")
	fmt.Println("index of cat is ",i)
}