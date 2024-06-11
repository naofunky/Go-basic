package main

import (
	"reflect"
	"testing"
)

func operateSlice() []int {
	s := make([]int, 0, 5)
	s = append(s, 1, 2, 3)
	return s
}

func Test_operateSlice(t *testing.T) {
	tests := []struct {
		name string
		want []int
	}{
		{
			name: "Test case 1: Check operateSlice function",
			want: []int{1, 2, 3},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := operateSlice(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("operateSlice() = %v, want %v", got, tt.want)
			}
		})
	}
}
