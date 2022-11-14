package dinopark

import (
	"reflect"
	"testing"
)

func TestNewDinosaur(t *testing.T) {
	type args struct {
		name    string
		species DinoSpecies
	}
	tests := []struct {
		name     string
		args     args
		wantDino *Dinosaur
		wantErr  *Error
	}{
		{
			name: "Create new dino",
			args: args{
				name:    "Blue",
				species: Velociraptor,
			},
			wantDino: &Dinosaur{
				Name:    "Blue",
				Species: Velociraptor,
			},
			wantErr: nil,
		},
		{
			name: "expect err: missing name",
			args: args{
				name:    "",
				species: Velociraptor,
			},
			wantDino: nil,
			wantErr:  &MissingArg,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotDino, gotErr := NewDinosaur(tt.args.name, tt.args.species)

			//Custom test: Not using deepequal because uuid is random->Test all other static fields
			if gotDino != nil &&
				gotDino.Name != tt.wantDino.Name &&
				gotDino.Species != tt.wantDino.Species &&
				len(gotDino.ID) == 0 {
				t.Errorf("NewDinosaur() gotDino = %v, want %v", gotDino, tt.wantDino)
			}
			if !reflect.DeepEqual(gotErr, tt.wantErr) {
				t.Errorf("NewDinosaur() gotErr = %v, want %v", gotErr, tt.wantErr)
			}
		})
	}
}
