package parking

import (
	"encoding/json"
	"testing"
)

func Test_createParkingLot(t *testing.T) {
	tmp := `{"parking_name": "phoenix mall parking","floors":[{"total":2,"space":[{"from":1,"to":2,"type":"DEFAULT"}]}]}`
	var cfg Config
	json.Unmarshal([]byte(tmp), &cfg)

	l, _ := createParkingLot(cfg)

	t.Run("Create parking lot", func(t *testing.T) {
		if l.GetTotal() != 2 {
			t.Errorf("Created total slot = %v, want %v", l.GetTotal(), 2)
		}
	})
}

func TestAddNewParking(t *testing.T) {
	tmp := `{"parking_name": "phoenix mall parking","floors":[{"total":2,"space":[{"from":1,"to":2,"type":"DEFAULT"}]}]}`
	var cfg Config
	json.Unmarshal([]byte(tmp), &cfg)

	err := AddNewParking(cfg)

	t.Run("Create parking lot", func(t *testing.T) {
		if err != nil {
			t.Errorf(err.Error())
		}
	})

	err = AddNewParking(cfg)

	t.Run("Create parking lot", func(t *testing.T) {
		if err == nil {
			t.Errorf("Able to create parking with duplicate name")
		}
	})
}

func Test_lot_GetTotal(t *testing.T) {
	tmp := `{"parking_name": "phoenix mall parking","floors":[{"total":2,"space":[{"from":1,"to":2,"type":"DEFAULT"}]}]}`
	var cfg Config
	json.Unmarshal([]byte(tmp), &cfg)

	l, _ := createParkingLot(cfg)

	t.Run("Get total parking slot", func(t *testing.T) {
		if l.GetTotal() != 2 {
			t.Errorf("Created total slot = %v, want %v", l.GetTotal(), 2)
		}
	})
}
