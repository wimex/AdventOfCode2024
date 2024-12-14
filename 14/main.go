package main

import (
	"bufio"
	"fmt"
	"github.com/samber/lo"
	"os"
	"sort"
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
	//width, height := 11, 7
	time := 0
	question1 := 0
	question2 := 0
	for true { //} 100000 {
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

		quad1, quad2, quad3, quad4 := calculate(robots, width, height)
		if time == 99 {
			question1 = quad1 * quad2 * quad3 * quad4
		}

		if check(robots) {
			question2 = time + 1
			draw(robots, width, height)

			break
		}

		time++
	}

	fmt.Println("Question 1:", question1)
	fmt.Println("Question 2:", question2)
}

func calculate(robots []Robot, width, height int) (int, int, int, int) {
	xmiddle, ymiddle := (width-1)/2, (height-1)/2
	quad1, quad2, quad3, quad4 := 0, 0, 0, 0
	for _, robot := range robots {
		if robot.position.x < xmiddle && robot.position.y < ymiddle {
			quad1++
		}
		if robot.position.x > xmiddle && robot.position.y < ymiddle {
			quad2++
		}
		if robot.position.x < xmiddle && robot.position.y > ymiddle {
			quad3++
		}
		if robot.position.x > xmiddle && robot.position.y > ymiddle {
			quad4++
		}
	}

	return quad1, quad2, quad3, quad4
}

// Simply check for robots that form a line that is long enough (which will be the frame of the image).
// I originally assumed that the image will be symmetrical, but that didn't work out, so had to check frames by hand.
// After finding the image, I came up with this solution.
func check(robots []Robot) bool {
	length := 29

	grouped := lo.GroupBy(robots, func(robot Robot) int { return robot.position.y })
	for _, group := range grouped {
		if len(group) < length {
			continue
		}

		adjacent := 0
		sort.Slice(group, func(i, j int) bool { return group[i].position.x < group[j].position.x })
		for i := 0; i < len(group)-1; i++ {
			if group[i+1].position.x-group[i].position.x == 1 {
				adjacent++

				if adjacent >= length {
					return true
				}
			} else {
				adjacent = 1
			}
		}
	}

	return false
}

func draw(robots []Robot, width, height int) {
	for j := 0; j < height; j++ {
		for i := 0; i < width; i++ {
			_, found := lo.Find(robots, func(robot Robot) bool { return robot.position.x == i && robot.position.y == j })
			if found {
				fmt.Print("*")
			} else {
				fmt.Print(" ")
			}
		}
		fmt.Println()
	}
}
