package main

import "testing"

func TestBwt(t *testing.T) {
	tests := []struct {
		name string
		t    string
		want string
	}{
		{
			name: "simple",
			t:    "abracadabra",
			want: "ard$rcaaaabb",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Bwt(tt.t)
			if got != tt.want {
				t.Errorf("\nwant: %+v\ngot : %+v\n", tt.want, got)
			}
		})
	}
}

// func TestBwtIncrement(t *testing.T) {
// 	tests := []struct {
// 		name string
// 		t    string
// 		want string
// 	}{
// 		{
// 			name: "simple",
// 			t:    "abracadabra",
// 			want: "ard$rcaaaabb",
// 		},
// 	}

// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			got := BwtIncrement(tt.t)
// 			if got != tt.want {
// 				t.Errorf("\nwant: %+v\ngot : %+v\n", tt.want, got)
// 			}
// 		})
// 	}
// }
