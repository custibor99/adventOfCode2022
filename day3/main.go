package main

import (
	"bufio"
	"fmt"
	"log"
	"os"

	utils "example.com/utils/day3utils"
)


func main() {
	utils.Hello()
	
	f, err := os.Open(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()
	scanner := bufio.NewScanner(f)
	/*
	scanner := bufio.NewScanner(f)
	score := 0
	for scanner.Scan() {
		line := scanner.Text()
		half := len(line) / 2
		list := utils.StrToIntList(line)
		s1 := list[:half]
		s2 := list[half:]
		duplicates := utils.FindDuplicateElements(s1, s2)
		score += utils.ScoreElements(duplicates)		
	}
	fmt.Println(score)
	*/

	var backpacks [3][]int
	i := 0
	score := 0
	for scanner.Scan() {
		input := scanner.Text()
		backpacks[i % 3] = utils.StrToIntList(input)
		if i % 3 == 2 {
			duplicates := utils.FindDuplicateElements(backpacks[0], backpacks[1])
			duplicates = utils.FindDuplicateElements(duplicates, backpacks[2])
			score += utils.ScoreElements(duplicates)
		}
		i++
	}
	fmt.Println(score)
}