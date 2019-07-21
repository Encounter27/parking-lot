package command

import (
	"fmt"
	"reflect"
	"testing"
)

func TestCreateCollection(t *testing.T) {
	var all1, all2 collection

	go func() {
		all1 = CreateCollection()
	}()

	go func() {
		all2 = CreateCollection()
	}()

	t.Run("Test create singleton command collection", func(t *testing.T) {
		// both the singleton object should have the same address and same structure
		if !reflect.DeepEqual(&all1, &all2) {
			t.Errorf("CreateCollection() = %v, want %v", all1, all2)
		}
	})
}

func Test_collection_RegisterCommand(t *testing.T) {
	all := CreateCollection()
	name := "create_parking_lot"

	t.Run("Register command", func(t *testing.T) {
		if err := all.RegisterCommand(name, &CreateParkingLot{}); err != nil {
			t.Errorf("collection.RegisterCommand() error = %v", err)
		}
	})

	expectedError := fmt.Errorf("Command %s already registered", name)
	t.Run("Register command", func(t *testing.T) {
		if err := all.RegisterCommand(name, &CreateParkingLot{}); err.Error() != expectedError.Error() {
			t.Errorf("collection.RegisterCommand() error = %v expectedError %v", err, expectedError)
		}
	})
}

func Test_collection_Get(t *testing.T) {
	all := CreateCollection()
	name := "create_parking_lot"

	t.Run("Register command", func(t *testing.T) {
		if err := all.RegisterCommand(name, &CreateParkingLot{}); err != nil {
			t.Errorf("collection.RegisterCommand() error = %v", err)
		}
	})

	t.Run("Get command", func(t *testing.T) {
		cmd, err := all.Get(name)
		_, ok := cmd.(*CreateParkingLot)
		if err != nil || !ok {
			t.Errorf("collection.Get() error = %v ", err)
		}
	})

	expectedError := fmt.Errorf("Command xyz not registered")

	t.Run("Get command", func(t *testing.T) {
		_, err := all.Get("xyz")
		if err.Error() != expectedError.Error() {
			t.Errorf("collection.Get() error = %v expectedError = %v", err, expectedError)
		}
	})
}
