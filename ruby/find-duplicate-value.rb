#Goal of this function is to find the first duplicate in the array
# and return its value
# Below algo will use O(n) for both time and space complexity
def find_duplicate(array)
    h = {}

    array.each do |elem|
        if h[elem]
            return elem
        else
            h[elem] = true
        end
    end

    return -1
end