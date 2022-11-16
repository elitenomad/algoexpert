def prime(n):
  if n == 1:
    return False

  if n == 2 or n == 3 or n == 5 or n == 7:
    return True

  for i in range(2, n):
    if n % i == 0:
      return False

  return True


def prime_eff(n):
  if n == 1:
    return False
  
  i = 2
  while i * i <= n:
    if n % i == 0:
      return False
    i += 1
  
  return True

if __name__ == "__main__":
  n = input("Enter a number :")
  print(prime(int(n)))
  print(prime_eff(int(n)))