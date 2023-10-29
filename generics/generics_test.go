package main

import "testing"

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
