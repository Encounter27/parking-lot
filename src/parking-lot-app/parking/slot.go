package parking

import (
	"encoding/json"
	"fmt"
)

type SlotType int

// Default -> Its only to simplify as in given problem no parking space type is mentioned.
// Eventualy the parking slot can be extended for different slot types.
const (
	DEFAULT        SlotType = 0
	TWO_WHEELER    SlotType = 1
	FOUR_WHEELER   SlotType = 2
	BIG_VEHIHLE    SlotType = 3
	TotalSlotTypes SlotType = 4
)

func (s *SlotType) UnmarshalJSON(b []byte) error {
	var str string
	if err := json.Unmarshal(b, &str); err != nil {
		return err
	}

	switch str {
	default:
		return fmt.Errorf("%v is not a valid parking slot type", s)
	case "DEFAULT":
		*s = DEFAULT
	case "TWO_WHEELER":
		*s = TWO_WHEELER
	case "FOUR_WHEELER":
		*s = FOUR_WHEELER

	case "BIG_VEHIHLE":
		*s = BIG_VEHIHLE
	}

	return nil
}

type iSlot interface {
	GetType() SlotType
	GetVehicle() IVehicle
}

type slot struct {
	Vehicle IVehicle
	Type    SlotType
}

func (s slot) GetType() SlotType {
	return s.Type
}

func (s slot) GetVehicle() IVehicle {
	return s.Vehicle
}

type slotTwoWheeler struct {
	Vehicle IVehicle
	Type    SlotType
}

func (s slotTwoWheeler) GetType() SlotType {
	return s.Type
}

func (s slotTwoWheeler) GetVehicle() IVehicle {
	return s.Vehicle
}

type slotFourWheeler struct {
	Vehicle IVehicle
	Type    SlotType
}

func (s slotFourWheeler) GetType() SlotType {
	return s.Type
}

func (s slotFourWheeler) GetVehicle() IVehicle {
	return s.Vehicle
}

type slotBigVehicle struct {
	Vehicle IVehicle
	Type    SlotType
}

func (s slotBigVehicle) GetType() SlotType {
	return s.Type
}

func (s slotBigVehicle) GetVehicle() IVehicle {
	return s.Vehicle
}

func createSlot(t SlotType) (iSlot, error) {
	switch t {
	case DEFAULT:
		return &slot{Type: t}, nil

	case TWO_WHEELER:
		return &slotTwoWheeler{Type: t}, nil

	case FOUR_WHEELER:
		return &slotFourWheeler{Type: t}, nil

	case BIG_VEHIHLE:
		return &slotBigVehicle{Type: t}, nil
	}

	return nil, fmt.Errorf("Slot type %q is not avilable", t)
}
