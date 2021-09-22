package main

import (
	"fmt"
)

func remove(slice **[]int, s int) []int {
    return append((**slice)[:s], (**slice)[s+1:]...)
}

func solution(input *[]int) int {
	temp := -1
	temp_index := 0
	for index := 0; index < len((*input)); index++ {
		if (*input)[index] >= temp {
			temp = (*input)[index]
			temp_index = index
		}
	}
	(*input) = remove(&input, temp_index)

	temp = -1
	temp_index = 0
	for index := 0; index < len((*input)); index++ {
		if (*input)[index] >= temp {
			temp = (*input)[index]
			temp_index = index
		}
	}
	(*input) = remove(&input, temp_index)

	temp = -1
	temp_index = 0
	for index := 0; index < len((*input)); index++ {
		if (*input)[index] >= temp {
			temp = (*input)[index]
			temp_index = index
		}
	}
	(*input) = remove(&input, temp_index)

	temp = -1
	temp_index = 0
	for index := 0; index < len((*input)); index++ {
		if (*input)[index] >= temp {
			temp = (*input)[index]
			temp_index = index
		}
	}
	(*input) = remove(&input, temp_index)

	temp = -1
	for index := 0; index < len((*input)); index++ {
		if (*input)[index] >= temp {
			temp = (*input)[index]
		}
	}

	return temp
}

func main() {
	input := []int{0,4,2,3,1,5,10,9}
	result := solution(&input)
	fmt.Println(result)

	input = []int{10,6,4,5,0,7,9,1}
	result = solution(&input)
	fmt.Println(result)
}