package day4

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func parseSection(section string) (int, int, error) {
	sectionSplit := strings.Split(section, "-")
	if len(sectionSplit) != 2 {
		return 0, 0, fmt.Errorf("cannot parse section: %s", section)
	}
	start, err1 := strconv.Atoi(sectionSplit[0])
	end, err2 := strconv.Atoi(sectionSplit[1])
	if err1 != nil || err2 != nil {
		return 0, 0, fmt.Errorf("cannot parse section: %s", section)
	}

	if end < start {
		return 0, 0, fmt.Errorf("end cannot be < start: %s", section)
	}

	return start, end, nil
}

func isContained(start1 int, end1 int, start2 int, end2 int) (bool, error) {
	if end1 < start1 {
		return false, fmt.Errorf("end (%d) cannot be lower than start (%d) ", end1, start1)
	}

	if end2 < start2 {
		return false, fmt.Errorf("end (%d) cannot be lower than start (%d) ", end2, start2)
	}
	if start1 >= start2 {
		if end1 <= end2 {
			return true, nil
		}
	}

	if start2 >= start1 {
		if end2 <= end1 {
			return true, nil
		}
	}
	return false, nil
}

func intersectPartial(start1 int, end1 int, start2 int, end2 int) (bool, error) {
	if end1 < start1 {
		return false, fmt.Errorf("end (%d) cannot be lower than start (%d) ", end1, start1)
	}

	if end2 < start2 {
		return false, fmt.Errorf("end (%d) cannot be lower than start (%d) ", end2, start2)
	}

	if start1 >= start2 && start1 <= end2 || start2 >= start1 && start2 <= end1 {
		return true, nil
	} else {
		return false, nil
	}
}

func RunSolution(filename string) int {
	fmt.Println("Running solution for day4, part1")

	f, err := os.Open(filename)
	if err != nil {
		log.Fatalf("Cannot find file %s", filename)
	}

	defer f.Close()

	scanner := bufio.NewScanner(f)
	numIntersects := 0
	for scanner.Scan() {
		line := scanner.Text()

		sections := strings.Split(line, ",")
		if len(sections) != 2 {
			log.Fatalf("cannot split line %s into sections", line)
		}
		start1, end1, err := parseSection(sections[0])
		if err != nil {
			log.Fatalf("Cannot parse section %s", sections[0])
		}

		start2, end2, err := parseSection(sections[1])
		if err != nil {
			log.Fatalf("Cannot parse section %s", sections[1])
		}
		contained, err := isContained(start1, end1, start2, end2)
		if err != nil {
			log.Fatal("Incorrect set data for intersects")
		}

		if contained {
			numIntersects += 1
		}
	}
	return numIntersects
}

func RunSolution2(filename string) int {
	fmt.Println("Running solution for day4, part2")

	f, err := os.Open(filename)
	if err != nil {
		log.Fatalf("Cannot find file %s", filename)
	}

	defer f.Close()

	scanner := bufio.NewScanner(f)
	numIntersects := 0
	for scanner.Scan() {
		line := scanner.Text()

		sections := strings.Split(line, ",")
		if len(sections) != 2 {
			log.Fatalf("cannot split line %s into sections", line)
		}
		start1, end1, err := parseSection(sections[0])
		if err != nil {
			log.Fatalf("Cannot parse section %s", sections[0])
		}

		start2, end2, err := parseSection(sections[1])
		if err != nil {
			log.Fatalf("Cannot parse section %s", sections[1])
		}
		contained, err := intersectPartial(start1, end1, start2, end2)
		if err != nil {
			log.Fatal("Incorrect set data for intersects")
		}

		if contained {
			numIntersects += 1
		}
	}
	return numIntersects
}
