package prefixtree

import (
	"testing"
)

func Test_main(t *testing.T) {

	want := 2
	got := doubleEntreOrden2((*root).left, (*root).right)

	if got != want {
		t.Errorf("Hello() = %q, want %q", got, want)
	}
}
