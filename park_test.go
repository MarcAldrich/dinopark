package dinopark

import (
	"reflect"
	"testing"
)

func TestPark_ListPlaces(t *testing.T) {
	type fields struct {
		Dinos      []Dinosaur
		Enclosures []Enclosure
		Places     []Place
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
				Places:     []Place{},
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
				Places: []Place{{
					//This place should NOT match filter
					Name:     "Not a match",
					Location: "lab",
					Kind:     ENCLOSURE,
				}},
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
		//TODO: Next 1-length return
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
