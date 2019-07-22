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
	floor := 0                     // Again hardcoded as given problem dont have any floor input
	reg := args[1]
	color := args[2]

	return nil, nil
}

type Leave struct {
}

func (c *Leave) Execute(args []string) ([]string, error) {
	return nil, nil
}

type Status struct {
}

func (c *Status) Execute(args []string) ([]string, error) {
	return nil, nil
}

type GetRegNumbers struct {
}

func (c *GetRegNumbers) Execute(args []string) ([]string, error) {
	return nil, nil
}

const (
	BY_COLOR      = 0
	BY_VEHICLE_NO = 1
)

type GetSlotNumbers struct {
	QueryBy int
}

func (c *GetSlotNumbers) Execute(args []string) ([]string, error) {
	return nil, nil
}
