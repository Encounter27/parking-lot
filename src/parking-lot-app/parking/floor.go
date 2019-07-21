package parking

type floor struct {
	spot      []iSlot
	available []int
}

func createFloorPlan(cfg FloorConfig) (floor, error) {
	var f floor
	f.available = make([]int, TotalSlotTypes)
	f.spot = make([]iSlot, cfg.Total+1) // +1 to simplify and treat array index starting from 1

	var err error
	for _, i := range cfg.Space {
		f.available[i.Type] = i.To - i.From + 1

		for e := i.From; e <= i.To; e++ {
			if f.spot[e], err = createSlot(i.Type); err != nil {
				return floor{}, err
			}
		}
	}

	return f, nil
}
