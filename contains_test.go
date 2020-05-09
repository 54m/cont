package cont

import (
	"testing"
	"time"
)

func TestContains(t *testing.T) {
	t.Run("string", func(tr *testing.T) {
		base := "1234567890"

		if !Contains(base, "56") {
			tr.Fatal("missing logic")
		}

		if Contains(base, "98") {
			tr.Fatal("missing logic")
		}
	})

	t.Run("[]string", func(tr *testing.T) {
		base := []string{"1", "2", "3"}

		if !Contains(base, "2") {
			tr.Fatal("missing logic")
		}

		if Contains(base, "4") {
			tr.Fatal("missing logic")
		}

		if Contains(base, 1) {
			tr.Fatal("missing logic")
		}
	})

	t.Run("[]int", func(tr *testing.T) {
		base := []int{1, 2, 3}

		if !Contains(base, 1) {
			tr.Fatal("missing logic")
		}

		if Contains(base, "2") {
			tr.Fatal("missing logic")
		}

		if Contains(base, 4) {
			tr.Fatal("missing logic")
		}
	})

	t.Run("map", func(tr *testing.T) {
		now := time.Now()
		base := map[string]interface{}{"a": "1", "b": "2", "c": 3, "d": now}

		if !Contains(base, "1") {
			tr.Fatal("missing logic")
		}

		if !Contains(base, "a") {
			tr.Fatal("missing logic")
		}

		if !Contains(base, 3) {
			tr.Fatal("missing logic")
		}

		if !Contains(base, now) {
			tr.Fatal("missing logic")
		}

		if Contains(base, "e") {
			tr.Fatal("missing logic")
		}
	})
}
