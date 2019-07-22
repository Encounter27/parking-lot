package parking

import (
	"fmt"
	"parking-lot-app/vehicle"
	"sync"
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

func AssignTicket(parkingName string, car vehicle.IVehicle) error {
	floor := 0 // Again hardcoded as given problem dont have any floor input
	return nil
}
