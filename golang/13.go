package main
import "fmt"
func main() {
	var number,temp,remainder int
	var reverse int=0
	fmt.Println("enter a number")
	fmt.Scan(&number)
	temp=number
	for{
		remainder=number%10
		reverse=reverse*10+remainder
		number/=10
		if number==0{
			break
		}	

	}
	if temp==reverse{
		fmt.Println("palindrome",temp)
	}else{
		fmt.Println("not palindrome",temp)	
	}
}