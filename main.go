package main
import "fmt"

type Vehicle struct {
	ID string 
	Battery float64
	Speed float64
	Temperature float64
	Location string
	Status string
}


func main() {

    var object1 Vehicle 

    var vehicleID = [5]string{"V-100", "V-101", "V-102", "V-103", "V-104"}
    var vehicleBattery = [5]float64{99.0, 58.2, 76.9, 28.23, 80.0}
    var vehicleSpeed = [5]float64{42, 76, 95, 26, 33}
    var vehicleTemp = [5]float64{100, 70, 80, 90, 59}
    var vehicleLocation = [5]string{"Austin", "Lexington", "Alabama", "New York", "Oklahoma"}
    var vehicleStatus = [5]string{"ACTIVE", "CHARGING", "OFFLINE", "ACTIVE", "ACTIVE"}

    
    for i:=0; i < 5; i++ {
        object1.ID = vehicleID[i]
        object1.Battery = vehicleBattery[i]
        object1.Speed = vehicleSpeed[i]
        object1.Temperature = vehicleTemp[i]
        object1.Location = vehicleLocation[i]
        object1.Status = vehicleStatus[i]
        
        fmt.Println("Vehicle", i)
        fmt.Println("Vehicle: " + object1.ID)
        fmt.Println("Battery:",  object1.Battery)
        fmt.Println("Speed:",  object1.Speed)
        fmt.Println("Temperature:",  object1.Temperature)
        fmt.Println("Location: " + object1.Location)
        fmt.Println("Status: " + object1.Status)

        fmt.Println()
    }


}




