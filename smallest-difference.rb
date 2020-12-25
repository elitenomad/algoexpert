def smallest_difference(arr_1, arr_2) # O(M * N)
    h = {}

    arr_1.each do |e1|
        arr_2.each do |e2|
            h[(e1 - e2).abs] = [e1, e2]
        end
    end

    return h[h.keys.sort{|a, b| a - b }.first]
end