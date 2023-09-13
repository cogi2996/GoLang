package main

import (
	"fmt"
	"time"
)

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


// what to do
//--> duyet qua tung phan tu objInters 
// how to do
//--> su dung for range

func say(threadName string, number int){
    for i:=0;i<number;i++{
        fmt.Println(threadName, i+1);
        time.Sleep(1*time.Millisecond)
    }
}

func testGoroutine(){
    number :=10
    go say("thread 1",number)
    say("thread 2",number)
    time.Sleep(1*time.Millisecond)
}

func hello(done chan bool){
    fmt.Println("hello is starting")
    time.Sleep(1*time.Millisecond)
    fmt.Println("hello is end")
    done<-true
}

func getSquare(x int64) int64{
    time.Sleep(2*time.Millisecond)
    return x*4
}

func getRectangle(x int64,y int64)int64{
    time.Sleep(2*time.Millisecond)
    return(x+y)*2
}

func test(){
    done1 :=make(chan bool)
    done2 :=make(chan bool)
    done3 :=make(chan bool)
    fmt.Println("Main is start")
    go hello(done1)
    go hello(done2)
    go hello(done3)
    result1 := <-done1
    result2 := <-done2
    result3 := <-done3
    fmt.Println("Main received data:" ,result1 ,result2 ,result3)
}
func main() {
   test()
}
