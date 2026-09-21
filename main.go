package main
import "fmt"

const MAX_TEMP = 100
const LOW_BATTERY_THRESHOLD = 45

// Vehicle struct tempplate
type Vehicle struct {
	ID          string
	Battery     float64
	Speed       float64
	Temperature float64
	Location    string
	Status      string
}

func main() {
	// Instantiate vehicle objects 
	var vehicle1 = Vehicle{"V-100", 99.0, 42, 100, "Austin", "ACTIVE"}
	var vehicle2 = Vehicle{"V-101", 58.2, 76, 70, "Lexington", "CHARGING"}
	var vehicle3 = Vehicle{"V-102", 76.9, 95, 80, "Alabama", "OFFLINE"}
	var vehicle4 = Vehicle{"V-103", 28.23, 26, 90, "New York", "ACTIVE"}
	var vehicle5 = Vehicle{"V-104", 80.0, 33, 59, "Oklahoma", "ACTIVE"}
	
	vehicles := []Vehicle{vehicle1, vehicle2, vehicle3, vehicle4, vehicle5}		// Populate vehicles slice
	
	for _, vehicle := range vehicles {
		var isBatteryLow = IsBatteryLow(vehicle)
		var batteryStatus = GetBatteryStatus(vehicle)
		var isOverheating = IsOverheating(vehicle)
		
		// Print characteristics of vehicle object
		GetVehicleID(vehicle)
		GetVehicleSpeed(vehicle)
		GetVehicleLocation(vehicle)
		fmt.Println("Battery Low:", isBatteryLow)
		fmt.Println("Battery Status:", batteryStatus)
		fmt.Println("Overheating:", isOverheating)
		fmt.Println("Vehicle Status:", VehicleStatus(vehicle))
		fmt.Println()
	}
}		// End of main


func GetVehicleID(vehicle Vehicle) {
	fmt.Println("Vehicle ID:", vehicle.ID)
}

func IsBatteryLow(vehicle Vehicle) bool {
	if (vehicle.Battery < LOW_BATTERY_THRESHOLD){
		return true
	} else {
		return false
	}
}

func GetBatteryStatus(vehicle Vehicle) string {
	if (vehicle.Battery >= LOW_BATTERY_THRESHOLD) {
		return "NORMAL"
	} else {
		return "LOW_BATTERY"
	}
}

func GetVehicleSpeed(vehicle Vehicle) {
	fmt.Println("Speed:", vehicle.Speed)
}

func IsOverheating(vehicle Vehicle) bool {
	return vehicle.Temperature >= MAX_TEMP
}

func GetVehicleLocation(vehicle Vehicle){
	fmt.Println("Location:", vehicle.Location)
}

func VehicleStatus(vehicle Vehicle) string {
	if IsBatteryLow(vehicle){
		return "LOW_BATTERY"
	} else if IsOverheating(vehicle){
		return "OVERHEATING"
	} else {
		return "HEALTHY"
	}
}


