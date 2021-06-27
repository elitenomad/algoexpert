def array_of_products(array)
    # Brute force
    products = []
    array.each_with_index do |val, i|
        array.each_with_index do |val2, j|
            if i != j
                products << va1 * val2
            end
        end
    end

    return products
end

# When the elements are non-zero
# Easy way of solving the problem but has its limitations
def v2(array)
    product = 1

    array.each do |elem|
        product *= elem
    end

    array.each_with_index do |elem, index|
        array[index] = product / elem
    end

    return array
end

# Get the product of the elements to the left of the array
# Get the product of the elements to the right of the array
# Multiply above two arrays by index
def v3(array)
    products = Array.new(array.length).fill(1)


    left = 1
    array.each_with_index do |elem, index|
        products[index] = left
        left *= elem
    end

    right = 1
    j = array.length - 1
    while j >= 0
        products[j] *= right
        right *= array[j]

        j = j -  1
    end

   return products
end


