package bigsubpalindrome

import (
	"reflect"
	"testing"
)

func Test_isPal(t *testing.T) {
	type args struct {
		array []byte
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			"test1", args{
				[]byte{9, 10, 2, 10, 9},
			},
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isPal(tt.args.array); got != tt.want {
				t.Errorf("isPal() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_bigSubPal(t *testing.T) {
	type args struct {
		array []byte
	}
	tests := []struct {
		name string
		args args
		want []byte
	}{
		{
			"test1", args{
				[]byte{4, 2, 10, 9, 9, 10, 2, 4},
			},
			[]byte{4, 2, 10, 9, 9, 10, 2, 4},
		},
		{
			"test2", args{
				[]byte{10, 9, 9, 10, 2, 10, 9, 7},
			},
			[]byte{9, 10, 2, 10, 9},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := bigSubPal(tt.args.array); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("igSubPal() = %v, want %v", got, tt.want)
			}
		})
	}
}
