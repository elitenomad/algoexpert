# Find the TARGET in the sorted Array
# If you don't find it return the next largest element of the target

def celing(input, target):
  low = 0
  high = len(input) - 1
  
  while(low < high):
    print(low, high)
    mid = (low + high)//2
    if(input[mid] == target):
      return target
    
    if(input[mid] < target):
      low = mid + 1
    elif(input[mid] > target):
      high = mid
      
  return input[high]

if __name__ == '__main__':
  i = [12,1,10,2,3,4,7,22,34,54,64,99,100]
  i.sort()
  # i = [1, 2, 3, 4, 7, 10, 12, 22, 34, 54, 64, 99, 100]
  print(celing(i, 5))