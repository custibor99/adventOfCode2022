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
}