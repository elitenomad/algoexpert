// This is my solution for AlgoDaily problem Dollar Sign Deletion
// Located at https://algodaily.com/challenges/dollar-sign-deletion

package main

import "fmt"

type Stack struct {
  Elements []rune
}

func New() *Stack {
  return &Stack{}
}

func (s *Stack) Push(element rune) {
  s.Elements = append(s.Elements, element)
}

func (s *Stack) Pop() rune {
  poppedElement := s.Elements[len(s.Elements)-1]
  s.Elements = s.Elements[0:len(s.Elements)-1]
  return poppedElement
}

func dollarSignDeletionHelper(s string) string{
  stack := New()
  fmt.Println(s)
  for i := 0; i < len(s); i++ {
    if rune(s[i]) != rune('$') {
      stack.Push(rune(s[i]))
    }else{
      stack.Pop()
    }
  }
  fmt.Println(stack)
  return string(stack.Elements)
}

func dollarSignDeletion(input []string) bool {
  // Loop through array
  // Pick a string, Delete previous char if you find a $ sign
  result := []string{}
  for i:= 0; i < len(input);i++ {
    result = append(result, dollarSignDeletionHelper(input[i]))
  }
    // Create a hash to store the frequencies of final string array
  h := map[string]int{}
  fmt.Println(result)
  for i := 0; i < len(result); i++ {
    h[result[i]] += 1
  }
  fmt.Println(h)
  return len(h) == 1
}

func main() {
    // write test cases
  input := []string{"f$st", "st" }
  fmt.Println(dollarSignDeletion(input))
}

