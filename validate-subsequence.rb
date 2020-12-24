def validate_subsequence(input, sub)
    if sub.length > input.length
        return false
    end

    sequence = 0
    input_len = 0

    while sequence < sub.length && input_len < input.length
        if sub[sequence] == input[input_len]
            sequence = sequence + 1
        end

        input_len = input_len + 1
    end

    return sequence == sub.length
end