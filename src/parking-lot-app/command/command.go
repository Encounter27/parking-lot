package command

import (
	"fmt"
	"parking-lot-app/parking"
	"strconv"
)

type ICommand interface {
	Execute(args []string) error // For simplicity not returning the response of the command, instead will print to stdout.
}

type CreateParkingLot struct {
}

func (c *CreateParkingLot) Execute(args []string) error {
	size, err := strconv.Atoi(args[0])
	if err != nil {
		fmt.Printf("Failed to convert argument [%s] to integer", args[0])
	}

	cfg := parking.GetSimpleParkingConfig(size)
	_ = cfg

	return nil
}

type Park struct {
}

func (c *Park) Execute(args []string) error {
	return nil
}
