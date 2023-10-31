package main

import (
	"reflect"
	"testing"
)

// generics/generics.goのコードをテストする
func TestGenerics(t *testing.T) {
	t.Run("int型の加算", func(t *testing.T) {
		if add(1, 2) != 3 {
			t.Errorf("add(1, 2) is not 3")
		}
	})
	t.Run("int16型の加算", func(t *testing.T) {
		if add(int16(1), int16(2)) != int16(3) {
			t.Errorf("add(int16(1), int16(2)) is not int16(3)")
		}
	})
	t.Run("float32型の加算", func(t *testing.T) {
		if add(float32(1.0), float32(2.0)) != float32(3.0) {
			t.Errorf("add(float32(1.0), float32(2.0)) is not float32(3.0)")
		}
	})
	t.Run("float64型の加算", func(t *testing.T) {
		if add(float64(1.0), float64(2.0)) != float64(3.0) {
			t.Errorf("add(float64(1.0), float64(2.0)) is not float64(3.0)")
		}
	})
	t.Run("string型の加算", func(t *testing.T) {
		if add("1", "2") != "12" {
			t.Errorf("add(\"1\", \"2\") is not \"12\"")
		}
	})
}

func Test_reverse(t *testing.T) {
	type args struct {
		s []string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		// TODO: Add test cases
		{
			name: "reverse",
			args: args{
				s: []string{"a", "b", "c"},
			},
			want: []string{"c", "b", "a"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := reverse(tt.args.s); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("reverse() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_sumValues(t *testing.T) {
	type args struct {
		m map[int]string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
		{
			name: "sumValues",
			args: args{
				m: map[int]string{
					1: "あ",
					2: "い",
				},
			},
			want: "あい",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := sumValues(tt.args.m); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("sumValues() = %v, want %v", got, tt.want)
			}
		})
	}
}
