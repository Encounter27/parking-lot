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

func Test_getSlotType(t *testing.T) {
	type args struct {
		t CarType
	}
	tests := []struct {
		name string
		args args
		want SlotType
	}{
		{"test 1", args{CAR}, DEFAULT},
		{"test 1", args{MOTORBIKE}, TWO_WHEELER},
		{"test 1", args{SEDAN}, FOUR_WHEELER},
		{"test 1", args{SUV}, FOUR_WHEELER},
		{"test 1", args{TRUCK}, BIG_VEHIHLE},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getSlotType(tt.args.t); got != tt.want {
				t.Errorf("getSlotType() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAssignTicket(t *testing.T) {
	tmp := `{"parking_name": "phoenix mall parking","floors":[{"total":2,"space":[{"from":1,"to":2,"type":"DEFAULT"}]}]}`
	var cfg Config
	json.Unmarshal([]byte(tmp), &cfg)

	AddNewParking(cfg)

	type args struct {
		parkingName string
		car         IVehicle
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			"test assign parking",
			args{
				parkingName: "phoenix mall parking",
				car: &Car{
					Type:   CAR,
					RegID:  "hdsvh",
					Colour: "white",
				},
			},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := AssignTicket(tt.args.parkingName, tt.args.car)
			if (err != nil) != tt.wantErr {
				t.Errorf("AssignTicket() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}
