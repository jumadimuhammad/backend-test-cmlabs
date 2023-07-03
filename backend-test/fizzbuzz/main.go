package main

import "fmt"

func main() {
	fizzBuzz(100)
}

func fizzBuzz(num int) {
	var fizz = "Fizz"
	var buzz = "Buzz"
	
	for i := 1; i <= num; i++ {
		if i%3 == 0 && i%5 == 0 {
			fmt.Println(i, "--> ", fizz+buzz)
		} else if i%3 == 0 {
			fmt.Println(i, "--> ", fizz)
		} else if i%5 == 0 {
			fmt.Println(i, "--> ", buzz)
		} else {
			fmt.Println(i)
		}

	}
}
