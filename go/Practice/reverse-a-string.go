// This is my solution for AlgoDaily problem Reverse a String
// Located at https://algodaily.com/challenges/reverse-a-string

package main

import "fmt"

func reverseAString(arg string) string {
  r := []rune(arg)
  
  for i,j := 0, len(r)-1; i < len(r)/2; i,j = i+1,j-1 {
   r[i],r[j] = r[j], r[i]
  }
  return string(r)
}

func main() {
    // write test cases
  
    fmt.Println(reverseAString("input"))
  fmt.Println(reverseAString("topper"))
  fmt.Println(reverseAString("i"))
  fmt.Println(reverseAString("it"))
  fmt.Println(reverseAString("NIT"))
  fmt.Println(reverseAString(" "))
}

