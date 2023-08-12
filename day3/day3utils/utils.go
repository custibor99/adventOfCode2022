package utils

import "fmt"

func StrToIntList(s string) []int {
	res := make([]int, len(s))
	for i, el := range  s {
		res[i] = int(el)
	}
	return res
}

func MergeSort(elements []int ) []int {
	if len(elements) < 2 {
		return elements
	}

	n := len(elements)
	half := n / 2
	sortedLeft := MergeSort(elements[:half])
	sortedRight := MergeSort(elements[half:])
	sorted := make([]int, n)
	indxL := 0
	indxR := 0
	for {
		if indxL >= len(sortedLeft) {
			copy(sorted[indxL + indxR:],sortedRight[indxR:])
			return sorted
		}
		if indxR >= len(sortedRight) {
			copy(sorted[indxL + indxR:],sortedLeft[indxL:])
			return sorted
		}
		if sortedLeft[indxL] <= sortedRight[indxR] {
			sorted[indxL + indxR] = sortedLeft[indxL]
			indxL += 1
		} else {
			sorted[indxL + indxR] = sortedRight[indxR]
			indxR += 1
		}
	}
}

func FindDuplicateElements(el1 []int, el2 []int) []int {
	el1 = MergeSort(el1)
	el2 = MergeSort(el2)
	duplicates := make([] int, 1)
	indx1 := 0
	indx2 := 0
	for {
		if indx1 >= len(el1) ||  indx2 >= len(el2) {
			return duplicates[1:]
		}
		if el1[indx1] == el2[indx2] {
			indx2 += 1
			if len(duplicates) != 0 && duplicates[len(duplicates) - 1] != el1[indx1] {
				duplicates = append(duplicates, el1[indx1])
			}
		} else if el1[indx1] < el2[indx2] {
			indx1 += 1
		} else {
			indx2 += 1
		}
	}
}

func ScoreElements(input []int) int {
	score := 0
	for _, el := range input {
		if el >= int('a') && int(el) <= int('z'){
			score += el - int('a') + 1
		} else {
			score += el - int('A') + 27
		}
	}
	return score
}

func Hello() {
	fmt.Println("Hello world!")
}