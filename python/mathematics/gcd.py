def gcd(a, b):
  if a == 0: 
    return b
  
  if b == 0:
    return a

  num = max(a, b)

  for i in range(1, num + 1):
    if a % i == 0 and b%i == 0:
      gcd = i

  return gcd

def gcd_euclidian(a, b):
  while a != b:
    if a > b:
      a = a - b
    elif b > a:
      b = b - a
  return a


if __name__ == "__main__":
  a = input("Enter A: ")
  b = input("Enter B: ")
  print(gcd(int(a), int(b)))
  print(gcd_euclidian(int(a), int(b)))