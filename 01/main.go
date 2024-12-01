package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	file, _ := os.Open("input.txt")
	defer file.Close()

	array1 := []int{}
	array2 := []int{}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		pieces := strings.Split(line, "   ")

		num1, _ := strconv.Atoi(pieces[0])
		num2, _ := strconv.Atoi(pieces[1])

		array1 = append(array1, num1)
		array2 = append(array2, num2)
	}

	sort.Ints(array1)
	sort.Ints(array2)

	total := 0.0
	for i := 0; i < len(array1); i++ {
		total += math.Abs(float64(array1[i]) - float64(array2[i]))
	}

	fmt.Println("Question 1:", int(total))
}
