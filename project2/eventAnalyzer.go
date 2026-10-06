package main

import (
	"bufio"
	"os"
	"strings"
)

/*
	This project is meant to gather information about vehicle data and analyze it
	Vehicle data is kept in eventData.txt
	Each event must be read from this file and stored within a map
		Map: [Vehicle -> Event number] + [Event -> Occurences]
	Tests will be conducted at the end of implementation to make sure everything works
*/

type Event struct {
	vehicle string
	event   string
	value   int
}

func main() {

	vehicleOccurences := make(map[string]int)

	file, fileError := os.Open("eventData.txt") // Open file

	if fileError != nil { // File error handling
		panic(fileError)
	}

	scanner := bufio.NewScanner(file) // Create new scanner

	for scanner.Scan() { // While scanner is reading file
		line := scanner.Text()                // Get text read from scanner
		lineSlice := strings.Split(line, ",") // Split text up by comma
		vehicleOccurences[lineSlice[0]]++
		scanner.Err()
	}

	// GetTotalEvents(vehicleOccurences) -> int
	// VehicleEvents(vehicleOccurences)
	// MostEvents(vehicleOccurences) -> string
	// GetOverheatingEvents -> int
}

func GetTotalEvents()       {} // This function will return the total number of events
func VehicleEvents()        {} // This function will list the number of events each vehicle has
func MostEvents()           {} // This function will return the vehicle with the most events
func GetOverheatingEvents() {} // This function will get the number of overheating events
func VehiclesLatestEvent()  {} // This function will get the latest event for each vehicle
