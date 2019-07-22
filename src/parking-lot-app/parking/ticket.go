package parking

import "time"
import "github.com/google/uuid"

type Status int

const (
	ACTIVE Status = 0
	PAID   Status = 1
)

type Ticket struct {
	No          uuid.UUID
	ParkingName string
	Floor       int // Which floor
	Slot        int // Which spot
	EntryTime   time.Time
	ExitTime    time.Time
	State       Status
	Charge      int
}
