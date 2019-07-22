package main

import (
	"bufio"
	"fmt"
	"os"
	"parking-lot-app/command"
	"strings"
)

var (
	all = command.CreateCollection()
)

func init() {
	// Register all the actions, user going to perform
	all.RegisterCommand("create_parking_lot", &command.CreateParkingLot{})
	all.RegisterCommand("park", &command.Park{})
	all.RegisterCommand("leave", &command.Leave{})
	all.RegisterCommand("status", &command.Status{})
	all.RegisterCommand("registration_numbers_for_cars_with_colour", &command.GetRegNumbers{})
	all.RegisterCommand("slot_numbers_for_cars_with_colour", &command.GetSlotNumbers{QueryBy: command.BY_COLOR})
	all.RegisterCommand("slot_number_for_registration_number", &command.GetSlotNumbers{QueryBy: command.BY_VEHICLE_NO})
}

func print(resp []string) {
	for _, s := range resp {
		fmt.Println(s)
	}
}

func main() {
	snr := bufio.NewScanner(os.Stdin)
	for {
		snr.Scan()
		line := snr.Text()
		if len(line) == 0 || line == "exit" {
			break
		}
		fields := strings.Fields(line)
		//fmt.Printf("Fields: %q\n", fields)

		if cmd, err := all.Get(fields[0]); err == nil {
			resp, err := cmd.Execute(fields)
			if err != nil {
				fmt.Println(err)

			} else {
				print(resp)
			}
		} else {
			fmt.Println(err)
		}
	}
}
