def IsPrime(n):
  if n == 1:
    return False

  if n == 2 or n == 3 or n == 5 or n == 7:
    return True

  for i in range(2, n):
    if n % i == 0:
      return False

  return True

def prime_factors(n):
  for i in range(2, n + 1):
    if IsPrime(i):
      x = i
      while n % x == 0:
        print(i)
        x = x * i

if __name__ == "__main__":
  n = input("Enter a number :")
  print(prime_factors(int(n)))
