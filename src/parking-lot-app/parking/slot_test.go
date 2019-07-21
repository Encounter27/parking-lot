package parking

import (
	"encoding/json"
	"reflect"
	"testing"
)

func Test_createSlot(t *testing.T) {
	type args struct {
		t SlotType
	}
	tests := []struct {
		name    string
		args    args
		want    iSlot
		wantErr bool
	}{
		{"Create slot for DEFAULT type", args{t: DEFAULT}, &slot{Type: DEFAULT}, false},
		{"Create slot for TWO_WHEELER type", args{t: TWO_WHEELER}, &slotTwoWheeler{Type: TWO_WHEELER}, false},
		{"Create slot for FOUR_WHEELER type", args{t: FOUR_WHEELER}, &slotFourWheeler{Type: FOUR_WHEELER}, false},
		{"Create slot for BIG_VEHIHLE type", args{t: BIG_VEHIHLE}, &slotBigVehicle{Type: BIG_VEHIHLE}, false},
		{"Create slot for random type", args{t: SlotType(10)}, nil, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := createSlot(tt.args.t)
			if (err != nil) != tt.wantErr {
				t.Errorf("createSlot() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("createSlot() = %v, want %v", got, tt.want)
			}
			if got != nil && got.GetType() != tt.args.t {
				t.Errorf("Input slot type = %v, slot created of type %v", tt.args.t, got.GetType())
			}
		})
	}
}

func TestSlotType_UnmarshalJSON(t *testing.T) {
	var slotType SlotType

	b0, _ := json.Marshal(string("DEFAULT"))
	b1, _ := json.Marshal(string("TWO_WHEELER"))
	b2, _ := json.Marshal(string("FOUR_WHEELER"))
	b3, _ := json.Marshal(string("BIG_VEHIHLE"))
	b4, _ := json.Marshal(string("TotalSlotTypes"))

	slotType.UnmarshalJSON(b0)

	type args struct {
		b []byte
	}
	tests := []struct {
		name    string
		s       *SlotType
		args    args
		wantErr bool
	}{
		{"Unmarshal to DEFAULT type", &slotType, args{b: b0}, false},
		{"Unmarshal to TWO_WHEELER type", &slotType, args{b: b1}, false},
		{"Unmarshal to FOUR_WHEELER type", &slotType, args{b: b2}, false},
		{"Unmarshal to BIG_VEHIHLE type", &slotType, args{b: b3}, false},
		{"Check error", &slotType, args{b: []byte("random")}, true},
		{"Unmarshal to invalid type", &slotType, args{b: b4}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.s.UnmarshalJSON(tt.args.b); (err != nil) != tt.wantErr {
				t.Errorf("SlotType.UnmarshalJSON() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
