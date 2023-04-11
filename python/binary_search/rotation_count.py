def find_pivot(input, low, high):  
  if(low == high):
    return low
  
  while(low <= high):
    mid = low + (high - low)//2
    if(mid < high and input[mid] > input[mid + 1]):
      return mid
    
    if(mid > low and input[mid] < input[mid - 1]):
      return mid - 1
    
    if(input[low] >= input[mid]):
      high = mid - 1
    else:
      low = mid + 1
    
  return -1


def rotation_count(input):
  low = 0
  high = len(input) - 1
  
  pivot = find_pivot(input, low, high)
  
  return pivot + 1

if __name__ == '__main__':
  i = [100, 500, 10, 20, 30, 40, 50]
  print(rotation_count(i))

