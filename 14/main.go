package main

import (
	"bufio"
	"fmt"
	"os"
)

type Coordinate struct {
	x, y int
}

type Robot struct {
	position Coordinate
	velocity Coordinate
}

func main() {
	file, _ := os.Open("input.txt")
	defer file.Close()

	scanner := bufio.NewScanner(file)
	robots := make([]Robot, 0)
	for scanner.Scan() {
		robot := scanner.Text()

		var px, py, vx, vy int
		fmt.Sscanf(robot, "p=%d,%d v=%d,%d", &px, &py, &vx, &vy)

		robots = append(robots, Robot{Coordinate{px, py}, Coordinate{vx, vy}})
	}

	width, height := 101, 103
	time := 100
	for time > 0 {
		for i := 0; i < len(robots); i++ {
			robot := &robots[i]

			px := robot.position.x + robot.velocity.x
			py := robot.position.y + robot.velocity.y

			if px < 0 {
				robot.position.x = width + px
			} else if px >= width {
				robot.position.x = px - width
			} else {
				robot.position.x = px
			}

			if py < 0 {
				robot.position.y = height + py
			} else if py >= height {
				robot.position.y = py - height
			} else {
				robot.position.y = py
			}
		}

		time--
	}

	xmiddle, ymiddle := (width-1)/2, (height-1)/2
	question1 := 0
	for _, robot := range robots {
		if robot.position.x != xmiddle || robot.position.y != ymiddle {
			question1++
		}
	}

	fmt.Println("Question 1:", question1)
}
