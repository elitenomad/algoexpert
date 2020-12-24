# Initialize array
# Initialize alphabests -> convert them to array
# loop through the str chars
# Get each letter and add key
# Fetch the newLetter from alphabets array
# Join the result array and return.
def caesar_cipher(str, key)
    result = []
    alphabets = "abcdefghijklmnopqrstuvwxyz".split("")

    str.chars.each do |letter|
        newLetterCode = alphabets.index(letter) + key
        result.push(alphabets[newLetterCode % 26])
    end

    return result.join("")
end