def palindrome(str)
    return str == str.reverse
end


def v2(str)
    i = 0
    j = str.length - 1

    while i < j  
        if str[i] != str[j]
            return false
        end

        i = i + 1
        j = j - 1
    end

    return true
end