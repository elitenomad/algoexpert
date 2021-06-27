# O(nlogn) and O(n) space
def sorted_square(input)
    result = []

    input.each_with_index do |val, i|
        result[i] = val * val
    end

    return result.sort
end