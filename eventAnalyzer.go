package main

import "fmt"

func EventAnalyzer() {

	/*
		This project is meant to gather information about vehicle data and analyze it
		Vehicle data is kept in eventData.txt
		Each event must be read from this file and stored within a map
			Map: [Vehicle -> Event number] + [Event -> Occurences]
		Tests will be conducted at the end of implementation to make sure everything works
	*/

	fmt.Println("Hello world")
}

func GetTotalEvents()       {} // This function will return the total number of events
func VehicleEvents()        {} // This function will list the number of events each vehicle has
func MostEvents()           {} // This function will return the vehicle with the most events
func GetOverheatingEvents() {} // This function will get the number of overheating events
func VehiclesLatestEvent()  {} // This function will get the latest event for each vehicle
