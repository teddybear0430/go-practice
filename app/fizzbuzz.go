package main

import "fmt"

func fizzbuzz(fizzBuzzNum int) {
	count := 1

	for count <= fizzBuzzNum {
		switch {
		case count % 2 == 0 && count % 7 == 0:
			fmt.Printf("%d is a fizzbuzz\n", count)
		case count % 2 == 0:
			fmt.Printf("%d is a fizz\n", count)
		case count % 7 == 0:
			fmt.Printf("%d is a buzz\n", count)
		default:
			fmt.Println(count)
		}

		count++
    }
}

func main() {
	fizzbuzz(100)
}
