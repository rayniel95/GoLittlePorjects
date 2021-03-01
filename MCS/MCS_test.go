package mcs

import "testing"

func Test_toMCS(t *testing.T) {
	type args struct {
		small string
		large string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "test1", args: args{small: "cambiamos", large: "comienzosas"},
			want: 7,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := toMCS(tt.args.small, tt.args.large); got != tt.want {
				t.Errorf("toMCS() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_mcs(t *testing.T) {
	type args struct {
		small       []byte
		large       []byte
		bestCount   *int
		actualCount int
		actualIndex int
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mcs(tt.args.small, tt.args.large, tt.args.bestCount, tt.args.actualCount, tt.args.actualIndex)
		})
	}
}
