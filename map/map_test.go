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

	t.Run("重複処理のテスト", func(t *testing.T) {
		duplicatedArray := []int{1, 2, 3, 3, 4, 5}
		notDuplicatedArray := []int{}

		myMap := map[int]struct{}{}

		for _, v := range duplicatedArray {
			_, loopValue := myMap[v]
			if !loopValue {
				myMap[v] = struct{}{}
				notDuplicatedArray = append(notDuplicatedArray, v)
			}
		}

		if len(notDuplicatedArray) != 5 {
			t.Errorf("notDuplicatedArray length is not 5")
		}
	})
}
