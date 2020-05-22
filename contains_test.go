package cont_test

import (
	"testing"
	"time"

	"github.com/54m/cont"
)

func TestContains(t *testing.T) {
	t.Run("string", func(tr *testing.T) {
		base := "1234567890"

		if !cont.Contains(base, "56") {
			tr.Fatal("missing logic")
		}

		if cont.Contains(base, "98") {
			tr.Fatal("missing logic")
		}
	})

	t.Run("[]string", func(tr *testing.T) {
		base := []string{"1", "2", "3"}

		if !cont.Contains(base, "2") {
			tr.Fatal("missing logic")
		}

		if cont.Contains(base, "4") {
			tr.Fatal("missing logic")
		}

		if cont.Contains(base, 1) {
			tr.Fatal("missing logic")
		}
	})

	t.Run("[]int", func(tr *testing.T) {
		base := []int{1, 2, 3}

		if !cont.Contains(base, 1) {
			tr.Fatal("missing logic")
		}

		if cont.Contains(base, "2") {
			tr.Fatal("missing logic")
		}

		if cont.Contains(base, 4) {
			tr.Fatal("missing logic")
		}
	})

	t.Run("map", func(tr *testing.T) {
		now := time.Now()
		base := map[string]interface{}{"a": "1", "b": "2", "c": 3, "d": now}

		if !cont.Contains(base, "1") {
			tr.Fatal("missing logic")
		}

		if !cont.Contains(base, "a") {
			tr.Fatal("missing logic")
		}

		if !cont.Contains(base, 3) {
			tr.Fatal("missing logic")
		}

		if !cont.Contains(base, now) {
			tr.Fatal("missing logic")
		}

		if cont.Contains(base, "e") {
			tr.Fatal("missing logic")
		}
	})
}
