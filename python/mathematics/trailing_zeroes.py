def trailing_zeroes(n):
  count = 0

  while n >= 5:
    n //= 5
    count += n
  return count


if __name__ == "__main__":
  n = input("Enter a number :")
  print(trailing_zeroes(int(n)))