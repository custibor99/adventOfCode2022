package main

import (
	"bufio"
	"errors"
	"fmt"
	"log"
	"os"
	"strings"
)

// A, X is ROCK
// B, Y is PAPER
// C, Z IS SCISORS

func scoreGame(oponent string, player string) (int, error) {
	var score int = 0
	
	switch player {
		case "X" : score += 1; 
		case "Y" : score += 2
		case "Z" : score += 3
		default : return 0, errors.New("Invalid player input")
	}
	
	if oponent != "A" && oponent != "B" && oponent != "C" {
		return 0, errors.New("Invalid oponent input")
	}

	if player == "X" {
		if oponent == "A" {
			score += 3
		} else if oponent == "C" {
			score += 6
		}
	} else if player == "Y" {
		if oponent == "A" {
			score += 6
		} else if oponent == "B" {
			score += 3
		} 
	} else if player == "Z" {
		if oponent == "B" {
			score += 6
		} else if oponent == "C" {
			score += 3
		} 
	}
	return score, nil
}

func main() {
	var filename string = os.Args[1]
	file, err := os.Open(filename)
	defer file.Close()

	if err != nil {
		log.Fatal(err)
	}
	
	scanner := bufio.NewScanner(file)
	totalScore := 0
	for scanner.Scan() {
		line := scanner.Text()
		elements := strings.Split(line, " ")

		if len(elements) != 2 {
			log.Fatal("Number of elements in row is off: ", len(elements), " ", elements)
		}

		gameScore, err := scoreGame(elements[0], elements[1])
		if err != nil {
			log.Fatal(err)
		}
		totalScore += gameScore 
	}
	fmt.Println(totalScore)
}