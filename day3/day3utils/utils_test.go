package utils

import (
	"math/rand"
	"sort"
	"testing"

	"golang.org/x/exp/slices"
)

func TestMergeSortEmpty(t *testing.T) {
	expected := [] int {}
	sorted := MergeSort(expected)

	if len(sorted) != 0 {
		t.Fatal("Array should be empty")
	}
}

func TestMergeSortShort(t *testing.T) {
	expected := [] int {1, 2, 3, 4}
	sorted := MergeSort(expected)

	for i := 0; i < len(expected); i++ {
		if expected[i] != sorted[i] {
			t.Fatal("Elements dont match", expected[i], " != ", sorted[i])
		}
	}
}


func TestMergeSortLong(t *testing.T) {
	for i := 0; i <= 4; i++ {
		expected := rand.Perm(100000)
		actuall := MergeSort(expected)
		sort.Ints(expected)
	
		for i := 0; i < len(expected); i++ {
			if expected[i] != actuall[i] {
				t.Fatal("Elements dont match", expected[i], " != ", actuall[i])
			}
		}
	}
}

func TestFindDuplicatesEmpty(t *testing.T) {
	list1 := [] int {1, 2, 3}
	list2 := [] int {4, 5, 6}

	res := FindDuplicateElements(list1, list2)
	if len(res) != 0 {
		t.Fatal("There are no intersectin elements")
	}
}

func TestFindDuplicate(t *testing.T) {
	list1 := [] int {1, 2, 3, 4}
	list2 := [] int {3, 2, 6, 7}

	res := FindDuplicateElements(list1, list2)
	if len(res) != 2 {
		t.Fatal("Intersection should contain 2 elements. It contains: ", res)
	}
	if !slices.Contains(res, 3) {
		t.Fatal("Does not contain element 3")
	}
	if !slices.Contains(res, 2) {
		t.Fatal("Does not contain element 2")
	}
}


func TestFindDuplicate1(t *testing.T) {
	list1 := [] int {1, 1, 2, 2}
	list2 := [] int {2, 2, 3, 3}

	res := FindDuplicateElements(list1, list2)
	if len(res) != 1 {
		t.Fatal("Intersection should contain 2 elements. It contains: ", res)
	}
}

func TestFindDuplicate2(t *testing.T) {
	list1 := [] int {90, 100, 114, 90, 115, 106, 81, 98, 110, 77}
	list2 := [] int {67, 71, 119, 114, 82, 119, 122, 109, 84, 114}

	res := FindDuplicateElements(list1, list2)
	if len(res) != 1 {
		t.Fatal("Intersection should contain 1 element. It contains: ", res)
	}
}

func TestScoreElementsEmpty(t *testing.T) {
	res := ScoreElements(StrToIntList(""))
	if res != 0 {
		t.Fatal("score should be 0, it is: ", res)
	}
}

func TestScoreElements(t *testing.T) {
	res := ScoreElements(StrToIntList("pPLvts"))
	if res != 157 {
		t.Fatal("score should be 157, it is: ", res)
	}


	res = ScoreElements(StrToIntList("pPpPpP"))
	if res != 174 {
		t.Fatal("score should be 171, it is: ", res)
	}
}