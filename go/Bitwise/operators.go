package main

import "fmt"

func main() {
	a := -3                                   // 0011
	b := 6                                    // 0110
	var c uint8 = 5                           // 2^8 - 2
	var d int8 = 5                            // 2^8 - 2
	fmt.Println("Bitwise AND:", a&b)          // 2
	fmt.Println("Bitwise OR:", a|b)           // 7
	fmt.Println("Bitwise XOR:", a^b)          // 5
	fmt.Println("Bitwise Left Shift:", a<<1)  // 6
	fmt.Println("Bitwise Right Shift:", a>>1) // 011 >> 1 => 001 => 1
	fmt.Println("Bitwise NOT: ", ^c)          // ^011 => 100 => 4
	fmt.Println("Bitwise NOT: ", ^d)          // ^011 => 100 => 4
	fmt.Printf("%b ", b)
}
