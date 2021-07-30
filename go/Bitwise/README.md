# Bitwise Operators in GO

## Number literals can be written in 3 forms
  - Decimals (never starts with 0)
  - Octals  (starts with 0 )
  - HexaDecimal (starts with 0x)

## Bitwise Operators
  - Bitwise AND &
  - Bitwise OR |
  - Bitwise NOT !
  - Bitwise XOR ^
  - Bitwise ANDNOT (Bit clear) &!
  - Bitwise Left Shift <<
  - Bitwise Right Shift >>

## Bitwise AND
  - 0 & 0 => 0
  - 0 & 1 => 0
  - 1 & 0 => 0
  - 1 & 1 => 1
  - Example: 0110 & 0101 -> 0100

## Bitwise OR
  - 0 & 0 => 0
  - 0 & 1 => 1
  - 1 & 0 => 1
  - 1 & 1 => 1
  - Example: 0110 & 0101 -> 0111

## Bitwise XOR
  - 0 ^ 0 => 0
  - 0 ^ 1 => 1
  - 1 ^ 0 => 1
  - 1 ^ 1 => 0
  - x ^ 0 => x
  - x ^ y => y ^ x
  - x ^ (y ^ z) => (x ^ y) ^ z
  - x ^ x => 0

## Bitwise Left Shift <<
  - 3 << 1 => 011 << 1 => 110 => 6
  - x * 2^y  

## Bitwise Right Shift >>
  - 3 >> 1 => 011 >> 1 => 001 => 1
  - 4 >> 1 => 100 >> 1 => 010 => 2
  - x / 2 ^ y

## Bitwise NOT
  - ^3 => ^011 =>100 => 4
  