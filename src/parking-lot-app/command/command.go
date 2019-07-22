package command

import (
	"fmt"
	"parking-lot-app/parking"
	"strconv"
)

type ICommand interface {
	Execute(args []string) ([]string, error) // For simplicity not returning the response of the command, instead will print to stdout.
}

type CreateParkingLot struct {
}

func (c *CreateParkingLot) Execute(args []string) ([]string, error) {
	size, err := strconv.Atoi(args[1])
	if err != nil {
		fmt.Printf("Failed to convert argument [%s] to integer", args[1])
	}

	cfg := parking.GetSimpleParkingConfig(size)
	if err := parking.AddNewParking(cfg); err != nil {
		return nil, err
	}

	var response []string

	response = append(response, fmt.Sprintf("Created a parking lot with %d slots", size))

	return response, nil
}

type Park struct {
}

func (c *Park) Execute(args []string) ([]string, error) {
	name := "phoenix mall parking" // As of now hardcoded, but can be expected as argument

	car := parking.Car{
		Type:   parking.CAR,
		RegID:  args[1],
		Colour: args[2],
	}

	var response []string
	t, err := car.Park(name)
	if err != nil {
		response = append(response, fmt.Sprintf(err.Error()))

		return response, err
	}

	response = append(response, fmt.Sprintf("Allocated slot number: %d", t.Slot))

	return response, nil
}

type Leave struct {
}

func (c *Leave) Execute(args []string) ([]string, error) {
	slotNo, err := strconv.Atoi(args[1])
	if err != nil {
		fmt.Printf("Failed to convert argument [%s] to integer", args[1])
	}

	var response []string

	if err = parking.FreeParkingSlot("phoenix mall parking", 0, slotNo); err != nil {
		return response, err
	}

	response = append(response, fmt.Sprintf("Slot number %d is free", slotNo))

	return response, nil
}

type Status struct {
}

func (c *Status) Execute(args []string) ([]string, error) {
	name := "phoenix mall parking" // As of now hardcoded, but can be expected as argument
	resp, err := parking.Status(name)

	return resp, err
}

type GetRegNumbers struct {
}

func (c *GetRegNumbers) Execute(args []string) ([]string, error) {
	name := "phoenix mall parking" // As of now hardcoded, but can be expected as argument
	resp, err := parking.GetAllCarRegNoByColour(name, args[1])

	return resp, err
}

const (
	BY_COLOR      = 0
	BY_VEHICLE_NO = 1
)

type GetSlotNumbers struct {
	QueryBy int
}

func (c *GetSlotNumbers) Execute(args []string) ([]string, error) {
	name := "phoenix mall parking" // As of now hardcoded, but can be expected as argument

	switch c.QueryBy {
	case BY_COLOR:
		return parking.GetAllSlotsByColour(name, args[1])
	case BY_VEHICLE_NO:
		return parking.GetSlotNoByReg(name, args[1])
	}

	return nil, fmt.Errorf("Invalid query")
}
