package main

import "fmt"

func dailyTemperatures(temperatures []int) []int {

	tempArrayLength := len(temperatures)
	tempWaitOutput := make([]int, tempArrayLength) // creating an array the same size as temperatures
	for index := range temperatures {
		if index < tempArrayLength {
			for j := index + 1; j < tempArrayLength; j++ {
				if temperatures[j] > temperatures[index] {
					tempWaitOutput[index] = j - index
					break
				}
			}
		}
	}
	return tempWaitOutput
}

func main() {

	fmt.Println(dailyTemperatures([]int{73, 74, 75, 71, 69, 72, 76, 73}))
	fmt.Println(dailyTemperatures([]int{30, 40, 50, 60}))
	fmt.Println(dailyTemperatures([]int{30, 60, 90}))

}
