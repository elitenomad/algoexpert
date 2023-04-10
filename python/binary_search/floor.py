# Find the TARGET in the sorted Array
# If you don't find it return the largest element smaller than target
def floor(input, target):
  if(target < input[0]):
    return -1
  
  low = 0
  high = len(input) - 1
  
  while(low <= high):
    mid = (low + high)//2
    if(input[mid] == target):
      return target
    
    if(input[mid] < target):
      low = mid + 1
    elif(input[mid] > target):
      high = mid - 1
    else:
      return input[mid]
  
  print(low, high)
  return input[high]

if __name__ == '__main__':
  i = [12,1,10,2,3,4,7,22,34,54,64,99,100]
  i.sort()
  # i = [1, 2, 3, 4, 7, 10, 12, 22, 34, 54, 64, 99, 100]
  # l = 0, high = 12, mid = 6, 12 > 5
  # l = 0, high = 6, mid = 3, 4 < 5
  # l = 3, high = 6, mid = 4, 7 > 5
  # l = 3, high = 3, mid = 3, 4 < 5
  # l = 3, high = 3, mid = 3, 4 < 5
  # l = 2, high = 5, mid = 3
  # l = 3, high = 5, mid = 4
  # l = 4, high = 5, mid = 4
  
  print(floor(i, 9))
  print(floor(i, 0))