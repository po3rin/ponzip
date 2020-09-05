package main

import "testing"

func TestGetChar(t *testing.T) {
	tests := []struct {
		name  string
		t     string
		index int
		want  string
	}{
		{
			name:  "simple",
			t:     "abracadabra",
			index: 3,
			want:  "a",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := GetChar(tt.t, tt.index)
			if got != tt.want {
				t.Errorf("\nwant: %+v\ngot : %+v\n", tt.want, got)
			}
		})
	}
}

func TestInsert(t *testing.T) {
	tests := []struct {
		name  string
		t     string
		c     string
		index int
		want  string
	}{
		{
			name:  "zero",
			t:     "aaaaa",
			index: 0,
			c:     "b",
			want:  "baaaaa",
		},
		{
			name:  "simple1",
			t:     "aaaaa",
			index: 3,
			c:     "b",
			want:  "aaabaa",
		},
		{
			name:  "simple2",
			t:     "aaaaa",
			index: 5,
			c:     "b",
			want:  "aaaaab",
		},
		{
			name:  "simple3",
			t:     "aaaaa",
			index: 2,
			c:     "b",
			want:  "aabaaa",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Insert(tt.t, tt.c, tt.index)
			if got != tt.want {
				t.Errorf("\nwant: %+v\ngot : %+v\n", tt.want, got)
			}
		})
	}
}
func TestRank(t *testing.T) {
	tests := []struct {
		name  string
		t     string
		c     string
		index int
		want  int
	}{
		{
			name:  "simple1",
			t:     "abracadabra",
			index: 3,
			c:     "a",
			want:  1,
		},
		{
			name:  "simple2",
			t:     "abracadabra",
			index: 5,
			c:     "c",
			want:  1,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Rank(tt.t, tt.c, tt.index)
			if got != tt.want {
				t.Errorf("\nwant: %+v\ngot : %+v\n", tt.want, got)
			}
		})
	}
}

func TestRankLessThan(t *testing.T) {
	tests := []struct {
		name  string
		t     string
		c     string
		index int
		want  int
	}{
		{
			name:  "simple1",
			t:     "abracadabra$",
			index: 5,
			c:     "c",
			want:  3,
		},
		{
			name:  "simple2",
			t:     "abracadabra$",
			index: 5,
			c:     "r",
			want:  4,
		},
		{
			name:  "full",
			t:     "abracadabra$",
			index: 12,
			c:     "a",
			want:  1,
		},
		{
			name:  "full2",
			t:     "abracadabra$",
			index: 12,
			c:     "b",
			want:  6,
		},
		{
			name:  "full3",
			t:     "abracadabra$",
			index: 12,
			c:     "c",
			want:  8,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := RankLessThan(tt.t, tt.c, tt.index)
			if got != tt.want {
				t.Errorf("\nwant: %+v\ngot : %+v\n", tt.want, got)
			}
		})
	}
}
