package day2

import (
	"fmt"
	"strconv"
	"strings"
)

// CONSTANTS
const redCubes int = 12
const greenCubes int = 13
const blueCubes int = 14

// SOLUTION

func solveGame(gameSets []string) (bool, int) {
	isSolved := true
	maxRed := 0
	maxGreen := 0
	maxBlue := 0

	for _, gameSet := range gameSets {
		cubes := strings.Split(gameSet, ", ")

		for _, cube := range cubes {
			splittedCubeString := strings.Split(strings.TrimSpace(cube), " ")

			cubeCount, err := strconv.Atoi(splittedCubeString[0])
			if err != nil {
				println("Cube Count : " + string(rune(cubeCount)) + " | Game set : " + gameSet)
				continue
			}

			cubeName := splittedCubeString[1]

			if cubeName == "red" {
				if maxRed < cubeCount {
					maxRed = cubeCount
				}

				if cubeCount > redCubes {
					isSolved = false
				}
			}

			if cubeName == "blue" {
				if maxBlue < cubeCount {
					maxBlue = cubeCount
				}

				if cubeCount > blueCubes {
					isSolved = false
				}
			}

			if cubeName == "green" {
				if maxGreen < cubeCount {
					maxGreen = cubeCount
				}

				if cubeCount > greenCubes {
					isSolved = false
				}
			}
		}
	}

	power := maxRed * maxGreen * maxBlue

	fmt.Printf("%v * %v * %v = %v", maxRed, maxGreen, maxBlue, power)

	return isSolved, power
}

func solveGames(input string) (int, int) {
	games := strings.Split(strings.ReplaceAll(strings.TrimSpace(input), "\r", ""), "\n")
	totalSolvableGames := 0
	totalPower := 0

	for _, val := range games {
		splitValue := strings.Split(val, ": ")
		gameNumber, err := strconv.Atoi(strings.Split(splitValue[0], " ")[1])
		if err != nil {
			println("Game Number : " + string(rune(gameNumber)))
			continue
		}
		gameSets := strings.Split(splitValue[1], ";")
		fmt.Printf("Game number : %v\n", gameNumber)

		isSolved, power := solveGame(gameSets)

		if isSolved {
			totalSolvableGames += gameNumber
		}

		totalPower += power
		fmt.Printf("\nPower: %v | Total Power: %v \n", power, totalPower)
	}

	return totalSolvableGames, totalPower
}

func Solve(input string) int {
	total, _ := solveGames(input)
	return total
}

func SolvePart2(input string) int {
	_, power := solveGames(input)
	return power
}
