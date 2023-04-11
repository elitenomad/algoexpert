# On the assumption that input array size >= 3
def search_in_rotated_array(input, target):
  low = 0
  high = len(input) - 1
  
  while(low <= high):
    mid = low + (high - low) // 2
    
    if input[mid] == target:
      return mid
    
    if(input[low] < input[mid]):
      if(target >= input[low] and target < input[mid]):
        high = mid - 1
      else:
        low = mid + 1
    else:
      if(target > input[mid] and target <= input[len(input)-1]):
        low = mid + 1
      else:
        high = mid - 1
        
  return -1
# CASE-1 : For given mid, When value[mid] > value[mid + 1]
# CASE-2 : For given mid, When value[mid] < value[mid - 1]
# CASE-3 : For given mid, When value[low] > value[mid], Ignore all the elements right of mid, -> high = mid - 1
# CASE-4 : For given mid, When value[low] < value[mid], Ignore all the elements left of mid, -> low = mid 
def find_pivot(input, low, high):  
  if(low == high):
    return low
  
  mid = low + (high - low)//2
  
  if(mid < high and input[mid] > input[mid + 1]):
    return mid
  
  if(mid > low and input[mid] < input[mid - 1]):
    return mid - 1
  
  if(input[low] >= input[mid]):
    return find_pivot(input, low, mid - 1)
  else:
    return find_pivot(input, mid + 1, high)

def find_pivot_non_recursive(input, low, high):  
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
  

def find_pivot_with_duplicates(input):  
  low = 0
  high = len(input) - 1
  
  while(low <= high):
    mid = low + (high - low)//2
    if(mid < high and input[mid] > input[mid + 1]):
      return mid
    
    if(mid > low and input[mid] < input[mid - 1]):
      return mid - 1
    
    if(input[mid] == input[low] and (input[mid] == input[high])):
      if(input[low] > input[low + 1]):
        return low
      low += 1
      
      if(input[high] < input[high-1]):
        return high - 1
      high -= 1
    elif(input[low] < input[high] or (input[low] == input[mid] and input[mid] > input[high])):
      low = mid + 1
    else:
      high = mid - 1      
      
    
  return -1

def binary_search(input, target, low, high):  
  while(low <= high):
    mid = low + (high - low)//2
    
    if(input[mid] == target):
      return mid
    
    if(target >= input[mid]):
      low = mid + 1
    elif(target < input[mid]):
      high = mid - 1
    
  return -1


def search_in_rotated_array_II(input, target):
  low = 0
  high = len(input) - 1
  
  # FIND PIVOT
  pivot = find_pivot_non_recursive(input, low, high)
  print("PIVOT",pivot)
  
  if(pivot == -1):
    return binary_search(input, target, low, high)
  
  if(target <= input[pivot]):
    return binary_search(input, target, low, pivot)
  else:
    return binary_search(input, target, pivot + 1, high)

if __name__ == '__main__':
  i = [100, 500, 10, 20, 30, 40, 50]
  print(search_in_rotated_array_II(i, 500))
  print(search_in_rotated_array(i, 40))

