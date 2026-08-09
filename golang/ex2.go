package main
import "fmt"

func main() {
    person := struct {
        name string
        age  int
    }{
        name: "ravi",
        age:  20,
    }

    fmt.Println("name:", person.name)
    fmt.Println("age:", person.age)
}
