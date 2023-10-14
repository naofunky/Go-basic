package main

import "testing"

func TestSum(t *testing.T) {
	sl := []int{}
	sl = append(sl, 1, 2, 3, 4, 5)
	if sl[0] != 1 {
		t.Errorf("sl[0] is not 1")
	}
	if sl[1] != 2 {
		t.Errorf("sl[1] is not 2")
	}
	if sl[2] != 3 {
		t.Errorf("sl[2] is not 3")
	}
	if sl[3] != 4 {
		t.Errorf("sl[3] is not 4")
	}
	if sl[4] != 5 {
		t.Errorf("sl[4] is not 5")
	}
}
