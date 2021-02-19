package compact

import (
	"reflect"
	"testing"
)

func Test_decompact(t *testing.T) {
	type args struct {
		s []byte
	}
	tests := []struct {
		name string
		args args
		want []byte
	}{
		{name: "test1", args: args{s: []byte("(abc)3")},
			want: []byte("abcabcabc")},
		{name: "test2", args: args{s: []byte("a(xy)4b")},
			want: []byte("axyxyxyxyb")},
		{name: "test3", args: args{s: []byte("(xx(ab)5xx)2")},
			want: []byte("xxabababababxxxxabababababxx")},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := decompact(tt.args.s); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("decompact() = %v, want %v", string(got), tt.want)
			}
		})
	}
}
