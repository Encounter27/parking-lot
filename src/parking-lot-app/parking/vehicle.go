package parking

type CarType int

// Default -> Its only to simplify as in given problem no vehicle type is mentioned
// Eventualy vehicle type can be extended for different vehicle types.
const (
	CAR       CarType = 0
	MOTORBIKE CarType = 1
	SEDAN     CarType = 2
	SUV       CarType = 3
	TRUCK     CarType = 4
	BUS       CarType = 5
)

type IVehicle interface {
	Park(parkingName string) (Ticket, error)
	GetType() CarType
	GetRegNO() string
	GetColor() string
}

// Car is just a default concrete representation of vehicle.
type Car struct {
	Type   CarType
	RegID  string
	Colour string
}

func (c *Car) Park(parkingName string) (Ticket, error) {
	t, err := AssignTicket(parkingName, c)
	return t, err
}

func (c Car) GetType() CarType {
	return c.Type
}

func (c Car) GetRegNO() string {
	return c.RegID
}

func (c Car) GetColor() string {
	return c.Colour
}
