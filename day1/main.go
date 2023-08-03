package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	fmt.Println("Hello world");

	f, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()
	var sumCalories = 0
	var maxCalories = 0

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		strCalories := scanner.Text()
		intCalories, err := strconv.Atoi(strCalories)
		if err != nil {
			if maxCalories < sumCalories {
				maxCalories = sumCalories
			}
			sumCalories = 0
		} else {
			sumCalories += intCalories
		}
	}
	if maxCalories < sumCalories {
		maxCalories = sumCalories
	}
	fmt.Println("MaxCalories: ", maxCalories)
}