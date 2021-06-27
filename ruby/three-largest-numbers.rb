# The way i want to approach the problem is 
# Empty result [nil, nil, nil]
# Loop through the array
# update result for each iteration
# manually verify the index position values wih array iteration value
# Replace if its greater or continue if its not
# At the end of the array iteration we will have a result array with 3 largest values.
def find_three_largest_numbers(array)
    result = [nil, nil, nil]

    array.each do |elem|
        update_result_with_largest_nums(result, elem)
    end

    return result
end


def update_result_with_largest_nums(result, num)
    if result[2].nil? || num > result[2] 
        shift_and_update(result, num, 2)
    elsif result[1].nil? || num > result[1] 
        puts shift_and_update(result, num, 1)
    elsif result[0].nil? || num > result[0]
        shift_and_update(result, num, 0)
    end
end

def shift_and_update(result, num, index)
    (index + 1).times.each do |i|
        if i == index
            result[i] = num
        else
            result[i] = result[i + 1]
        end
    end
end