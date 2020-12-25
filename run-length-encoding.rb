# Goal of this problem is to encode a string in such a way
# We can decode that as well with out ambiguity
# Rules are like 'AAA' => '3A'
def run_length_encoding(str)
    # Initialize encodingChars array
    # go through the all the chars of the string
    # push the counr and char value to chars array if previous wont match with current
    # If not continue

    i = 1
    chars = []
    len = 1

    while i < str.length
        current_c = str[i]
        previous_c = str[i - 1]

        if current_c != previous_c || len === 9
            chars.push(len.to_s)
            chars.push(previous_c)
            len = 0
        end

        i = i + 1
        len = len + 1
    end

    # Last Index Validation
    chars << len.to_s
    chars << str[str.length - 1]

    return chars.join('')
end