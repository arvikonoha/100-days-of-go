package main

import (
	"fmt"
	"math"
	"runtime"
	"time"
)

func sqrRoot(a float64) string {
	// if construct example
	if a < 0 {
		return sqrRoot(-a) + "i"
	}
	return fmt.Sprint(math.Sqrt(a))
}

func pow(x, n, lim float64) float64 {
	if v := math.Pow(x, n); v < lim {
		return v
	}
	return lim
}

func sqrRootAc(a float64) float64 {
	z, d := float64(1), float64(1)
	for d > 1e-15 {
		z0 := z
		z = z - (z*z-a)/(2*z)
		d = math.Abs(z - z0)
	}
	return z
}

func main() {
	sum := 1
	fmt.Println("For loop illustration")
	// for i := 0; i <= 10; i++ {
	// 	sum += i
	// }
	for sum < 100 {
		sum += sum
	}
	fmt.Println(sum)
	fmt.Println(sqrRoot(-10), " square root of a negative number") // illustration of if construct
	fmt.Println(pow(3, 2, 10))                                     // if with initialization
	fmt.Println(pow(3, 3, 20))
	fmt.Println(sqrRootAc(3)) // illustration of how math.Sqrt works
	// switch case
	switch os := runtime.GOOS; os {
	case "darvin":
		fmt.Println("darwin.")
	case "linux":
		fmt.Println("linux")
	default:
		fmt.Println(os)
	}

	// switch construct can have expression in its case which is executed in the order top to bottom
	fmt.Println("When is saturday ?")
	switch today := time.Now().Weekday(); time.Saturday {
	case today:
		fmt.Println("Today is saturday")
	case today + 1:
		fmt.Println("Tomorrow is saturday")
	case today + 2:
		fmt.Println("Saturday is two days away")
	default:
		fmt.Println("Saturday is too far away")
	}

	// better way of writing if else construct , the case that results in true is excuted

	fmt.Println("Greetings")
	switch {
	case time.Now().Hour() < 12:
		fmt.Println("Good morning")
	case time.Now().Hour() < 17:
		fmt.Println("Good afternoon")
	default:
		fmt.Println("Good night")
	}

	defer fmt.Println("world") // defer statement is executed after the control reaches the end of the current function

	fmt.Println("Hello")

	// defer makes the function to be pushed to a stack, hence they execute in last in first out manner

	for i := 0; i < 10; i++ {
		defer fmt.Println(i)
	}

}
