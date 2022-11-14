package dinopark

import (
	"reflect"
	"sync"
	"testing"

	"github.com/google/uuid"
)

func TestPlaces_ListPlaces(t *testing.T) {
	type fields struct {
		Places *Places
	}
	type args struct {
		filter *PlaceFilter
	}
	tests := []struct {
		name       string
		fields     fields
		args       args
		wantPlaces []Place
		wantErr    *Error
	}{
		{
			name: "0-length return: no filter",
			fields: fields{
				Places: NewPlaces(),
			},
			args: args{
				filter: nil,
			},
			wantPlaces: []Place{},
			wantErr:    nil,
		},
		{
			name: "0-length return: place-kind filter",
			fields: fields{
				Places: &Places{
					Places: map[uuid.UUID]*Place{[16]byte{'a', 'a'}: &Place{
						ID:       [16]byte{'a', 'a'},
						Name:     "Entry doesn't match filter",
						Location: "Loc1",
						Kind:     ENCLOSURE,
					}},
					mu: &sync.Mutex{},
				},
			},
			args: args{
				filter: &PlaceFilter{
					//Filter does not match entry->should return 0 places
					ByKind: NewPlaceKind(LAB),
				},
			},
			wantPlaces: []Place{},
			wantErr:    nil,
		},
		{
			name: "1-length return: filter by place-kind",
			fields: fields{
				Places: &Places{
					Places: map[uuid.UUID]*Place{
						[16]byte{'a', 'a'}: &Place{
							ID:       [16]byte{'a', 'a'},
							Name:     "Place1",
							Location: "Loc1",
							Kind:     ENCLOSURE,
						},
						[16]byte{'b', 'b'}: &Place{
							ID:       [16]byte{'b', 'b'},
							Name:     "Place2",
							Location: "Loc2",
							Kind:     LAB,
						},
					},
					mu: &sync.Mutex{},
				},
			},
			args: args{
				filter: &PlaceFilter{
					ByKind: NewPlaceKind(LAB),
				},
			},
			wantPlaces: []Place{{
				Name:     "Place2",
				Location: "Loc2",
				Kind:     LAB,
			}},
			wantErr: nil,
		},
		{
			name: "All-return: no-filter",
			fields: fields{
				Places: &Places{
					Places: map[uuid.UUID]*Place{
						[16]byte{'a', 'a'}: &Place{
							ID:       [16]byte{'a', 'a'},
							Name:     "Place1",
							Location: "Loc1",
							Kind:     ENCLOSURE,
						},
						[16]byte{'b', 'b'}: &Place{
							ID:       [16]byte{'b', 'b'},
							Name:     "Place2",
							Location: "Loc2",
							Kind:     LAB,
						},
					},
					mu: &sync.Mutex{},
				},
			},
			args: args{
				filter: nil,
			},
			wantPlaces: []Place{
				{
					Name:     "Place1",
					Location: "Loc1",
					Kind:     LAB,
				},
				{
					Name:     "Place2",
					Location: "Loc2",
					Kind:     ENCLOSURE,
				},
			},
			wantErr: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := Places{
				Places: tt.fields.Places.Places,
				mu:     &sync.Mutex{},
			}
			gotPlaces, gotErr := p.ListPlaces(tt.args.filter)
			if gotPlaces != nil {
				//Length must match
				if len(gotPlaces) != len(tt.wantPlaces) {
					t.Errorf("Park.ListPlaces() gotPlaces = %v, want %v", gotPlaces, tt.wantPlaces)
				}
				//Compare entries
				for _, plToValidate := range tt.wantPlaces {
					found := false
					//Check got places for the place to validate
					for _, v := range gotPlaces {
						if v.Name == plToValidate.Name {
							found = true
							break
						}
					}
					if !found {
						t.Errorf("Park.ListPlaces() gotPlaces = %v, want %v", gotPlaces, tt.wantPlaces)
					}
				}
			}

			if !reflect.DeepEqual(gotErr, tt.wantErr) {
				t.Errorf("Park.ListPlaces() gotErr = %v, want %v", gotErr, tt.wantErr)
			}
		})
	}
}

func TestPlaces_RegisterPlace(t *testing.T) {
	type fields struct {
		Places *Places
	}
	type args struct {
		plcToReg *Place
	}
	tests := []struct {
		name               string
		fields             fields
		expectedPlaceCount int
		args               args
		wantPlcAdded       *Place
		wantErr            *Error
	}{
		{
			name: "add 1 to none",
			fields: fields{
				Places: NewPlaces(),
			},
			expectedPlaceCount: 1,
			args: args{
				plcToReg: &Place{
					ID:       [16]byte{},
					Name:     "Place1",
					Location: "Loc1",
					Kind:     LAB,
				},
			},
			wantPlcAdded: &Place{
				ID:       [16]byte{},
				Name:     "Place1",
				Location: "Loc1",
				Kind:     LAB,
			},
			wantErr: nil,
		},
		{
			name: "add 1 to 1 -> expect length of 2",
			fields: fields{
				Places: &Places{
					Places: map[uuid.UUID]*Place{
						[16]byte{'a', 'a'}: &Place{
							ID:       [16]byte{'a', 'a'},
							Name:     "Place1",
							Location: "Loc1",
							Kind:     ENCLOSURE,
						}},
					mu: &sync.Mutex{},
				},
			},
			expectedPlaceCount: 2,
			args: args{
				plcToReg: &Place{
					ID:       [16]byte{},
					Name:     "Place2",
					Location: "Loc2",
					Kind:     LAB,
				},
			},
			wantPlcAdded: &Place{
				ID:       [16]byte{},
				Name:     "Place2",
				Location: "Loc2",
				Kind:     LAB,
			},
			wantErr: nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &Places{
				Places: tt.fields.Places.Places,
				mu:     &sync.Mutex{},
			}
			gotPlcAdded, gotErr := p.RegisterPlace(tt.args.plcToReg)
			if !reflect.DeepEqual(gotPlcAdded, tt.wantPlcAdded) {
				t.Errorf("Park.RegisterPlace() gotPlcAdded = %v, want %v", gotPlcAdded, tt.wantPlcAdded)
			}
			if !reflect.DeepEqual(gotErr, tt.wantErr) {
				t.Errorf("Park.RegisterPlace() gotErr = %v, want %v", gotErr, tt.wantErr)
			}

			//COMPARE BACKING SLICE FOR CORRECT LENGTH
			if tt.expectedPlaceCount != len(p.Places) {
				t.Errorf("Place not added: expected number of place entries to be %d; instead had %d entries", tt.expectedPlaceCount, len(p.Places))
			}
		})
	}
}
