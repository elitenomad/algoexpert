# Pattern matching

# Goal is to take the pattern and find it through the string given

# Naive : O(N-M+1 * M) - average  O(N) whwne all the chars in the pattern are differnt
  - N length of a string
  - M Length of a pattern
  - (N-M+1), Index until which you parse the string and look for pattern
  - Happens on the fly and no pre-processing

# Rabin Karp algorithm
  - O(N-M+1 * M)
  - Preprocessing pattern
  - Better than Naive pattern on an average
# KMP
  - O(N)
  - Preprocessing pattern
  
# Suffix tree
  - O(M)