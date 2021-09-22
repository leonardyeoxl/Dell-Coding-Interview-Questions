# Time Complexity: O(n+(n-1)+(n-2)+(n-3)+(n-4)). Given n is the number of elements.
# Space Complexity; O(1)
def Solution(input):
    temp = -1
    for index in range(len(input)):
        if input[index] >= temp:
            temp = input[index]
    input.remove(temp)

    temp = -1
    for index in range(len(input)):
        if input[index] >= temp:
            temp = input[index]
    input.remove(temp)

    temp = -1
    for index in range(len(input)):
        if input[index] >= temp:
            temp = input[index]
    input.remove(temp)

    temp = -1
    for index in range(len(input)):
        if input[index] >= temp:
            temp = input[index]
    input.remove(temp)

    temp = -1
    for index in range(len(input)):
        if input[index] >= temp:
            temp = input[index]

    return temp


input = [0,4,2,3,1,5,10,9]
assert Solution(input) == 3, "Wrong answer"

input = [10,6,4,5,0,7,9,1]
assert Solution(input) == 5, "Wrong answer"