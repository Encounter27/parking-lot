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

	// tmp := `{"parking_name": "phoenix mall parking","floors":[{"total":2,"space":[{"from":1,"to":2,"type":"DEFAULT"}]}]}`
	// var cfg1 Config
	// json.Unmarshal([]byte(tmp), &cfg)

	// floor, err := createFloorPlan(cfg.Floors[0])
}
