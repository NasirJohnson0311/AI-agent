package main
import "fmt"

// Create a vehicle stuct
type Vehicle struct {
	ID          string
	Battery     float64
	Speed       float64
	Temperature float64
	Location    string
	Status      string
}

func main() {

	var object1 Vehicle 	// Create object of vehicle type

	// Create arrays for each vehicle characteristic 
	var vehicleID = [5]string{"V-100", "V-101", "V-102", "V-103", "V-104"}
	var vehicleBattery = [5]float64{99.0, 58.2, 76.9, 28.23, 80.0}
	var vehicleSpeed = [5]float64{42, 76, 95, 26, 33}
	var vehicleTemp = [5]float64{100, 70, 80, 90, 59}
	var vehicleLocation = [5]string{"Austin", "Lexington", "Alabama", "New York", "Oklahoma"}
	var vehicleStatus = [5]string{"ACTIVE", "CHARGING", "OFFLINE", "ACTIVE", "ACTIVE"}
	
	// Assign characteristics to vehicle object
	for i := 0; i < 5; i++ {
		object1.ID = vehicleID[i]
		object1.Battery = vehicleBattery[i]
		object1.Speed = vehicleSpeed[i]
		object1.Temperature = vehicleTemp[i]
		object1.Location = vehicleLocation[i]
		object1.Status = vehicleStatus[i]

		// Print characteristics of vehicle object
		GetVehicleID(object1)
		fmt.Println("Battery Low: ",IsBatteryLow(object1))
		GetVehicleSpeed(object1)
		GetVehicleTemperature(object1)
		GetVehicleLocation(object1)
		GetVehicleStatus(object1)
		fmt.Println()

	}

}	// End of main


// Get functions 
func GetVehicleID(object1 Vehicle) {
	fmt.Println("Vehicle ID: ",object1.ID)
}

func IsBatteryLow(object1 Vehicle) bool {
	
	if (object1.Battery < 45){
		return true
	} else {
		return false
	}
}

func GetVehicleSpeed(object1 Vehicle) {
	fmt.Println("Speed: ",object1.Speed)
}

func GetVehicleTemperature(object1 Vehicle){
	fmt.Println("Temperature: ",object1.Temperature)
}

func GetVehicleLocation(object1 Vehicle){
	fmt.Println("Location: ",object1.Location)
}

func GetVehicleStatus(object1 Vehicle){
	fmt.Println("Status: ",object1.Status)
}


