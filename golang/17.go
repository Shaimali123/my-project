package main
import ("fmt")
func main(){
	var s1,s2 string
	fmt.Printf("enter 1st string:")
	fmt.Scanln(&s1)
	fmt.Printf("enter 2nd string:")
	fmt.Scanln(&s2)
	result:=s1+s2
	fmt.Printf("concatenated string is %s",result) 
}
 