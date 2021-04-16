package main

import "fmt"

func main() {
	fmt.Println("Start Fizz Buzz test")
	for i := 1; i <= 100; i++ {

		if i%15 == 0 {
			fmt.Println("FizzBuzz")
			continue
		}
		if i%3 == 0 {
			fmt.Println("Fizz")
			continue
		}
		if i%5 == 0 {
			fmt.Println("Buzz")
			continue
		}
		fmt.Println(i)

	}
	fmt.Println("End")
}
