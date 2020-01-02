package others

import (
	"fmt"
	"strings"
)

func Match(head, tail string) bool {
	if head == tail {
		return false
	}
	indexHead := strings.LastIndex(head, " ")
	indexTail := strings.Index(tail, " ")
	var headLastWord, tailFirstWord string
	if indexHead == -1 {
		headLastWord = head
	} else {
		headLastWord = head[indexHead+1 : len(head)]
	}
	if indexTail == -1 {
		tailFirstWord = tail
	} else {
		tailFirstWord = tail[0:indexTail]
	}
	if headLastWord == tailFirstWord {
		return true
	}
	return false
}

func buildMatrix(strs []string) [][]int {
	result := [][]int{}
	for i := 0; i < len(strs); i++ {
		mid := []int{}
		for j := 0; j < len(strs); j++ {
			if i != j && Match(strs[i], strs[j]) {
				mid = append(mid, 1)
			} else {
				mid = append(mid, 0)
			}
		}
		result = append(result, mid)
	}
	return result
}

func FindLongestPathInGraph(strs []string) int {
	matrix := buildMatrix(strs)
	outMatrix(matrix)

	resultLen := 0
	for i := 0; i < len(strs); i++ {
		length := 0
		result := []int{}
		for i := 0; i < len(strs); i++ {
			result = append(result, 0)
		}
		FindLongestPathForVertex(matrix, i, &result, &length)
		if length > resultLen {
			resultLen = length
		}
	}

	return resultLen

}
func FindLongestPathForVertex(gragh [][]int, i int, result *[]int, length *int) {
	(*result)[i] = 1
	for j := 0; j < len(gragh); j++ {
		var tempMax int
		if i != j && (*result)[j] == 0 && gragh[i][j] == 1 {
			(*result)[j] = 1
			tempMax = getLength(*result)
			if tempMax > *length {
				*length = tempMax
			}
			FindLongestPathForVertex(gragh, j, result, length)
			tempMax = getLength(*result)
			if tempMax > *length {
				*length = tempMax
			}
			(*result)[j] = 0

		} else {
			tempMax = getLength(*result)
			if tempMax > *length {
				*length = tempMax
			}
		}

	}
	(*result)[i] = 0
}
func outMatrix(matrix [][]int) {
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[i]); j++ {
			fmt.Printf("%3d", matrix[i][j])
		}
		fmt.Printf("\n")
	}
}
func getLength(a []int) int {
	length := 0
	for i := 0; i < len(a); i++ {
		if a[i] == 1 {
			length++

		}
	}
	return length
}
