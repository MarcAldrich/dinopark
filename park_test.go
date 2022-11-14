package dinopark

import (
	"reflect"
	"sync"
	"testing"
)

func TestPark_ListPlaces(t *testing.T) {
	type fields struct {
		Dinos      []Dinosaur
		Enclosures []Enclosure
		Places     Places
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
				Dinos:      []Dinosaur{},
				Enclosures: []Enclosure{},
				Places:     Places{},
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
				Dinos:      []Dinosaur{},
				Enclosures: []Enclosure{},
				Places: Places{
					Places: []Place{{
						//This place should NOT match filter
						Name:     "Not a match",
						Location: "lab",
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
				Dinos:      []Dinosaur{},
				Enclosures: []Enclosure{},
				Places: Places{
					Places: []Place{{
						Name:     "Place1",
						Location: "Loc1",
						Kind:     ENCLOSURE,
					},
						{
							Name:     "Place2",
							Location: "Loc2",
							Kind:     LAB,
						}},
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
				Dinos:      []Dinosaur{},
				Enclosures: []Enclosure{},
				Places: Places{
					Places: []Place{{
						Name:     "Place1",
						Location: "Loc1",
						Kind:     LAB,
					},
						{
							Name:     "Place2",
							Location: "Loc2",
							Kind:     ENCLOSURE,
						}},
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
			p := Park{
				Dinos:      tt.fields.Dinos,
				Enclosures: tt.fields.Enclosures,
				Places:     tt.fields.Places,
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
						if v == plToValidate {
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

func TestPark_RegisterPlace(t *testing.T) {
	type fields struct {
		Dinos      []Dinosaur
		Enclosures []Enclosure
		Places     *Places
	}
	type args struct {
		plcToReg *Place
	}
	tests := []struct {
		name         string
		fields       fields
		args         args
		wantPlcAdded *Place
		wantErr      *Error
	}{
		{
			name: "add 1 to none",
			fields: fields{
				Dinos:      []Dinosaur{},
				Enclosures: []Enclosure{},
				Places: &Places{
					Places: []Place{},
					mu:     &sync.Mutex{},
				},
			},
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
				Dinos:      []Dinosaur{},
				Enclosures: []Enclosure{},
				Places: &Places{
					Places: []Place{Place{
						ID:       [16]byte{},
						Name:     "Place1",
						Location: "Loc1",
						Kind:     LAB,
					}},
					mu: &sync.Mutex{},
				},
			},
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
			p := &Park{
				Dinos:      tt.fields.Dinos,
				Enclosures: tt.fields.Enclosures,
				Places:     *tt.fields.Places,
			}
			gotPlcAdded, gotErr := p.RegisterPlace(tt.args.plcToReg)
			if !reflect.DeepEqual(gotPlcAdded, tt.wantPlcAdded) {
				t.Errorf("Park.RegisterPlace() gotPlcAdded = %v, want %v", gotPlcAdded, tt.wantPlcAdded)
			}
			if !reflect.DeepEqual(gotErr, tt.wantErr) {
				t.Errorf("Park.RegisterPlace() gotErr = %v, want %v", gotErr, tt.wantErr)
			}

			//COMPARE BACKING SLICE FOR CORRECT LENGTH
			expectedLen := len(tt.fields.Places.Places) + 1 //+1 because arg takes a single place, not a list.
			if expectedLen != len(p.Places.Places) {
				t.Errorf("Place not added: expected number of place entries to be %d; instead had %d entries", expectedLen, len(p.Places.Places))
			}
		})
	}
}
