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

	var array1 []int
	var array2 []int

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

	question1 := 0.0
	for i := 0; i < len(array1); i++ {
		question1 += math.Abs(float64(array1[i]) - float64(array2[i]))
	}

	dict := make(map[int]int)
	for _, num := range array2 {
		dict[num] = dict[num] + 1
	}

	question2 := 0
	for _, num := range array1 {
		question2 += num * dict[num]
	}

	fmt.Println("Question 1:", int(question1))
	fmt.Println("Question 2:", question2)
}
