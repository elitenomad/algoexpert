def lcm(a, b):
  l = max(a, b)
  while True:
    if l % a == 0 and l % b == 0:
      return l
    l += 1

  return l

def gcd(a, b):
  while a != b:
    if a > b:
      a = a - b
    elif b > a:
      b = b - a
    
  return a

def lcm_eff(a, b):
  return (a * b)//gcd(a, b)


if __name__ == "__main__":
  a = input("Enter a number :")
  b = input("Enter a number :")
  print(lcm(int(a), int(b)))
  print(lcm_eff(int(a), int(b)))