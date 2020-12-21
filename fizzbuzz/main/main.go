package main

import "fmt"

// Write a program that prints the numbers from 1 to 100.
// But for multiples of three print “Fizz” instead of the number and for the multiples of five print “Buzz”.
// For numbers which are multiples of both three and five print “FizzBuzz”
func main(){
	i:= 1
	for i <= 100{
		isFizzOrBuzz := false
		//fmt.Sprintf("i: %d -", i)
		if i % 3 == 0{
			fmt.Print("fizz")
			isFizzOrBuzz = true
		}
		if i % 5 == 0{
			fmt.Print("buzz")
			isFizzOrBuzz = true
		}
		if !isFizzOrBuzz{
			fmt.Print(i)
		}
		fmt.Print("\n")
		i++
	}
}