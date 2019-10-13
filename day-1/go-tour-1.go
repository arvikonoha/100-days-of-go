package main

import (
	"fmt"
	"math"
	"math/rand"
)

var isAlive bool
var num1, num2 int = 2, 3

func add(a int, b int) int {
	return a + b
}

func multi(a, b int) int {
	return a * b
}

func greet() (string, string) {
	return "hello", "world"
}

func arith(a, b int) (sum, mul int) {
	sum, mul = add(a, b), multi(a, b)
	return
}

func main() {

	var myNumber int

	fmt.Println("Hello world")
	fmt.Printf("A random number %d", rand.Intn(20))  // package example
	fmt.Println("Sqare root of 9 is ", math.Sqrt(9)) // look at how you could structure imports
	fmt.Println("Value of Pi is ", math.Pi)          // exported entities have a capitalized first letter
	fmt.Println(add(10, 20))                         // example for function syntax
	fmt.Println(multi(10, 20))
	fmt.Println(greet())           // multiple function returns
	fmt.Println(arith(10, 20))     // named return aka naked return
	fmt.Println(isAlive, myNumber) // variable declaration using var , default values for variables is zero variable
	fmt.Println(num1, num2)        // variable declaration along with initialization
	fact, right := "Javascript is awesome", true
	fmt.Println(fact, right)
	// all datatypes -> int,int8,int16,int32,int64,rune,byte,uint8,uint16,uint32,uint64,float,float32,float64,complex64,complex128
	comp := 5 + 6i
	fmt.Printf("%T\n", comp) // %T is used to check type of variable
	var i32 int32 = 10
	var i64 int64 = int64(i32)
	fmt.Println("ugh...", i64) // though int32 is the smaller type it must still be converted cause like java go doesn't fuck arount when it comes to types
	const myName = "Aravind"
	fmt.Println("My name won't change", myName) // constant is declared with const keyword and must be immediately assigned
}
