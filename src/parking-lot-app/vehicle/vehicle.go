package vehicle

import "parking-lot-app/parking"

type Type int

// Default -> Its only to simplify as in given problem no vehicle type is mentioned
// Eventualy vehicle type can be extended for different vehicle types.
const (
	CAR       Type = 0
	MOTORBIKE Type = 1
	SEDAN     Type = 2
	SUV       Type = 3
	TRUCK     Type = 4
	BUS       Type = 5
)

type IVehicle interface {
	Park() parking.Ticket
}

// Car is just a default concrete representation of vehicle.
type Car struct {
	Type  Type
	RegID string
	Color string
}
