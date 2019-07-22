package parking

import (
	"encoding/json"
	"testing"
)

func Test_createFloorPlan(t *testing.T) {
	tmp := `{"parking_name": "phoenix mall parking","floors":[{"total":2,"space":[{"from":1,"to":2,"type":"DEFAULT"}]}]}`
	var cfg Config
	json.Unmarshal([]byte(tmp), &cfg)

	floor, err := createFloorPlan(cfg.Floors[0])

	t.Run("Create floor plan", func(t *testing.T) {
		if err != nil {
			t.Error(err.Error())
		}

		if len(floor.spot)-1 != 2 {
			t.Errorf("Created total slot = %v, want %v", len(floor.spot)-1, 2)
		}
	})
}

func Test_floor_getNearestAvaibleSlot(t *testing.T) {
	tmp := `{"parking_name": "phoenix mall parking","floors":[{"total":2,"space":[{"from":1,"to":2,"type":"DEFAULT"}]}]}`
	var cfg Config
	json.Unmarshal([]byte(tmp), &cfg)

	floor, _ := createFloorPlan(cfg.Floors[0])

	t.Run("Get nearest slots", func(t *testing.T) {
		for i := 1; i <= 2; i++ {
			c := floor.getNearestAvaibleSlot(DEFAULT)
			if c != i {
				t.Errorf("Created total slot = %v, want %v", c, i)
			}
		}

		c := floor.getNearestAvaibleSlot(DEFAULT)
		if c != 0 {
			t.Errorf("Created total slot = %v, want %v", c, 0)
		}

		c = floor.getNearestAvaibleSlot(TWO_WHEELER)
		if c != 0 {
			t.Errorf("Created total slot = %v, want %v", c, 0)
		}
	})
}
