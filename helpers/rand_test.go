package helpers

import (
	"fmt"
	"testing"
)

func TestRandomStr(t *testing.T) {
	val := RandomStr(10)
	fmt.Print(val)
	if len(val) != 10 {
		t.Errorf("length must be 10, but got %d", len(val))
	}
}
