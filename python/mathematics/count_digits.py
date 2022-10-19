def count_digits(n): # O(N)
  count = 0
  while n > 0:
    count += 1
    n = n // 10
  return count

if __name__ == '__main__':
  n = input("Enter a number: ")
  print(count_digits(int(n)))