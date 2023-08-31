package day4

import (
	"fmt"
	"testing"
)

func TestParseSection(t *testing.T) {
	section := "6-10"
	start, end, err := parseSection(section)
	if err != nil || start != 6 || end != 10 {
		t.Fatal("Expected start: 6, end: 10. Got: ", start, end)
	}

	section = "60143-1023424"
	start, end, err = parseSection(section)
	if err != nil && start != 60143 || end != 1023424 {
		t.Fatal("Expected start: 60143, end: 1023424. Got: ", start, end)
	}

	section = "10-5"
	start, end, err = parseSection(section)
	if err == nil {
		t.Fatal("Expected error for section", section)
	}

	section = ""
	start, end, err = parseSection(section)
	if err == nil {
		t.Fatal("Expected error for section", section)
	}
}

func TestIsContained(t *testing.T) {
	start1 := 6
	end1 := 8
	start2 := 7
	end2 := 7
	contained, err := isContained(start1, end1, start2, end2)
	if err != nil || !contained {
		fmt.Println(err, contained)
		t.Fatalf("%d-%d and %d-%d should have a valid ful intersection", start1, end1, start2, end2)
	}

	start1 = 6
	end1 = 8
	start2 = 6
	end2 = 7
	contained, err = isContained(start1, end1, start2, end2)
	if err != nil || !contained {
		fmt.Println(err, contained)
		t.Fatalf("%d-%d and %d-%d should have a valid ful intersection", start1, end1, start2, end2)
	}

	start1 = 6
	end1 = 7
	start2 = 5
	end2 = 8
	contained, err = isContained(start1, end1, start2, end2)
	if err != nil || !contained {
		t.Fatalf("%d-%d and %d-%d should have a valid ful intersection", start1, end1, start2, end2)
	}

	start1 = 1
	end1 = 1
	start2 = 1
	end2 = 1
	contained, err = isContained(start1, end1, start2, end2)
	if err != nil || !contained {
		t.Fatalf("%d-%d and %d-%d should have a valid ful intersection", start1, end1, start2, end2)
	}

	start1 = 6
	end1 = 7
	start2 = 5
	end2 = 8
	contained, err = isContained(start1, end1, start2, end2)
	if err != nil || !contained {
		t.Fatalf("%d-%d and %d-%d should have a valid ful intersection", start1, end1, start2, end2)
	}

	start1 = 1
	end1 = 3
	start2 = 5
	end2 = 8
	contained, err = isContained(start1, end1, start2, end2)
	if err != nil || contained {
		t.Fatalf("%d-%d and %d-%d should not have a valid ful intersection", start1, end1, start2, end2)
	}

	start1 = 1
	end1 = 3
	start2 = 2
	end2 = 8
	contained, err = isContained(start1, end1, start2, end2)
	if err != nil || contained {
		t.Fatalf("%d-%d and %d-%d should not have a valid ful intersection", start1, end1, start2, end2)
	}

	start1 = 1
	end1 = 3
	start2 = 0
	end2 = 2
	contained, err = isContained(start1, end1, start2, end2)
	if err != nil || contained {
		t.Fatalf("%d-%d and %d-%d should not have a valid ful intersection", start1, end1, start2, end2)
	}
}

func TestCountIntersectionLength(t *testing.T) {
	start1 := 6
	end1 := 8
	start2 := 7
	end2 := 7
	doesIntersect, err := intersectPartial(start1, end1, start2, end2)
	if err != nil || !doesIntersect {
		t.Fatalf("Excpected true, got %t", doesIntersect)
	}

	start1 = 6
	end1 = 8
	start2 = 6
	end2 = 7
	doesIntersect, err = intersectPartial(start1, end1, start2, end2)
	if err != nil || !doesIntersect {
		t.Fatalf("Excpected true, got %t", doesIntersect)
	}

	start1 = 6
	end1 = 7
	start2 = 5
	end2 = 8
	doesIntersect, err = intersectPartial(start1, end1, start2, end2)
	if err != nil || !doesIntersect {
		t.Fatalf("Excpected true, got %t", doesIntersect)
	}

	start1 = 1
	end1 = 1
	start2 = 1
	end2 = 1
	doesIntersect, err = intersectPartial(start1, end1, start2, end2)
	if err != nil || !doesIntersect {
		t.Fatalf("Excpected true, got %t", doesIntersect)
	}

	start1 = 1
	end1 = 3
	start2 = 5
	end2 = 8
	doesIntersect, err = intersectPartial(start1, end1, start2, end2)
	if err != nil || doesIntersect {
		t.Fatalf("Excpected true, got %t", doesIntersect)
	}

	start1 = 2
	end1 = 8
	start2 = 1
	end2 = 3
	doesIntersect, err = intersectPartial(start1, end1, start2, end2)
	if err != nil || !doesIntersect {
		t.Fatalf("Excpected true, got %t", doesIntersect)
	}

	start1 = 1
	end1 = 3
	start2 = 2
	end2 = 4
	doesIntersect, err = intersectPartial(start1, end1, start2, end2)
	if err != nil || !doesIntersect {
		t.Fatalf("Excpected true, got %t", doesIntersect)
	}
}

func TestRunSolution2(t *testing.T) {
	res := RunSolution2("../inputFiles/day4/input1.txt")
	fmt.Println(res)
}
