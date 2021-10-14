// This is my solution for AlgoDaily problem Power of Three
// Located at https://algodaily.com/challenges/power-of-three

package main

import "fmt"

func powerOfThree(input int) bool {
  for input % 3 == 0 {
    input = input / 3
  }
    return input == 1
}

func main() {
    // write test cases
    fmt.Println(powerOfThree(91))
}

