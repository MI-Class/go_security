package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup
var chs = make([]chan int, 10)

func Add(x, y int, yy chan int) {
	defer wg.Done()
	<-yy
	z := x + y
	fmt.Println(z)
	if x+1 < len(chs) {
		chs[x+1] <- 1
	}

}
func main() {
	a := Hh{5}
	//b := nr{5.0}
	// var i float32 = b // compile-error: cannot use b (type nr) as type float32 in assignment
	// var i = float32(b) // compile-error: cannot convert b (type nr) to type float32
	// var c number = b // compile-error: cannot use b (type nr) as type number in assignment
	// needs a conversion:
	//var c = number(b)
	fmt.Println(a)

}

type Hh struct {
	id int
}
type number struct {
	f float32
}
type nr number
