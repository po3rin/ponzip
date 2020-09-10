package main

import "testing"

func TestBWT(t *testing.T) {
	tests := []struct {
		name string
		t    string
		want string
	}{
		{
			name: "normal",
			t:    "abracadabra",
			want: "ard$rcaaaabb",
		},
		{
			name: "simple",
			t:    "acaacg",
			want: "gc$aaac",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := BWT(tt.t)
			if got != tt.want {
				t.Errorf("\nwant: %+v\ngot : %+v\n", tt.want, got)
			}
		})
	}
}

func TestBWTInverse(t *testing.T) {
	tests := []struct {
		name string
		t    string
		want string
	}{
		{
			name: "normal",
			t:    "ard$rcaaaabb",
			want: "abracadabra",
		},
		{
			name: "simple",
			t:    "gc$aaac",
			want: "acaacg",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := BWTInverse(tt.t)
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
