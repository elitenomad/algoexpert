# Input: [5, 2, [7, -1], 3, [6, [-13, 8], 4]]
# Output: 12
# []: Level is 1, [[]]: Level is 2, [[[]]]: Level is 3
def product_sum(array)
    return product_sum_helper(array, 1)
end

def product_sum_helper(array, level)
    acc = 0

    array.each do |elem|
       if elem.is_a?(Array)
            acc += product_sum_helper(elem, level + 1)
        else
            acc += elem
       end 
    end

    return acc * level
end