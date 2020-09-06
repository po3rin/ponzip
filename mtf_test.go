package main

// func TestMTF(t *testing.T) {
// 	tests := []struct {
// 		name        string
// 		t           string
// 		want        string
// 		wantSymbols string
// 	}{
// 		{
// 			name:        "simple",
// 			t:           "ababacaca",
// 			want:        "ab222c222",
// 			wantSymbols: "acb",
// 		},
// 	}

// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			got, symbols := MTF(tt.t)
// 			if string(got) != tt.want {
// 				t.Errorf("\nwant: %+v\ngot : %+v\n", tt.want, string(got))
// 			}

// 			if symbols == tt.wantSymbols {
// 				t.Errorf("\nwant: %+v\ngot : %+v\n", tt.wantSymbols, string(symbols))
// 			}
// 		})
// 	}
// }
