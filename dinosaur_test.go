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
			if !reflect.DeepEqual(gotDino, tt.wantDino) {
				t.Errorf("NewDinosaur() gotDino = %v, want %v", gotDino, tt.wantDino)
			}
			if !reflect.DeepEqual(gotErr, tt.wantErr) {
				t.Errorf("NewDinosaur() gotErr = %v, want %v", gotErr, tt.wantErr)
			}
		})
	}
}
