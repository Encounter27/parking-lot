package parking

import (
	"bytes"
	"container/heap"
	"fmt"
	"sync"
	"time"

	"github.com/google/uuid"
)

type lot struct {
	floors []floor
}

func (l lot) GetTotal() int {
	total := 0

	for _, f := range l.floors {
		total += len(f.spot) - 1
	}

	return total
}

// Considering the fact that the parking management software will be completely
// won by parking woner. Parking woner can have more than one parking in the city
// and he wants to centralize the management of all parking he owns.
var allParkings map[string]lot

var once sync.Once
var lock = &sync.Mutex{}

func createParkingLot(cfg Config) (lot, error) {
	var l lot

	l.floors = make([]floor, len(cfg.Floors))

	var err error
	for i, _ := range cfg.Floors {
		if l.floors[i], err = createFloorPlan(cfg.Floors[i]); err != nil {
			return l, err
		}
	}

	return l, nil
}

func AddNewParking(cfg Config) error {
	once.Do(func() { // atomic, does not allow repeating
		allParkings = make(map[string]lot)
	})

	lock.Lock()
	defer lock.Unlock()

	if _, ok := allParkings[cfg.Name]; ok {
		return fmt.Errorf("Failed to create parking, parking with name %s already exists", cfg.Name)
	}

	l, err := createParkingLot(cfg)

	if err != nil {
		return err
	}

	allParkings[cfg.Name] = l

	return nil
}

// Ignoring validation for anvalid type as of now.
func getSlotType(t CarType) SlotType {
	var st SlotType

	switch t {
	case CAR:
		st = DEFAULT
	case MOTORBIKE:
		st = TWO_WHEELER
	case SEDAN, SUV:
		st = FOUR_WHEELER
	case TRUCK, BUS:
		st = BIG_VEHIHLE
	}

	return st
}

func AssignTicket(parkingName string, car IVehicle) (Ticket, error) {
	var t Ticket
	parking, ok := allParkings[parkingName]
	if !ok {
		return t, fmt.Errorf("Parking with name %s doesn't exists", parkingName)
	}

	st := getSlotType(car.GetType())
	var slotNo int
	fc := 0
	for _, f := range parking.floors {
		slotNo = f.getNearestAvaibleSlot(st)
		if slotNo != 0 {
			break
		}
		fc++
	}

	if slotNo == 0 {
		return t, fmt.Errorf("Sorry, parking lot is full")
	}

	parking.floors[fc].spot[slotNo] = slot{Vehicle: car, Type: st}

	// Generate a ticket
	t.No = uuid.New()
	t.ParkingName = parkingName
	t.Floor = fc
	t.Slot = slotNo
	t.EntryTime = time.Now()
	t.State = ACTIVE
	t.Charge = 0

	// Ticket can be store for audit purpouse or history management

	return t, nil
}

func FreeParkingSlot(parkingName string, floorNo int, slotNo int) error {
	parking, ok := allParkings[parkingName]
	if !ok {
		return fmt.Errorf("Parking with name %s doesn't exists", parkingName)
	}

	if nil == parking.floors[floorNo].spot[slotNo] {
		return fmt.Errorf("No vehicle parked in the slot %d", slotNo)
	}

	car := parking.floors[floorNo].spot[slotNo].GetVehicle()
	parking.floors[floorNo].spot[slotNo] = nil
	heap.Push(parking.floors[floorNo].available[getSlotType(car.GetType())], slotNo)

	// change the state of the ticket and charge accordingly
	_ = car

	return nil
}

func Status(parkingName string) ([]string, error) {
	parking, ok := allParkings[parkingName]
	if !ok {
		return nil, fmt.Errorf("Parking with name %s doesn't exists", parkingName)
	}

	var response []string
	response = append(response, "Slot No.    Registration No    Colour")

	for _, f := range parking.floors {
		for i, es := range f.spot {
			if es != nil {
				if car := es.GetVehicle(); car != nil {
					response = append(response, fmt.Sprintf("%d           %s      %s", i, car.GetRegNO(), car.GetColor()))
				}
			}
		}
	}

	return response, nil
}

func GetAllCarRegNoByColour(parkingName string, colour string) ([]string, error) {
	parking, ok := allParkings[parkingName]
	if !ok {
		return nil, fmt.Errorf("Parking with name %s doesn't exists", parkingName)
	}

	var response []string

	var buffer bytes.Buffer
	for _, f := range parking.floors {
		for _, es := range f.spot {
			if es != nil {
				if car := es.GetVehicle(); car != nil && colour == car.GetColor() {
					tmp := fmt.Sprintf("%s, ", car.GetRegNO())
					buffer.WriteString(tmp)
				}
			}
		}
	}
	tmp := buffer.String()
	if tmp == "" {
		return response, nil
	}

	response = append(response, tmp[:len(tmp)-2])

	return response, nil
}

func GetAllSlotsByColour(parkingName string, colour string) ([]string, error) {
	parking, ok := allParkings[parkingName]
	if !ok {
		return nil, fmt.Errorf("Parking with name %s doesn't exists", parkingName)
	}

	var response []string

	var buffer bytes.Buffer
	for _, f := range parking.floors {
		for i, es := range f.spot {
			if es != nil {
				if car := es.GetVehicle(); car != nil && colour == car.GetColor() {
					tmp := fmt.Sprintf("%d, ", i)
					buffer.WriteString(tmp)
				}
			}
		}
	}
	tmp := buffer.String()
	if tmp == "" {
		return response, nil
	}

	response = append(response, tmp[:len(tmp)-2])

	return response, nil
}

func GetSlotNoByReg(parkingName string, reg string) ([]string, error) {
	parking, ok := allParkings[parkingName]
	if !ok {
		return nil, fmt.Errorf("Parking with name %s doesn't exists", parkingName)
	}

	var response []string

	var buffer bytes.Buffer
	for _, f := range parking.floors {
		for i, es := range f.spot {
			if es != nil {
				if car := es.GetVehicle(); car != nil && reg == car.GetRegNO() {
					tmp := fmt.Sprintf("%d, ", i)
					buffer.WriteString(tmp)
					break
				}
			}
		}
	}
	tmp := buffer.String()
	if tmp == "" {
		response = append(response, "Not found")
		return response, nil
	}

	response = append(response, tmp[:len(tmp)-2])

	return response, nil
}
