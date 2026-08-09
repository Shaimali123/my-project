package main

import "fmt"

type Person struct {
    name string
    age  int
}

type Employee struct {
    Person 
    id int
}

func main() {
    e := Employee{
        Person: Person{name: "Ravi", age: 25},
        id:101,
    }

    fmt.Println(e.name) 
    fmt.Println(e.age)
    fmt.Println(e.id)
}