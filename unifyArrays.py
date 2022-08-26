# solve takes as input a matrix containing n sorted arrays and
# has to return a unified sorted array composed only by the elements of
# the n sorted arrays

# Example:
#   Input:
#   [
#     [1, 5, 8],
#     [3, 4],
#     [4, 11]
#   ]

#   Output:
#   [1, 3, 4, 4, 5, 8, 11]



def solve(input):
  indexes = [0] * len(input)
  
  okay = True
  resulting_array = []

  while okay:
  
    minValue = int(1e9 + 7)
    index = 0
    for j in range(len(input)):
      if indexes[j] < len(input[j]) and input[j][indexes[j]] < minValue:
        index = j
        minValue = input[j][indexes[j]]
    resulting_array.append(minValue)
    indexes[index] += 1

    check = False
    for i in range(len(indexes)):
      if indexes[i] < len(input[i]):
        check = True
    
    okay = check

  return resulting_array



input = [
  [1, 2, 5, 8, 10, 11],
  [2, 3, 5, 6, 8, 9],
  [7, 10, 24, 59]
]

print(solve(input))