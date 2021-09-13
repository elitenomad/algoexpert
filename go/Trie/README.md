  # Trie Datastructure

  - Its best suited for implementing Dictionaries
  - Dictionaries which contains set of words and keen to perform following operations
    - Insert
    - Search
    - Delete
    - PrefixSearch
    - Lexicographical ordering of words
  
e.g: Insert("Geek"), Insert("Geeks"), Insert("Bad"), Insert("Bat")
     Search("Bad"), Search("eek"), Delete("Geeks"), Search("Geeks")
     Print()
O/P: Yes, No, No (After Delete)
     "Bad", "Bat", "Geeks"


# Comparison of Hashing and Trie

- One of the key characteristics of Hashing is Fast Search, Insertion and Deletion operations

Search                  ----------> O(word-length) worst case --- O(word-length) average case
Insert                  ----------> O(word-length) worst case --- O(word-length) average case
Delete                  ----------> O(word-length) worst case --- O(word-length) average case
Prefix Search           ----------> O(p-length + O/P length)  --- Not Supported
Lexicographical Ordering----------> O(output-length)          --- Not Supported