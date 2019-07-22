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
	type args struct {
		args []string
	}
	tests := []struct {
		name    string
		c       *Park
		args    args
		want    []string
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Park{}
			got, err := c.Execute(tt.args.args)
			if (err != nil) != tt.wantErr {
				t.Errorf("Park.Execute() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Park.Execute() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestLeave_Execute(t *testing.T) {
	type args struct {
		args []string
	}
	tests := []struct {
		name    string
		c       *Leave
		args    args
		want    []string
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Leave{}
			got, err := c.Execute(tt.args.args)
			if (err != nil) != tt.wantErr {
				t.Errorf("Leave.Execute() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Leave.Execute() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestStatus_Execute(t *testing.T) {
	type args struct {
		args []string
	}
	tests := []struct {
		name    string
		c       *Status
		args    args
		want    []string
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Status{}
			got, err := c.Execute(tt.args.args)
			if (err != nil) != tt.wantErr {
				t.Errorf("Status.Execute() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Status.Execute() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetRegNumbers_Execute(t *testing.T) {
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
		// TODO: Add test cases.
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
		want    []string
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &GetSlotNumbers{
				QueryBy: tt.fields.QueryBy,
			}
			got, err := c.Execute(tt.args.args)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetSlotNumbers.Execute() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetSlotNumbers.Execute() = %v, want %v", got, tt.want)
			}
		})
	}
}
