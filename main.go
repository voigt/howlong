package main

import (
	"fmt"
	"math"
	"time"
)

func main() {
	timeFormat := "2006-01-02 15:04 MST"
	birthday := "1989-11-29 10:30 UTC"

	then, err := time.Parse(timeFormat, birthday)
	if err != nil {
		fmt.Println(err)
		return
	}

	duration := time.Since(then).Hours()
	weeks := getWeekCount(duration)

	printLife(weeks)
	fmt.Printf("Your Birthday: %s\n", birthday)
	fmt.Printf("%d weeks have passed since then!\n", weeks)
}

func Round(f float64) float64 {
	return math.Floor(f + .5)
}

// Requests total life expectancy from api.population.io
// Expects date of birth and country
// {
//   "dob": "1989-11-29",
//   "country": "Germany",
//   "total_life_expectancy": 84.9352193923596,
//   "sex": "male"
// }
func expectedAge(birthday time, country string) float64 {
	// http://api.population.io:80/1.0/life-expectancy/total/male/Germany/1989-11-29/

	return expectedAge
}

func getWeekCount(duration float64) int {

	weeks := int(Round(duration / 24 / 7))

	return weeks
}

func yearInWeeks(livedWeeks int) {

	i := 0

	for i < 52 {
		if i <= livedWeeks {

			if i == 51 {
				fmt.Print("o\n")
			} else {
				fmt.Print("o")
			}

		} else if i > livedWeeks {

			if i == 51 {
				fmt.Print(".\n")
			} else {
				fmt.Print(".")
			}
		}
		i++
	}

}

func printLife(weeks int) {
	i := 0

	weeksLeft := weeks

	for i <= int(Round(float64(weeks)/52)) {

		fmt.Printf("%.2d ", i)
		yearInWeeks(weeksLeft)
		weeksLeft = weeksLeft - 52
		i++
	}
}
