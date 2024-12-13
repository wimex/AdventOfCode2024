package main

import (
	"bufio"
	"fmt"
	"math/big"
	"os"
)

type Coordinate struct {
	x, y *big.Int
}

type Machine struct {
	button1, button2, prize Coordinate
}

type Result struct {
	reached              bool
	press1, press2, cost *big.Int
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

		machine := Machine{
			Coordinate{
				new(big.Int).SetInt64(int64(ax)),
				new(big.Int).SetInt64(int64(ay))},
			Coordinate{
				new(big.Int).SetInt64(int64(bx)),
				new(big.Int).SetInt64(int64(by))},
			Coordinate{
				new(big.Int).SetInt64(int64(tx)),
				new(big.Int).SetInt64(int64(ty))},
		}

		machines = append(machines, machine)
		scanner.Text() //Throw away empty line
	}

	question1 := new(big.Int).SetInt64(0)
	for _, machine := range machines {
		//fmt.Println(machine)
		result := simulate(machine)
		question1.Add(question1, result)
	}

	fmt.Println("Question 1:", question1)
}

func simulate(machine Machine) *big.Int {
	result := step(machine)
	//minimum := lo.MinBy(result, func(a Result, b Result) bool { return cmp.Compare(a.cost, b.cost) != 0 })
	fmt.Println("RESULT:", result.cost.String())

	return result.cost
}

func step(machine Machine) Result {
	dig1 := new(big.Int).Sub(
		new(big.Int).Mul(new(big.Int).Set(machine.button1.y), machine.prize.x),
		new(big.Int).Mul(new(big.Int).Set(machine.button1.x), machine.prize.y))

	dig2 := new(big.Int).Sub(
		new(big.Int).Mul(new(big.Int).Set(machine.button1.y), machine.button2.x),
		new(big.Int).Mul(new(big.Int).Set(machine.button1.x), machine.button2.y))

	test1 := new(big.Int).Mod(new(big.Int).Set(dig1), dig2).Int64()
	if test1 != 0 {
		zero := new(big.Int).SetInt64(0)
		return Result{false, zero, zero, zero}
	}

	bpr := new(big.Int).Div(new(big.Int).Set(dig1), dig2)
	apr := new(big.Int).Div(
		new(big.Int).Sub(new(big.Int).Set(machine.prize.x), new(big.Int).Mul(new(big.Int).Set(machine.button2.x), bpr)),
		machine.button1.x)

	c1 := new(big.Int).Mul(new(big.Int).Set(apr), new(big.Int).SetInt64(3))
	c2 := new(big.Int).Mul(new(big.Int).Set(bpr), new(big.Int).SetInt64(1))
	sm := new(big.Int).Add(c1, c2)

	return Result{true, apr, bpr, sm}

	//axval = machine.button1.x, movement of button 1 in x
	//ayval = machine.button1.y, movement of button 1 in y
	//bxval = machine.button2.x, movement of button 2 in x
	//byval = machine.button2.y, movement of button 2 in y
	//apr = number of times button 1 is pressed
	//bpr = number of times button 2 is pressed
	//pricex = machine.prize.x, x coordinate of prize
	//pricey = machine.prize.y, y coordinate of prize

	//axval * apr + bxval * bpr = pricex
	//ayval * apr + byval * bpr = pricey
	//apr = (pricex - bxval * bpr) / axval

	//ayval * ((pricex - bxval * bpr) / axval) + byval * bpr = pricey
	//ayval * (pricex - bxval * bpr) + axval * byval * bpr = axval * pricey
	//ayval * pricex - ayval * bxval * bpr + axval * byval * bpr = axval * pricey
	//ayval * pricex - (ayval * bxval + axval * byval) * bpr = axval * pricey
	//-(ayval * bxval + axval * byval) * bpr = ayval * pricex - axval * pricey

	//machine.button1.x*a * (machine.prize.y - machine.button2.y*b) / machine.button1.y + machine.button2.x*b = machine.prize.x
	//machine.button1.x*a * (machine.prize.y - machine.button2.y*b) + machine.button2.x*b*machine.button1.y = machine.prize.x*machine.button1.y
	//machine.button1.x*a * machine.prize.y - machine.button1.x*a * machine.button2.y*b + machine.button2.x*b*machine.button1.y = machine.prize.x*machine.button1.y
	//machine.button1.x*a * machine.prize.y - (machine.button1.x*a
}

func check(machine Machine, i, j int, channel chan Result) {
}
