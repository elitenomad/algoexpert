# Find the TARGET in the sorted Array (chars)
# If you don't find it return the next smallest element of the target (ignore equal to target)
# Return wrap around element if you dont find the target
def celing(input, target):
  low = 0
  high = len(input) - 1
  
  while(low <= high):
    # mid = (low + high)//2
    mid = low + (high - low)//2
    
    if(input[mid] < target):
      low = mid + 1
    else:
      high = mid - 1
  
  print(low, high)
  return input[low % len(input)]

if __name__ == '__main__':
  i = ['a', 'i', 'j']
  i.sort()
  # i = [1, 2, 3, 4, 7, 10, 12, 22, 34, 54, 64, 99]
  print(celing(i, 'a'))