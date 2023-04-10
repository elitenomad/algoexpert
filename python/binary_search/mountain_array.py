# On the assumption that input array size >= 3
def peak_in_mountain_array(input):
  low = 0
  high = len(input) - 1
  
  while(low < high):
    mid = low + (high - low)//2

    if(input[mid] > input[mid + 1]):
      high = mid
    else:
      low = mid + 1
      
  return input[low]

if __name__ == '__main__':
  i = [0,1,0]
  print(peak_in_mountain_array(i))
  i = [0,1,2,3,4,5,4,3,2,1,0]
  print(peak_in_mountain_array(i))
  i = [0,1,2,3,4,5]
  print(peak_in_mountain_array(i))
  i = [0,1,2,1,0]
  print(peak_in_mountain_array(i))
  i = [1,2,3,5,7,6,3,2]
  print(peak_in_mountain_array(i))
  i = [0,2,1,0]
  print(peak_in_mountain_array(i))

