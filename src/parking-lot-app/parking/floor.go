package parking

import (
	"container/heap"
	"parking-lot-app/utils"
)

type floor struct {
	spot      []iSlot
	available []*utils.PriorityQueue
}

func createFloorPlan(cfg FloorConfig) (floor, error) {
	var f floor
	f.available = make([]*utils.PriorityQueue, TotalSlotTypes)
	f.spot = make([]iSlot, cfg.Total+1) // +1 to simplify and treat array index starting from 1

	for i := 0; i < int(TotalSlotTypes); i++ {
		f.available[i] = new(utils.PriorityQueue)
		heap.Init(f.available[i])
	}

	var err error
	for _, i := range cfg.Space {
		for x := i.From; x <= i.To; x++ {
			heap.Push(f.available[i.Type], x)
		}

		for e := i.From; e <= i.To; e++ {
			if f.spot[e], err = createSlot(i.Type); err != nil {
				return floor{}, err
			}
		}
	}

	return f, nil
}

func (f floor) getNearestAvaibleSlot(t SlotType) int {
	if f.available[t].Len() > 0 {
		return heap.Pop(f.available[t]).(int)
	}

	return 0
}
