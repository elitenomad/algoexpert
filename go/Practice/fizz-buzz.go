// This is my solution for AlgoDaily problem Fizz Buzz
// Located at https://algodaily.com/challenges/fizz-buzz

package main

import "fmt"

func fizzBuzz(arg int) string {
  d_3 := (arg % 3) == 0
  d_5 := (arg % 5) == 0
  
  if d_3 && d_5 {
    return "fizzbuzz"
  }
  
  if d_3 {
    return "fizz"
  }
  
  if d_5 {
    return "buzz"
  }
  
  return "output"
}

func main() {
    // write test cases
    fmt.Println(fizzBuzz(15))
}

