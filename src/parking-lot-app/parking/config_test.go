package parking

import (
	"encoding/json"
	"reflect"
	"testing"
)

func TestGetSimpleParkingConfig(t *testing.T) {
	cfg1 := GetSimpleParkingConfig(6)
	tmp := `{"parking_name": "phoenix mall parking","floors":[{"total":6,"space":[{"from":1,"to":6,"type":"DEFAULT"}]}]}`
	var cfg2 Config
	json.Unmarshal([]byte(tmp), &cfg2)

	t.Run("Get simple config", func(t *testing.T) {
		if !reflect.DeepEqual(&cfg1, &cfg2) {
			t.Errorf("GetSimpleParkingConfig() = %v, want %v", cfg1, cfg2)
		}
	})
}

func TestGetMultiFloorParkingConfig(t *testing.T) {
	tmp := `{"parking_name": "phoenix mall parking","floors":[{"total":6,"space":[{"from":1,"to":6,"type":"TWO_WHEELER"}]}]}`
	cfg1 := GetMultiFloorParkingConfig(tmp)
	var cfg2 Config
	json.Unmarshal([]byte(tmp), &cfg2)

	t.Run("Get simple config", func(t *testing.T) {
		if !reflect.DeepEqual(&cfg1, &cfg2) {
			t.Errorf("GetMultiFloorParkingConfig() = %v, want %v", cfg1, cfg2)
		}
	})
}
