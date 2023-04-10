# On the assumption that input array size >= 3
def search_in_mountain_array(input, target):
  peak = peak_in_mountain_array(input)
  print("PEAK: ", peak)
  index = binary_search(input, target, 0, peak)
  print("INDEX: ", index)
  if(index == -1):
    index = binary_search(input, target, peak, len(input) - 1)
    print("INDEX: ", index)
  return index

def peak_in_mountain_array(input):
  low = 0
  high = len(input) - 1
  
  while(low < high):
    mid = low + (high - low)//2

    if(input[mid] > input[mid + 1]):
      high = mid
    else:
      low = mid + 1
      
  return low


def binary_search(input, target, low, high):   # ORDER AGNOSTIC BINARY SEARCH
  isAsc = input[low] < input[high]
  
  while(low <= high):
    mid = low + (high - low)//2
    
    if(isAsc):
      if(target < input[mid]):
        high = mid - 1
      elif(target > input[mid]):
        low = mid + 1
      else:
        return mid
    else:
      if(target < input[mid]):
        low = mid + 1
      elif(target > input[mid]):
        high = mid - 1
      else:
        return mid
    
  return -1

if __name__ == '__main__':
  i = [1,4,6,9,7,5,3,2]
  print(search_in_mountain_array(i, 7))

