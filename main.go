package main

import "fmt"

func one() {
	fmt.Println("test")
}

func main() {
	one()
}

// Ask user city information
// Check if the city is in csv file // func checkCityExist(city)  isCityExist
// If user gives wrong input, ask again

// Ask user state information

// Check if state is in csv file.  // func checkStateExist(state) returns true or false
// If it is not ask again.

// If it is check if the state belongs to the given city // func checkStateBelongsToCity(city, state) // returns city or false
// If it is not say it and print the available states in that city, and ask state again.  // func getCityStates(city)  // returns all the states in given city.

// Return the latitude and longitude information for that specific city and state. // func returnLocation(city, state) // returns longitude and latitude values by looking to the csv file.
// Send request using openweather api and assign the response to a variable. // func fetchWeatherInformation // returns weatherInformation
// Print variable in appropriate format. fmtPrintln(weatherInformation)

//
