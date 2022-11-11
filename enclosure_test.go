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

func TestEnclosure_SetEnclosureCapacity(t *testing.T) {
	type fields struct {
		contains   []*Dinosaur
		capacity   *EnclosureCapacity
		powerState bool
	}
	type args struct {
		newEncCap *EnclosureCapacity
	}
	tests := []struct {
		name                   string
		fields                 fields
		args                   args
		wantConfiguredCapacity *EnclosureCapacity
		wantErr                *Error
	}{
		{
			name: "from init to 1",
			fields: fields{
				contains: []*Dinosaur{},
				capacity: &EnclosureCapacity{
					//Starting state
					species:  "",
					capacity: 0,
				},
				powerState: false,
			},
			args: args{
				newEncCap: &EnclosureCapacity{
					species:  "species1",
					capacity: 1,
				},
			},
			wantConfiguredCapacity: &EnclosureCapacity{
				species:  "species1",
				capacity: 1,
			},
			wantErr: nil,
		},
		{
			name: "expect conflict: attempt reconfigure capacity on non-empty enclosure",
			fields: fields{
				contains: []*Dinosaur{{}},
				capacity: &EnclosureCapacity{
					species:  "species1",
					capacity: 1,
				},
				powerState: false,
			},
			args: args{
				newEncCap: &EnclosureCapacity{
					species:  "species2",
					capacity: 2,
				},
			},
			wantConfiguredCapacity: nil,
			wantErr:                &EncNotEmpty,
		},
		{
			name: "expect error: nil arg input",
			fields: fields{
				contains:   []*Dinosaur{},
				capacity:   &EnclosureCapacity{},
				powerState: false,
			},
			args: args{
				newEncCap: nil,
			},
			wantConfiguredCapacity: nil,
			wantErr:                &EncInvalidConfig,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := &Enclosure{
				contains:   tt.fields.contains,
				capacity:   tt.fields.capacity,
				powerState: tt.fields.powerState,
			}
			gotConfiguredCapacity, gotErr := e.SetEnclosureCapacity(tt.args.newEncCap)
			if !reflect.DeepEqual(gotConfiguredCapacity, tt.wantConfiguredCapacity) {
				t.Errorf("Enclosure.SetEnclosureCapacity() gotConfiguredCapacity = %v, want %v", gotConfiguredCapacity, tt.wantConfiguredCapacity)
			}
			if !reflect.DeepEqual(gotErr, tt.wantErr) {
				t.Errorf("Enclosure.SetEnclosureCapacity() gotErr = %v, want %v", gotErr, tt.wantErr)
			}
		})
	}
}

func TestEnclosure_CmdPwrState(t *testing.T) {
	type fields struct {
		contains   []*Dinosaur
		capacity   *EnclosureCapacity
		powerState bool
	}
	type args struct {
		commandedPwrState bool
	}
	tests := []struct {
		name                string
		fields              fields
		args                args
		wantCurrentPwrState *bool
		wantErr             *Error
	}{
		{
			name: "expect success: init to power on",
			fields: fields{
				contains:   []*Dinosaur{},
				capacity:   &EnclosureCapacity{},
				powerState: false,
			},
			args: args{
				commandedPwrState: true,
			},
			wantCurrentPwrState: ptrToBool(true),
			wantErr:             nil,
		},
		{
			name: "expect error: dino in enc to start-> attempt pwr down",
			fields: fields{
				contains: []*Dinosaur{{}},
				capacity: &EnclosureCapacity{
					species:  "species1",
					capacity: 1,
				},
				powerState: true,
			},
			args: args{
				commandedPwrState: false,
			},
			wantCurrentPwrState: nil,
			wantErr:             &EncNotEmpty,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := &Enclosure{
				contains:   tt.fields.contains,
				capacity:   tt.fields.capacity,
				powerState: tt.fields.powerState,
			}
			gotCurrentPwrState, gotErr := e.CmdPwrState(tt.args.commandedPwrState)
			//NOTE: protecting test from dereferencing a nil ptr
			if gotCurrentPwrState != nil && tt.wantCurrentPwrState != nil &&
				*gotCurrentPwrState != *tt.wantCurrentPwrState {
				t.Errorf("Enclosure.CmdPwrState() gotCurrentPwrState = %v, want %v", gotCurrentPwrState, tt.wantCurrentPwrState)
			}
			if !reflect.DeepEqual(gotErr, tt.wantErr) {
				t.Errorf("Enclosure.CmdPwrState() gotErr = %v, want %v", gotErr, tt.wantErr)
			}
		})
	}
}

func ptrToBool(val bool) *bool {
	return &val
}
