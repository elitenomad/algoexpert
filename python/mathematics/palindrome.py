def palindrome(n):
  original = n
  rev = 0
  while n > 0:
    rem = n % 10
    rev = rev * 10 + rem
    n = n // 10
  return rev == original


if __name__ == "__main__":
  n = input("Enter a number: ")
  print(palindrome(int(n)))


