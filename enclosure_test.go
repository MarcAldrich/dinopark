package dinopark

import (
	"reflect"
	"testing"
)

func TestEnclosure_ReadEnclosureState(t *testing.T) {
	type fields struct {
		contains   []*Dinosaur
		capacity   *EnclosureCapacity
		powerState bool
	}
	tests := []struct {
		name               string
		fields             fields
		wantEnclosureState *Enclosure
		wantErr            *Error
	}{
		{
			name: "capacityOf1",
			fields: fields{
				contains: []*Dinosaur{},
				capacity: &EnclosureCapacity{
					species:  "testSpecies",
					capacity: 1,
				},
				powerState: false,
			},
			wantEnclosureState: &Enclosure{
				contains: []*Dinosaur{},
				capacity: &EnclosureCapacity{
					species:  "testSpecies",
					capacity: 1,
				},
				powerState: false,
			},
			wantErr: nil,
		},
		{
			name: "capacity of 0",
			fields: fields{
				contains: []*Dinosaur{},
				capacity: &EnclosureCapacity{
					species:  "",
					capacity: 0,
				},
				powerState: false,
			},
			wantEnclosureState: &Enclosure{
				contains: []*Dinosaur{},
				capacity: &EnclosureCapacity{
					species:  "",
					capacity: 0,
				},
				powerState: false,
			},
			wantErr: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := &Enclosure{
				contains:   tt.fields.contains,
				capacity:   tt.fields.capacity,
				powerState: tt.fields.powerState,
			}
			gotEnclosureState, gotErr := e.ReadEnclosureState()
			if !reflect.DeepEqual(gotEnclosureState, tt.wantEnclosureState) {
				t.Errorf("Enclosure.ReadEnclosureState() gotEnclosureState = %v, want %v", gotEnclosureState, tt.wantEnclosureState)
			}
			if !reflect.DeepEqual(gotErr, tt.wantErr) {
				t.Errorf("Enclosure.ReadEnclosureState() gotErr = %v, want %v", gotErr, tt.wantErr)
			}
		})
	}
}
