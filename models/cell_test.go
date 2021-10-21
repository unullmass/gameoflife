package models

import (
	"reflect"
	"testing"
)

func TestCell_String(t *testing.T) {
	tests := []struct {
		name string
		c    Cell
		want string
	}{
		{
			name: "Alive Cell",
			c: Cell{
				alive: true,
			},
			want: "A",
		},
		{
			name: "Dead Cell",
			c: Cell{
				alive: false,
			},
			want: "D",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.c.String(); got != tt.want {
				t.Errorf("Cell.String() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCell_IsAlive(t *testing.T) {
	tests := []struct {
		name string
		c    Cell
		want bool
	}{
		{
			name: "Alive Cell",
			c: Cell{
				alive: true,
			},
			want: true,
		},
		{
			name: "Dead Cell",
			c: Cell{
				alive: false,
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.c.IsAlive(); got != tt.want {
				t.Errorf("Cell.IsAlive() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewCell(t *testing.T) {
	type args struct {
		a  int
		nc int
	}
	tests := []struct {
		name string
		args args
		want Cell
	}{
		{
			name: "Alive Cell",
			args: args{
				a: 1,
			},
			want: Cell{
				alive:         true,
				liveNeighbors: 0,
			},
		},
		{
			name: "Dead Cell",
			args: args{
				a: 0,
			},
			want: Cell{
				alive:         false,
				liveNeighbors: 0,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewCell(tt.args.a, tt.args.nc); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewCell() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCell_CountLiveNeighbors(t *testing.T) {
	type args struct {
		neighbors []Cell
	}
	tests := []struct {
		name    string
		c       *Cell
		args    args
		wantErr bool
	}{
		{
			name: "Valid NeighborSet",
			c: &Cell{
				alive:         false,
				liveNeighbors: 0,
			},
			args: args{
				neighbors: []Cell{
					NewCell(0, 0),
					NewCell(0, 1),
					NewCell(0, 1),
					NewCell(0, 1),
					NewCell(0, 0),
					NewCell(0, 0),
					NewCell(0, 1),
					NewCell(0, 0)},
			},
			wantErr: false,
		},
		{
			name: "Invalid NeighborSet",
			c: &Cell{
				alive:         true,
				liveNeighbors: 0,
			},
			args: args{
				neighbors: []Cell{
					NewCell(0, 0),
					NewCell(0, 1),
					NewCell(0, 0)},
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.c.CountLiveNeighbors(tt.args.neighbors); (err != nil) != tt.wantErr {
				t.Errorf("Cell.CountLiveNeighbors() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
