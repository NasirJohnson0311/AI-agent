package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

	/*
		This project is meant to gather information about vehicle data and analyze it
		Vehicle data is kept in eventData.txt
		Each event must be read from this file and stored within a map
			Map: [Vehicle -> Event number] + [Event -> Occurences]
		Tests will be conducted at the end of implementation to make sure everything works
	*/

	fmt.Println("Inside of event analyzer")

	// Open file using package, store store results in file if correctly opened, if not store in file error
	file, fileError := os.Open("eventData.txt")

	// Check to see if file error had returned something
	if fileError != nil {
		panic(fileError)
	} else {
		fmt.Println("No errors %v", file)
	}

	// Create a new scanner for our file using bufio
	scanner := bufio.NewScanner(file)

	// While scanner is still reading file
	for scanner.Scan() {
		line := scanner.Text() // Grab text that scanner has just received
		fmt.Println(line)
	}

}

func GetTotalEvents()       {} // This function will return the total number of events
func VehicleEvents()        {} // This function will list the number of events each vehicle has
func MostEvents()           {} // This function will return the vehicle with the most events
func GetOverheatingEvents() {} // This function will get the number of overheating events
func VehiclesLatestEvent()  {} // This function will get the latest event for each vehicle
