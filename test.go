package main

import "fmt"

func setMap(numbers map[string]int) {
    numbers["A"] = 1
    numbers["B"] = 2
    numbers["C"] = 3
    numbers["D"] = 4
    numbers["E"] = 5
    numbers["B"] = 5
}

type Move interface{
    move() string
}
type Car struct {
    name string 
}

type Animal struct{
    name string
}

func (car Car) move() string{
    return "Moving by wheel"
}

func (a Animal) move() string{
    return "Moving by leg"
}



func main() {

    car1 := Move{Car{"Mescedes"}}

}
