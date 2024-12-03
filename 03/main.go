package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"sort"
	"strconv"
)

func main() {
	file, _ := os.Open("input.txt")
	defer file.Close()

	var regex_muls = regexp.MustCompile(`mul\((\d{1,3}),(\d{1,3})\)`)
	var regex_dos = regexp.MustCompile(`do\(\)`)
	var regex_donts = regexp.MustCompile(`don\'t\(\)`)
	var instructions = make([][]int, 0)

	scanner := bufio.NewScanner(file)
	concat := ""
	for scanner.Scan() {
		line := scanner.Text()
		concat += line
	}

	muls := regex_muls.FindAllStringSubmatch(concat, -1)
	muli := regex_muls.FindAllStringSubmatchIndex(concat, -1)

	dos := regex_dos.FindAllStringSubmatch(concat, -1)
	doi := regex_dos.FindAllStringSubmatchIndex(concat, -1)

	donts := regex_donts.FindAllStringSubmatch(concat, -1)
	donti := regex_donts.FindAllStringSubmatchIndex(concat, -1)

	for index, match := range muls {
		num1, _ := strconv.Atoi(match[1])
		num2, _ := strconv.Atoi(match[2])
		instructions = append(instructions, []int{muli[index][0], 0, num1, num2})
	}

	for index, _ := range dos {
		instructions = append(instructions, []int{doi[index][0], 1, 0, 0})
	}

	for index, _ := range donts {
		instructions = append(instructions, []int{donti[index][0], 2, 0, 0})
	}

	sort.Slice(instructions, func(i, j int) bool {
		return instructions[i][0] < instructions[j][0]
	})

	question1 := 0
	question2 := 0
	enabled := true
	for _, nums := range instructions {
		if nums[1] == 0 {
			question1 += nums[2] * nums[3]
		}

		if nums[1] == 0 && enabled {
			question2 += nums[2] * nums[3]
		}

		if nums[1] == 1 {
			enabled = true
		}

		if nums[1] == 2 {
			enabled = false
		}
	}

	fmt.Println("Question 1:", question1)
	fmt.Println("Question 2:", question2)
}
