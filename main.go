package main

import (
	"fmt"
	"log"
	"net/url"
	"time"

	"github.com/voigt/howlong/helper"
)

type Me struct {
	Dob        string  `json:"dob"`
	Country    string  `json:"country"`
	Expectancy float64 `json:"total_life_expectancy"`
	Sex        string  `json:"sex"`
	Diabetes   string  `json:"diabetes"`
}

func main() {
	me := &Me{
		Sex:     "male",
		Country: "Germany",
		Dob:     "1989-11-29",
	}
	timeFormat := "2006-01-02 15:04 MST"
	birthday := me.Dob + " 10:30 UTC"

	me.Expectancy = expectedAge("1989-11-29", me.Sex, me.Country)

	then, err := time.Parse(timeFormat, birthday)
	if err != nil {
		fmt.Println(err)
		return
	}

	duration := time.Since(then).Hours()
	weeks := getWeekCount(duration)

	printLife(weeks)
	fmt.Printf("Your Birthday: %s\n", me.Dob)
	fmt.Printf("%d weeks have passed since then!\n", weeks)
	fmt.Printf("You will probably live until you are %f years old.\n", me.Expectancy)

}

// Requests total life expectancy from api.population.io
// Expects date of birth and country
// Sample Request:
// curl 'http://api.population.io:80/1.0/life-expectancy/total/male/Germany/1989-11-29/'
// {
//   "dob": "1989-11-29",
//   "country": "Germany",
//   "total_life_expectancy": 84.9352193923596,
//   "sex": "male"
// }
func expectedAge(dob string, gender string, country string) float64 {

	u, err := url.Parse("http://api.population.io:80/1.0/life-expectancy/total")
	if err != nil {
		log.Fatal(err)
	}
	u.Path = u.Path + "/" + gender + "/" + country + "/" + dob
	me := &Me{}
	helper.GetJson(u.String(), me)

	return me.Expectancy
}

func getWeekCount(duration float64) int {

	weeks := int(helper.Round(duration / 24 / 7))

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

	for i <= int(helper.Round(float64(weeks)/52)) {

		fmt.Printf("%.2d ", i)
		yearInWeeks(weeksLeft)
		weeksLeft = weeksLeft - 52
		i++
	}
}
