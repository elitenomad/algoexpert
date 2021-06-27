# Input: [0, 1, 21, 33, 45, 45, 61, 71, 72, 73]
# Output: 33
def binary_search(array, target)
    start = 0
    finish = array.length - 1
    mid = (start + finish) / 2
   
    while start <= finish
        if array[mid] == target
            return mid     
        end

        if array[mid] > target 
            finish = mid - 1
        elsif array[mid] < target
            start = mid + 1
        end

        mid = (start + finish) / 2
    end

    return -1
end