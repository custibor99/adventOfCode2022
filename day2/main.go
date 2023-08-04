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

const(
	rock 		int = 1
	paper 		int = 2
	scissors 	int = 3
)

var shapeMap = map[string]int {
	"X": rock,
	"A": rock,
	"Y": paper,
	"B": paper,
	"Z": scissors,
	"C": scissors,
}

func scoreRound(playerShape int, oponentShape int) int {
	if playerShape == oponentShape {
		return 3 + playerShape
	}
	if playerShape - oponentShape == -1 || playerShape - oponentShape == 2 {
		return playerShape
	}
	return 6 + playerShape
}


func parseInput(oponent string, player string) (int, int, error) {	
	playerShape, exists := shapeMap[player]
	if (exists == false) {
		return 0, 0, errors.New("Invalid player input")
	}

	oponentShape, exists := shapeMap[oponent]
	if (exists == false) {
		return 0, 0,  errors.New("Invalid oponent input")
	}
	return playerShape, oponentShape, nil
}

func parseInput2(oponent string, player string) (int, int, error) {	
	oponentShape, exists := shapeMap[oponent]
	if (exists == false) {
		return 0, 0,  errors.New("Invalid oponent input")
	}

	_, exists = shapeMap[player]
	if (exists == false) {
		return 0, 0, errors.New("Invalid player input")
	}

	if player == "Y" {
		return oponentShape, oponentShape, nil
	} else if player == "Z" {
		playerShape := (oponentShape + 1) % 3
		if playerShape == 0 {
			playerShape = 3
		}
		return playerShape, oponentShape, nil
	} else {
		playerShape := oponentShape - 1
		if playerShape == 0{
			playerShape = 3
		}
		return playerShape, oponentShape, nil
	}
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
	totalScore2 := 0
	for scanner.Scan() {
		line := scanner.Text()
		elements := strings.Split(line, " ")

		if len(elements) != 2 {
			log.Fatal("Number of elements in row is off: ", len(elements), " ", elements)
		}

		playerShape, oponentShape, err := parseInput(elements[0], elements[1])
		if err != nil {
			log.Fatal(err)
		}

		playerShape2, oponentShape2, err := parseInput2(elements[0], elements[1])
		if err != nil {
			log.Fatal(err)
		}
		roundScore := scoreRound(playerShape, oponentShape)
		roundScore2 := scoreRound(playerShape2, oponentShape2)
		totalScore += roundScore 
		totalScore2 += roundScore2
	}
	fmt.Println(totalScore)
	fmt.Println(totalScore2)
}