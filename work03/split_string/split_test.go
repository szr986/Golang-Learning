package split_string

import (
	"reflect"
	"testing"
)

// func TestSplit(t *testing.T) {
// 	ret := Split("babcbef", "b")
// 	want := []string{"", "a", "c", "f"}
// 	if !reflect.DeepEqual(ret, want) {
// 		t.Errorf("got %#v, want %#v", ret, want)
// 	}
// }

// func Test2Split(t *testing.T) {
// 	ret := Split("a:b:c", ":")
// 	want := []string{"a", "b", "c"}
// 	if !reflect.DeepEqual(ret, want) {
// 		t.Errorf("got %#v, want %#v", ret, want)
// 	}
// }

func TestSplit(t *testing.T) {
	type testCase struct {
		str  string
		sep  string
		want []string
	}

	testGroup := []testCase{
		testCase{"babcbef", "b", []string{"", "a", "c", "f"}},
		testCase{"a:b:c", ":", []string{"a", "b", "c"}},
		testCase{"abcef", "bc", []string{"a", "ef"}},
	}

	for _, tc := range testGroup {
		got := Split(tc.str, tc.sep)
		if !reflect.DeepEqual(got, tc.want) {
			t.Fatalf("want:%#v got:%#v", tc.want, got)
		}
	}
}

// 基准测试
func BenchmarkSplit(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Split("a:b:c:d:e", ":")
	}
}
