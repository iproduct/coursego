package main

import (
	"fmt"
)

func calculateArea(radius int) (int, error) {
	if radius < 0 {
		//return 0, errors.New("provide positive radius: " + strconv.Itoa(radius))
		return 0, fmt.Errorf("provide positive radius: %d", radius)
	}
	return radius * radius, nil
}

func main() {
	areaValue, err := calculateArea(-12)
	if err != nil {
		fmt.Printf("Error: %v", err) //Error: provide positive radius: -12
		return
	}

	fmt.Println(areaValue)
}
