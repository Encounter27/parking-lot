package parking

import "fmt"

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

type iSlot interface {
	GetType() SlotType
}

type slot struct {
	Type SlotType
}

func (s slot) GetType() SlotType {
	return s.Type
}

type slotTwoWheeler struct {
	Type SlotType
}

func (s slotTwoWheeler) GetType() SlotType {
	return s.Type
}

type slotFourWheeler struct {
	Type SlotType
}

func (s slotFourWheeler) GetType() SlotType {
	return s.Type
}

type slotBigVehicle struct {
	Type SlotType
}

func (s slotBigVehicle) GetType() SlotType {
	return s.Type
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
