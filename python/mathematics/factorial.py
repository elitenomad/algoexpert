def factorial(n):
  product = 1
  for i in range(1, n + 1):
    product *= i
  
  return product


def factorial_rec(n):
  if n == 1:
    return 1
  
  return n * factorial_rec(n - 1)

if __name__ == "__main__":
  print(factorial(5))
  print(factorial_rec(5))