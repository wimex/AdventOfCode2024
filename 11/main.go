package main

import (
	"bufio"
	"fmt"
	"math/big"
	"os"
	"strconv"
	"strings"
)

var CACHE = map[string]int{}

func main() {
	file, _ := os.Open("input.txt")
	defer file.Close()

	scanner := bufio.NewScanner(file)
	stones := []*big.Int{}
	for scanner.Scan() {
		line := scanner.Text()
		pieces := strings.Split(line, " ")
		for _, piece := range pieces {
			parsed, _ := strconv.ParseInt(piece, 10, 64)
			number := big.NewInt(parsed)
			stones = append(stones, number)
		}
	}

	length1 := 0
	length2 := 0
	for _, stone := range stones {
		length1 += blink(stone, 25)
		length2 += blink(stone, 75)
	}

	fmt.Println("Question 1:", length1)
	fmt.Println("Question 2:", length2)
}

func blink(stone *big.Int, steps int) int {
	cache := "L" + stone.String() + "S" + strconv.Itoa(steps)
	if cached, ok := CACHE[cache]; ok {
		return cached
	}

	next := apply(stone)
	result := 0
	for _, expanded := range next {
		if steps == 1 {
			result++
		} else {
			result += blink(expanded, steps-1)
		}
	}

	CACHE[cache] = result

	return result
}

var ZERO = big.NewInt(0)
var ONE = big.NewInt(1)
var TTF = big.NewInt(2024)

func apply(stone *big.Int) []*big.Int {
	if stone.Cmp(ZERO) == 0 {
		clone := new(big.Int).Set(ONE)
		return []*big.Int{clone}
	}

	clone := new(big.Int).Set(stone)
	digits := clone.String()
	if len(digits)%2 == 0 {
		left := digits[:len(digits)/2]
		right := digits[len(digits)/2:]

		num1, _ := big.NewInt(0).SetString(left, 10)
		num2, _ := big.NewInt(0).SetString(right, 10)

		return []*big.Int{num1, num2}
	}

	multiply := clone.Mul(clone, TTF)
	return []*big.Int{multiply}
}
