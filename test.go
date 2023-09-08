// You can edit this code!
// Click here and start typing.
package main

import "fmt"

func add(x,y int64) int64{
    return  x+ y
}

func mul(x,y int64) int64{
    return x*y
}

func cal(x,y int64,fn func(int64,int64)int64)int64{
    a := x*10
    var b int64
    b = y*10
    return fn(a,b)
}
func main() {
    x :=cal(10,20,add)
    y := cal(10,20,mul)
    fmt.Print(x,y)
    
}
