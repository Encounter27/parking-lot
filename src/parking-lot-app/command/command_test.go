package command

import (
	"reflect"
	"testing"
)

func TestCreateParkingLot_Execute(t *testing.T) {

	type args struct {
		args []string
	}
	tests := []struct {
		name    string
		c       *CreateParkingLot
		args    args
		want    []string
		wantErr bool
	}{
		{
			"Create parking lot",
			&CreateParkingLot{},
			args{args: []string{"create_parking_lot", "6"}},
			[]string{"Created a parking lot with 6 slots"},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &CreateParkingLot{}
			got, err := c.Execute(tt.args.args)
			if (err != nil) != tt.wantErr {
				t.Errorf("CreateParkingLot.Execute() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CreateParkingLot.Execute() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPark_Execute(t *testing.T) {
	c := new(CreateParkingLot)
	arg := []string{"create_parking_lot", "6"}
	c.Execute(arg)

	type args struct {
		args []string
	}
	tests := []struct {
		name    string
		c       *Park
		args    args
		wantErr bool
	}{
		{
			"Park",
			&Park{},
			args{args: []string{"park", "KA-01-HH-1234", "White"}},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Park{}
			_, err := c.Execute(tt.args.args)
			if (err != nil) != tt.wantErr {
				t.Errorf("Park.Execute() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}

func TestLeave_Execute(t *testing.T) {
	c := new(CreateParkingLot)
	arg := []string{"create_parking_lot", "6"}
	c.Execute(arg)

	type args struct {
		args []string
	}
	tests := []struct {
		name    string
		c       *Leave
		args    args
		wantErr bool
	}{
		{
			"Park",
			&Leave{},
			args{args: []string{"leave", "1"}},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Leave{}
			_, err := c.Execute(tt.args.args)
			if (err != nil) != tt.wantErr {
				t.Errorf("Leave.Execute() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}

func TestStatus_Execute(t *testing.T) {
	c := new(CreateParkingLot)
	arg := []string{"create_parking_lot", "6"}
	c.Execute(arg)

	type args struct {
		args []string
	}
	tests := []struct {
		name    string
		c       *Status
		args    args
		wantErr bool
	}{
		{
			"status",
			&Status{},
			args{args: []string{"status"}},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Status{}
			_, err := c.Execute(tt.args.args)
			if (err != nil) != tt.wantErr {
				t.Errorf("Status.Execute() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}

func TestGetRegNumbers_Execute(t *testing.T) {
	c := new(CreateParkingLot)
	arg := []string{"create_parking_lot", "6"}
	c.Execute(arg)

	p := new(Park)
	arg = []string{"park", "PP-01-HH-1234", "White"}
	p.Execute(arg)

	type args struct {
		args []string
	}
	tests := []struct {
		name    string
		c       *GetRegNumbers
		args    args
		want    []string
		wantErr bool
	}{
		{
			"GetRegNumbers",
			&GetRegNumbers{},
			args{args: []string{"registration_numbers_for_cars_with_colour", "White"}},
			[]string{"PP-01-HH-1234"},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &GetRegNumbers{}
			got, err := c.Execute(tt.args.args)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetRegNumbers.Execute() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetRegNumbers.Execute() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetSlotNumbers_Execute(t *testing.T) {
	c := new(CreateParkingLot)
	arg := []string{"create_parking_lot", "6"}
	c.Execute(arg)

	p := new(Park)
	arg = []string{"park", "PX-01-HH-123&", "White"}
	p.Execute(arg)

	type fields struct {
		QueryBy int
	}
	type args struct {
		args []string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			"GetSlotNumbers",
			fields{QueryBy: BY_COLOR},
			args{args: []string{"slot_numbers_for_cars_with_colour", "White"}},
			false,
		},
		{
			"GetSlotNumbers",
			fields{QueryBy: BY_VEHICLE_NO},
			args{args: []string{"slot_numbers_for_cars_with_colour", "White"}},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &GetSlotNumbers{
				QueryBy: tt.fields.QueryBy,
			}
			_, err := c.Execute(tt.args.args)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetSlotNumbers.Execute() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}
