# for every loop we push the largest one to the end
def bubble_sort(array)
    sorted = false
    len = array.length - 1

    while !sorted
        sorted = true
        len.times do |index|
            if array[index] > array[index + 1]
                temp = array[index]
                array[index] = array[index + 1]
                array[index + 1] = temp
                sorted = false
            end
        end

        len = len - 1
    end

    array
end