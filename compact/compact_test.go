package compact

import (
	"testing"
)

func Test_compact(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{name: "test1", args: args{s: "abcabcabc"},
			want: "(abc)3"},
		{name: "test2", args: args{s: "axyxyxyxyb"},
			want: "a((xy)2)2b"},
		{name: "test3", args: args{s: "xxabababxxxxabababxx"},
			want: "(xx(ab)3xx)2"},
		{name: "test4", args: args{s: "abcdefghijkl"},
			want: "abcdefghijkl"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := compact(tt.args.s); got != tt.want {
				t.Errorf("compact() = %v, want %v", got, tt.want)
			}
		})
	}
}

// func Test_compactA(t *testing.T) {
// 	type args struct {
// 		s          string
// 		startIndex int
// 		endIndex   int
// 	}
// 	tests := []struct {
// 		name  string
// 		args  args
// 		want  []byte
// 		want1 int
// 		want2 int
// 	}{
// 		// TODO: Add test cases.
// 	}
// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			got, got1, got2 := compactA(tt.args.s, tt.args.startIndex, tt.args.endIndex)
// 			if !reflect.DeepEqual(got, tt.want) {
// 				t.Errorf("compactA() got = %v, want %v", got, tt.want)
// 			}
// 			if got1 != tt.want1 {
// 				t.Errorf("compactA() got1 = %v, want %v", got1, tt.want1)
// 			}
// 			if got2 != tt.want2 {
// 				t.Errorf("compactA() got2 = %v, want %v", got2, tt.want2)
// 			}
// 		})
// 	}
// }

// func Test_searchPattern(t *testing.T) {
// 	type args struct {
// 		s          string
// 		startIndex int
// 		endIndex   int
// 	}
// 	tests := []struct {
// 		name  string
// 		args  args
// 		want  int
// 		want1 int
// 		want2 int
// 	}{
// 		// TODO: Add test cases.
// 	}
// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			got, got1, got2 := searchPattern(tt.args.s, tt.args.startIndex, tt.args.endIndex)
// 			if got != tt.want {
// 				t.Errorf("searchPattern() got = %v, want %v", got, tt.want)
// 			}
// 			if got1 != tt.want1 {
// 				t.Errorf("searchPattern() got1 = %v, want %v", got1, tt.want1)
// 			}
// 			if got2 != tt.want2 {
// 				t.Errorf("searchPattern() got2 = %v, want %v", got2, tt.want2)
// 			}
// 		})
// 	}
// }

// func Test_checkPatternTimes(t *testing.T) {
// 	type args struct {
// 		s             string
// 		startIndex    int
// 		endIndex      int
// 		startIndexPat int
// 		endIndexPat   int
// 	}
// 	tests := []struct {
// 		name string
// 		args args
// 		want int
// 	}{
// 		// TODO: Add test cases.
// 	}
// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			if got := checkPatternTimes(tt.args.s, tt.args.startIndex, tt.args.endIndex, tt.args.startIndexPat, tt.args.endIndexPat); got != tt.want {
// 				t.Errorf("checkPatternTimes() = %v, want %v", got, tt.want)
// 			}
// 		})
// 	}
// }
