package lcs

import "testing"

func Test_toLCS(t *testing.T) {
	type args struct {
		small string
		large string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "test1", args: args{small: "cambiamos", large: "comienzos"},
			want: "cmios",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := toLCS(tt.args.small, tt.args.large); got != tt.want {
				t.Errorf("toLCS() = %v, want %v", got, tt.want)
			}
		})
	}
}

// func Test_lcs(t *testing.T) {
// 	type args struct {
// 		small            []byte
// 		large            []byte
// 		bestString       string
// 		actualString     []byte
// 		actualIndex      int
// 		actualLargeIndex int
// 		rec              int
// 	}
// 	tests := []struct {
// 		name string
// 		args args
// 	}{
// 		// TODO: Add test cases.
// 	}
// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			lcs(tt.args.small, tt.args.large, tt.args.bestString, tt.args.actualString, tt.args.actualIndex, tt.args.actualLargeIndex, tt.args.rec)
// 		})
// 	}
// }
