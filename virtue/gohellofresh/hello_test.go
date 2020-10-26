package main

import (
	"encoding/json"
	"fmt"
	"os"
	"testing"
)

func TestPostcodeInTimeRange(t *testing.T) {

	fd := FoodDelivery{PostCode: "10161", Recipe: "Carbonara real no cream", Delivery: "Saturday 10AM - 6PM"}

	perfData := PostcodePerformanceTimerange{Postcode: "10161", From: "10AM", To: "5PM", Count: 0}
	if postcodeInTimeRange(&fd, perfData.Postcode, perfData.From, perfData.To) {
		fmt.Printf("%v is in range %v\n", perfData, fd)
	} else {
		fmt.Printf("%v is NOT in range %v\n", perfData, fd)
	}

	perfData = PostcodePerformanceTimerange{Postcode: "10161", From: "11AM", To: "5PM", Count: 0}
	if postcodeInTimeRange(&fd, perfData.Postcode, perfData.From, perfData.To) {
		fmt.Printf("%v is in range %v\n", perfData, fd)
	} else {
		fmt.Printf("%v is NOT in range %v\n", perfData, fd)
	}

	perfData = PostcodePerformanceTimerange{Postcode: "10161", From: "10AM", To: "7PM", Count: 0}
	if postcodeInTimeRange(&fd, perfData.Postcode, perfData.From, perfData.To) {
		fmt.Printf("%v is in range %v\n", perfData, fd)
	} else {
		fmt.Printf("%v is NOT in range %v\n", perfData, fd)
	}

	perfData = PostcodePerformanceTimerange{Postcode: "10161", From: "9AM", To: "5PM", Count: 0}
	if postcodeInTimeRange(&fd, perfData.Postcode, perfData.From, perfData.To) {
		fmt.Printf("%v is in range %v\n", perfData, fd)
	} else {
		fmt.Printf("%v is NOT in range %v\n", perfData, fd)
	}
}

func TestOutput(t *testing.T) {

	myOutput := Output{
		UniqueRecipeCount: 15,
		CountPerRecipe: []RecipeCount{
			{
				Recipe: "Mediterranean Baked Veggies",
				Count:  1,
			},
			{
				Recipe: "Speedy Steak Fajitas",
				Count:  1,
			},
			{
				Recipe: "Tex-Mex Tilapia",
				Count:  3,
			},
		},
		BusiestPostcode: PostcodeCount{
			Postcode: "10120",
			Count:    1000,
		},
		CountPerPostcodeAndTime: PostcodePerformanceTimerange{
			Postcode: "10120",
			From:     "11AM",
			To:       "3PM",
			Count:    500,
		},
		MatchByName: []string{
			"Mediterranean Baked Veggies", "Speedy Steak Fajitas", "Tex-Mex Tilapia",
		},
	}

	encoder := json.NewEncoder(os.Stdout) //(f)
	encoder.SetIndent("", "  ")           // leave empty prefix for first arg, second arg is the indent spaces
	if err := encoder.Encode(&myOutput); err != nil {
		t.Errorf(err.Error())
	}
}
