def move_element_to_end(array, move)
    # Goal is to move the 'move' to the end of an array
    start = 0
    finish = array.length - 1

    while start < finish
        while(start < finish &&  array[finish] == move)
            finish = finish - 1
        end
        
        if array[start] == move
            tmp = array[start]
            array[start] = array[finish]
            array[finish] = tmp
        end 

        start = start + 1
    end

    return array
end