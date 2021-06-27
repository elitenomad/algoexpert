# Goal is to find all the combinations of the target
# from the array of elements (triplets).
def three_sum(array, target)
    # Sort the array
    # We gonna ignore O(nlogn) complexity anyways
    # as there are two nested loops which constitute for O(n^2)
    array.sort!{|a, b| a - b } # [2,4,3,1] => [1,2,3,4]

    triplets = []
    i = 0

    while i < array.length - 2
        left = i + 1
        right = array.length - 1

        while left < right
            sum = array[i] + array[left] + array[right]
            
            if sum == target
                triplets.push([array[i], array[left], array[right]])
                
                left = left + 1
                right = right - 1
            elsif sum < target
                left = left + 1
            elsif sum > target
                right = right - 1
            end
        end

        i = i + 1
    end 

    return triplets
end