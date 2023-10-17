package main

import "testing"

func TestSum(t *testing.T) {

	t.Run("スライスの初期化", func(t *testing.T) {
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
	})

	t.Run("nilスライスと空スライスの検証", func(t *testing.T) {
		var nilSlice []int
		emptySlice := []int{}

		if nilSlice != nil {
			t.Errorf("nilSlice is nil")
		}
		if emptySlice == nil {
			t.Errorf("emptySlice is not nil")
		}
		if len(nilSlice) != 0 {
			t.Errorf("nilSlice length is not 0")
		}
		if len(emptySlice) != 0 {
			t.Errorf("emptySlice length is not 0")
		}
		if cap(nilSlice) != 0 {
			t.Errorf("nilSlice capacity is not 0")
		}
		if cap(emptySlice) != 0 {
			t.Errorf("emptySlice capacity is not 0")
		}
	})
}
