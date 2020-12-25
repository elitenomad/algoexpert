def monotonic(array)
    increasing = true
    decreasing = true

    i = 1
    while i < array.length
        if(array[i] < array[i - 1])
            increasing = false   
        end

        if(array[i] > array[i - 1])
            decreasing = false
        end

        i = i + 1
    end

    return increasing || decreasing
end