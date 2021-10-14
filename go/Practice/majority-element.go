// This is my solution for AlgoDaily problem Majority Element
// Located at https://algodaily.com/challenges/majority-element

package main

import "fmt"

func majorityElement(input []int) int {
  h := map[int]int{}
  
  for i := 0; i < len(input); i++ {
    h[input[i]] += 1
  }
  
  for k, v := range h {
    if v > len(input)/2.0 {
      return k
    }
  }

  return -1
}

func main() {
  fmt.Println(majorityElement([]int{3,2,3}))
}

