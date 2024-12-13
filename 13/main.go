package main

import (
	"bufio"
	"cmp"
	"fmt"
	"github.com/samber/lo"
	"os"
)

type Coordinate struct {
	x, y int
}

type Machine struct {
	button1, button2, prize Coordinate
}

type Result struct {
	reached              bool
	press1, press2, cost int
}

func main() {
	file, _ := os.Open("input.txt")
	defer file.Close()

	scanner := bufio.NewScanner(file)
	machines := make([]Machine, 0)
	for scanner.Scan() {
		line1 := scanner.Text()
		scanner.Scan()
		line2 := scanner.Text()
		scanner.Scan()
		line3 := scanner.Text()
		scanner.Scan()

		ax, ay := 0, 0
		bx, by := 0, 0
		tx, ty := 0, 0
		fmt.Sscanf(line1, "Button A: X+%d, Y+%d", &ax, &ay)
		fmt.Sscanf(line2, "Button B: X+%d, Y+%d", &bx, &by)
		fmt.Sscanf(line3, "Prize: X=%d, Y=%d", &tx, &ty)

		machine := Machine{Coordinate{ax, ay}, Coordinate{bx, by}, Coordinate{tx, ty}}
		machines = append(machines, machine)

		scanner.Text() //Throw away empty line
	}

	question1 := 0
	for _, machine := range machines {
		//fmt.Println(machine)
		question1 += simulate(machine)
	}

	fmt.Println("Question 1:", question1)
}

func simulate(machine Machine) int {
	channel := make(chan []Result)
	go step(machine, channel)

	result := <-channel
	minimum := lo.MinBy(result, func(a Result, b Result) bool { return cmp.Compare(a.cost, b.cost) != 0 })
	fmt.Println("LEN:", len(result), "MIN:", minimum)

	return minimum.cost
}

func step(machine Machine, channel chan []Result) {
	results := make([]Result, 0)

	for i := 0; i < machine.prize.x; i++ {
		for j := 0; j < machine.prize.y; j++ {
			v1 := i*machine.button1.x + j*machine.button2.x
			v2 := i*machine.button1.y + j*machine.button2.y
			if v1 == machine.prize.x && v2 == machine.prize.y {
				results = append(results, Result{true, i, j, i*3 + j*1})
			}
		}
	}

	channel <- results
}

func check(machine Machine, i, j int, channel chan Result) {
}
