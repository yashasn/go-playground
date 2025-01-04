package main

import "fmt"

type cost struct {
	day   int
	value float64
}

func getCostsByDay(costs []cost) []float64 {
	noOfDays := costs[len(costs)-1].day + 1
	costsByDay := make([]float64, noOfDays)
	for i := 0; i < len(costs); i++ {
		cost := costs[i]

		costsByDay[cost.day] += cost.value
	}
	return costsByDay
}

func getCostsByDayAlt(costs []cost) []float64 {
	costsByDay := []float64{}
	for i := 0; i < len(costs); i++ {
		cost := costs[i]
		for cost.day >= len(costsByDay) {
			costsByDay = append(costsByDay, 0.0)
		}
		costsByDay[cost.day] += cost.value
	}
	return costsByDay
}

func TestCostsByDay() {
	costs := []cost{
		{0, 4.0},
		{1, 2.1},
		{1, 3.1},
		{5, 2.5},
	}
	costsByDay := getCostsByDayAlt(costs)

	for i := 0; i < len(costsByDay); i++ {
		fmt.Printf("%f \n", costsByDay[i])
	}
}
