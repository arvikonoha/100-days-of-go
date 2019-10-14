package main

import (
	"fmt"
	"math"
	"strings"
)

type vertex struct {
	x, y int
}

func wordSmith(myString string) map[string]int {
	mp := make(map[string]int)
	for _, v := range strings.Fields(myString) {
		mp[v]++
	}
	return mp
}

func compute(fn func(float64, float64) float64) float64 {
	return fn(3, 4)
}

func adder() func(int) int {
	sum := 0 // this value is remembered by the inner functions of this outer function as and when they were created
	return func(a int) int {
		sum += a // every time function is called sum is added with the param and new value of sum is returned the value of sum associated with the givven inner function has nothing to do with the inner function created seperately
		return sum
	}
}

func fibonacci() func() int {
	x, y, z := 0, 1, 0
	return func() int {
		z = y + x
		x = y
		y = z
		return x
	}
}

func main() {
	i, j := 10, 20

	p := &i
	fmt.Println(*p) // p holds the pointer to i
	*p = 5          // value of i changes
	fmt.Println(i)

	// similarly...
	p = &j
	fmt.Println(*p)
	*p = j / 5
	fmt.Println(j)
	fmt.Println(vertex{3, 4}) // struct consists of collection of fields
	v := vertex{6, 7}
	v.x = 3 // fields are assigned/accessed via .
	fmt.Println(v.y)

	// array declaration
	var a [2]string
	a[0] = "Hello"
	a[1] = "World"
	fmt.Println(a[0], a[1])

	// another syntax
	b := []int{2, 4, 6, 8, 10}
	fmt.Println(b)

	sl := b[1:4] // slice , which is dynamic array, formed by 2 indices low and high
	fmt.Println(sl)

	// slices just references the underlying array
	// change in the slice actually affect the underlying array

	greetings := [6]string{"Hello", "Hi", "Good morning", "Hola", "Namaste", "Gracias"}
	sl1 := greetings[0:4]
	sl2 := greetings[3:] // default low bound is 0 and high bound is the length of the array
	sl1[3] = "Bye"
	fmt.Println(sl1, sl2, greetings)

	// length of a sclice is its current number of elements while capacity is the elements in its underlying array

	sl3 := []int{2, 4, 6, 7, 2, 1}

	sl3 = sl3[:0] // chopping higher bound does'nt reduce the capacity

	fmt.Println(sl3, " len - ", len(sl3), " cap - ", cap(sl3))

	sl3 = sl3[:4]

	fmt.Println(sl3, " len - ", len(sl3), " cap - ", cap(sl3))

	sl3 = sl3[2:] // chopping the lower bound does.

	fmt.Println(sl3, " len - ", len(sl3), " cap - ", cap(sl3))

	sl4 := [][]int{[]int{1, 2, 3}, []int{4, 5, 6}}
	fmt.Println(sl4)

	// there can be slices of slices as well

	sl5 := make([]int, 1)      // makes an arraqy of size 1 and returns a slice to it
	sl5 = append(sl5, 4, 5, 6) // appends the given set of elements to the given slice and returns the newly formed slice

	fmt.Println(sl5)

	//  map[type1]type2 maps the type1 key to type2 value, can be declared and assigned on the spot, avoid this dickhead give meaningful name instead

	mp1 := map[string]vertex{"Bell labs": vertex{x: 10, y: 20}, "Joe's Pizza": vertex{x: 24, y: 69}}
	fmt.Println(mp1)

	// just like for each in JavaScript

	for i, v := range mp1 {
		fmt.Println("KEY - ", i, " VALUE - ", v)
	}

	mp2 := make(map[string]int) // always use atleast make construct to make a map

	mp2["Answer"] = 10

	fmt.Println(mp2)

	delete(mp2, "Answer") // gets rid of the key answer

	answer, ok := mp2["Answer"] // ok gets the boolean true if given key answer is present and answer gets the possible value, zero value otherwise

	fmt.Println("isAlive - ", ok, " answer - ", answer)

	// function to count the number of occurance of each words in a string

	fmt.Println(wordSmith("Hi hello bye hello hi good great bye"))

	// the function compute takes the function of the type func(float64,float64) float64 and returns the value of the given function run with arguements 3,4

	fmt.Println(compute(math.Pow)) //returns 3**4

	sum := adder()
	for i := 0; i < 10; i++ {
		fmt.Println(sum(i)) // this keeps adding i to sum remembered by the inner function of adder returned by it
	}

	// the fibonacci function returns a function which on running each time returns the next fibonacci number this works based on the fact that outer function's variables are remembered by the inner function returned and on each time of inner function call variables are modified as per fibonacci series and number returned corresponds to next element in the fibonacci series

	fibo := fibonacci()

	for i := 0; i < 10; i++ {
		fmt.Println(fibo())
	}

}
