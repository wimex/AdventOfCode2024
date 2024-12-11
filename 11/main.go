package main

import (
	"bufio"
	"fmt"
	"github.com/samber/lo"
	"math/big"
	"os"
	"strconv"
	"strings"
)

type coord struct {
	x int
	y int
}

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

	history := [][]*big.Int{}
	step := stones
	for i := 0; i < 25; i++ {
		step = blink(step)
		history = append(history, step)
		fmt.Println("Step", i, ":", lo.Map(step, func(s *big.Int, _ int) string { return s.String() }))
	}

	fmt.Println("Question 1:", len(step))
}

func blink(stones []*big.Int) []*big.Int {
	results := []*big.Int{}

	for _, stone := range stones {
		result := apply(stone)
		results = append(results, result...)
	}

	return results
}

var ZERO = big.NewInt(0)
var ONE = big.NewInt(1)
var TWO = big.NewInt(2)
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

	/*res, _ := stone.DivMod(stone, TWO, new(big.Int))
	if res == ZERO {
		copy := new(big.Int).Set(res)
		return []*big.Int{copy}
	}*/

	multiply := clone.Mul(clone, TTF)
	return []*big.Int{multiply}
}
