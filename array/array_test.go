package main

import "testing"

// ayyay.go内に記述されているmain関数に配列の内容をテスト
func TestSum(t *testing.T) {
	arr := [5]int{1, 2, 3, 4, 5}
	if arr[0] != 1 {
		t.Errorf("arr[0] is not 1")
	}
	if arr[1] != 2 {
		t.Errorf("arr[1] is not 2")
	}
	if arr[2] != 3 {
		t.Errorf("arr[2] is not 3")
	}
	if arr[3] != 4 {
		t.Errorf("arr[3] is not 4")
	}
	if arr[4] != 5 {
		t.Errorf("arr[4] is not 5")
	}
}

func TestArrayLength(t *testing.T) {
	arr := [5]int{1, 2, 3, 4, 5}
	if len(arr) != 5 {
		t.Errorf("arr length is not 5")
	}
}
