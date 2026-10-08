package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"
)

/*
	This project is meant to gather information about vehicle data and analyze it
	Vehicle data is kept in eventData.txt
	Each event must be read from this file and stored within a map
		Map: [Vehicle -> Event number] + [Event -> Occurences]
	Tests will be conducted at the end of implementation to make sure everything works
*/

type Event struct {
	VehicleID string
	Type      string
	Value     float64
	Timestamp time.Time
}

func main() {

	events := []Event{}

	file, fileError := os.Open("eventData.txt") // Open file
	if fileError != nil {                       // File error handling
		panic(fileError)
	}

	scanner := bufio.NewScanner(file) // Create new scanner
	for scanner.Scan() {              // While scanner is reading file
		currLine := scanner.Text()                // Get text read from scanner
		lineSlice := strings.Split(currLine, ",") // Split text up by comma

		newEvent := Event{lineSlice[0], lineSlice[1], float64(currLine[3]), time.Now()}
		events = append(events, newEvent)

		scanner.Err()
	}

	eventMap := make(map[string]int)
	for _, event := range events {
		eventMap[event.VehicleID]++
	}

	GetTotalEvents(events)
	GetVehicleEvents(eventMap)
	GetMostEvents(eventMap)
	GetOverheatingEvents(events)
	GetVehiclesLatestEvent(events)

}

func GetTotalEvents(events []Event) { // This function will return the total number of events

	numEvents := 0
	for numEvents < len(events) {
		numEvents++
	}

	fmt.Printf("Total events: %v \n", numEvents)
	fmt.Println()
}

func GetVehicleEvents(eventMap map[string]int) { // This function will list the number of events each vehicle has

	for event := range eventMap {
		fmt.Printf("%v has %v events", event, eventMap[event])
		fmt.Println()
	}

	fmt.Println()

}

func GetMostEvents(eventMap map[string]int) { // This function will return the vehicle with the most events

	currMost := ""
	numEvents := 0

	for currEvent := range eventMap {
		if eventMap[currEvent] > numEvents {
			currMost = currEvent
			numEvents = eventMap[currEvent]
		}
	}

	fmt.Printf("%v has the most events (%v)\n", currMost, numEvents)
	fmt.Println()

}

func GetOverheatingEvents(events []Event) { // This function will get the number of overheating events

	numOverheatingEvents := 0

	for _, event := range events {
		if event.Type == "OVERHEATING" {
			numOverheatingEvents++
		}
	}

	fmt.Printf("Overheating events: %v \n", numOverheatingEvents)
	fmt.Println()

}

func GetVehiclesLatestEvent(events []Event) { // This function will get the latest event for each vehicle

	for _, event := range events {
		fmt.Println(event.Timestamp)
	}
}
