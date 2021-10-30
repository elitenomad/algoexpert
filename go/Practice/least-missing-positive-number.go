// This is my solution for AlgoDaily problem Least Missing Positive Number
// Located at https://algodaily.com/challenges/least-missing-positive-number

package main

import "fmt"

func leastMissingPositiveNumber(input []int) int {
    // First step would be to find min and max values
  min := input[0]
  max := input[0]
  for _, v := range input { // O(N)
    if min > v {
      min = v
    }
    
    if max < v {
      max = v
    }
  }
  fmt.Println(min, max)
  
  result := 0
  for _, v := range input { // O(N)
    result = result ^ v
  }
  
  for i := min; i <= max; i++ { // O(N)
    result = result ^ i
  }
  
  
  return result
}

func main() {
    // write test cases
  fmt.Println(leastMissingPositiveNumber([]int{0,3,-1,-2,1}))
}

