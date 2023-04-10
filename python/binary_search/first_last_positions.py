# Find the first and last positions of a target
def first_position(input, target):
  low = 0
  high = len(input) - 1
  if(target > input[high] or target < input[low]):
    return - 1
  
  while(low <= high):
    # mid = (low + high)//2
    mid = low + (high - low)//2
    
    if(input[mid] < target):
      low = mid + 1
    elif(input[mid] > target):
      high = mid - 1
    else:
      if input[mid] != input[mid - 1]:
        return mid
      else:
        high = mid - 1
    
  return mid

# Find the first and last positions of a target
def last_position(input, target):
  low = 0
  high = len(input) - 1
  if(target > input[high] or target < input[low]):
    return - 1
  
  while(low <= high):
    # mid = (low + high)//2
    mid = low + (high - low)//2
    
    if(input[mid] < target):
      low = mid + 1
    elif(input[mid] > target):
      high = mid - 1
    else:
      if input[mid] != input[mid + 1]:
        return mid
      else:
        low = mid + 1
    
  return mid

if __name__ == '__main__':
  i = [1, 2, 3, 4, 5, 6, 10, 10, 12, 54, 64, 99]
  # i = [1, 2, 3, 4, 7, 10, 10, 10, 10, 54, 64, 99]
  print(first_position(i, 10), last_position(i, 10))
