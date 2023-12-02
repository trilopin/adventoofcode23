package main

import (
	"strings"
	"testing"
)

var input = strings.NewReader(`Game 1: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green
Game 2: 1 blue, 2 green; 3 green, 4 blue, 1 red; 1 green, 1 blue
Game 3: 8 green, 6 blue, 20 red; 5 blue, 4 red, 13 green; 5 green, 1 red
Game 4: 1 green, 3 red, 6 blue; 3 green, 6 red; 3 green, 15 blue, 14 red
Game 5: 6 red, 1 blue, 3 green; 2 blue, 1 red, 2 green
`)

func Test_sumValidGames(t *testing.T) {
	got := sumValidGames(input)
	want := 8
	if got != want {
		t.Errorf("got %d, want %d", got, want)
	}
}
func Test_sumMinimumBlocks(t *testing.T) {
	got := sumMinimumBlocks(input)
	want := 2286
	if got != want {
		t.Errorf("got %d, want %d", got, want)
	}
}

func Test_NewGame(t *testing.T) {
	testCases := []struct {
		name string
		in   string
		want *Game
	}{
		{
			name: "Game 1",
			in:   "Game 1: 3 blue, 4 red",
			want: &Game{
				ID: 1,
				Reveals: []Reveal{
					{Blue: 3, Red: 4, Green: 0},
				},
			},
		},
		{
			name: "Game 2",
			in:   "Game 2: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green",
			want: &Game{
				ID: 2,
				Reveals: []Reveal{
					{Blue: 3, Red: 4, Green: 0},
					{Blue: 6, Red: 1, Green: 2},
					{Blue: 0, Red: 0, Green: 2},
				},
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(tt *testing.T) {
			got, err := NewGame(tc.in)
			if err != nil {
				tt.Errorf("unexpected error on NewGame %v", err)
			}
			if got.ID != tc.want.ID {
				tt.Errorf("got %d, want %d", got.ID, tc.want.ID)
			}
			if len(got.Reveals) != len(tc.want.Reveals) {
				tt.Errorf("got %d reveals, want %d", len(got.Reveals), len(tc.want.Reveals))
			}
			for i, reveal := range got.Reveals {
				if reveal.Blue != tc.want.Reveals[i].Blue {
					tt.Errorf("got %d blue, want %d", reveal.Blue, tc.want.Reveals[i].Blue)
				}
				if reveal.Red != tc.want.Reveals[i].Red {
					tt.Errorf("got %d red, want %d", reveal.Red, tc.want.Reveals[i].Red)
				}
				if reveal.Green != tc.want.Reveals[i].Green {
					tt.Errorf("got %d green, want %d", reveal.Green, tc.want.Reveals[i].Green)
				}
			}
		})
	}
}
