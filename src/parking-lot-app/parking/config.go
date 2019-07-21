package parking

import (
	"encoding/json"
	"fmt"
)

/* Parking config is important so that we can create a parking with multiple floors.
And ofcouse each floor can have its own parking space configurations.

{
  "parking_name": "phoenix mall parking",
  "floors": [
    // Floor1 config
    {
      "total": 20,
      "space": [
        {
          "from": 1,
          "to": 10,
          "type": "TWO_WHEELER"
        },
        {
          "from": 11,
          "to": 20,
          "type": "FOUR_WHEELER"
        }
      ]
    },
    // Floor2 config
    {
      "total": 15,
      "space": [
        {
          "from": 1,
          "to": 10,
          "type": "FOUR_WHEELER"
        },
        {
          "from": 11,
          "to": 15,
          "type": "BIG_VEHIHLE"
        }
      ]
    }
  ]
}
*/

type SpaceConfig struct {
	From int      `json:"from"`
	To   int      `json:"to"`
	Type SlotType `json:"type"`
}

type FloorConfig struct {
	Total int           `json:"total"`
	Space []SpaceConfig `json:"space"`
}

type Config struct {
	Name   string        `json:"parking_name"`
	Floors []FloorConfig `json:"floors"`
}

// GetSimple, well this is to considered the fact that, for this assignment input is only parking Size
func GetSimpleParkingConfig(size int) Config {
	var cfg Config

	cfg.Floors = make([]FloorConfig, 0)
	floor := FloorConfig{}
	floor.Total = size
	floor.Space = make([]SpaceConfig, 0)
	floor.Space = append(floor.Space, SpaceConfig{
		From: 1,
		To:   size,
		Type: DEFAULT,
	})

	cfg.Floors = append(cfg.Floors, floor)
	cfg.Name = "phoenix mall parking"

	return cfg
}

// GetMultiFloorParkingConfig
func GetMultiFloorParkingConfig(configJson string) Config {
	var cfg Config

	err := json.Unmarshal([]byte(configJson), &cfg)

	if err != nil {
		fmt.Println(err)
	}

	return cfg
}
