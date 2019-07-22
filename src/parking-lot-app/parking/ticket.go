package parking

import "time"
import "github.com/google/uuid"

type TicketStatus int

const (
	ACTIVE TicketStatus = 0
	PAID   TicketStatus = 1
)

type Ticket struct {
	No          uuid.UUID
	ParkingName string
	Floor       int // Which floor
	Slot        int // Which spot
	EntryTime   time.Time
	ExitTime    time.Time
	State       TicketStatus
	Charge      int
}
