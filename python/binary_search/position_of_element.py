# Find the position of an element in infinite array
def binary_search(input, target, low, high):
  while(low <= high):
    # mid = (low + high)//2
    mid = low + (high - low)//2
    
    if(input[mid] < target):
      low = mid + 1
    elif(input[mid] > target):
      high = mid - 1
    else:
      return mid
    
  return -1

def find_range(input, target):
  start = 0
  end = 1
  
  while(target > input[end]):
    temp = end + 1
    end = end + (end - start + 1) * 2
    start = temp
  
  return [start, end]

def find_element_in_infinite_array(input, target):
  indexes = find_range(input, target)
  index = binary_search(input, target, indexes[0], indexes[1])
  return index

if __name__ == '__main__':
  i = [i + 2 for i in range(100000000)]
  # i = [1, 2, 3, 4, 7, 10, 12, 22, 34, 54, 64, 99]
  print(find_element_in_infinite_array(i, 233))
  # print(find_element_in_infinite_array(i, 101))