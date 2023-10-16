package main

import "testing"

func TestMain(t *testing.T) {

	t.Run("初期構文テスト", func(t *testing.T) {
		m := map[string]string{"age": "18", "name": "Bob", "sex": "woman"}

		if m["name"] != "Bob" {
			t.Errorf("m[name] is not Bob")
		}
		if m["age"] != "18" {
			t.Errorf("m[age] is not 18")
		}
		if m["sex"] != "woman" {
			t.Errorf("m[sex] is not man")
		}
		if len(m) != 3 {
			t.Errorf("m length is not 3")
		}
	})
}
