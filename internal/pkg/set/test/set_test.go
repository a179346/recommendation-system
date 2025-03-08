package set_test

import (
	"testing"

	"github.com/a179346/recommendation-system/internal/pkg/set"
)

func TestSet(t *testing.T) {
	t.Run("Has", func(t *testing.T) {
		s := set.New[int]()
		s.Add(1)
		if !s.Has(1) {
			t.Error("expected to have 1 after adding")
		}
		if s.Has(2) {
			t.Error("expected to not have 2")
		}
	})

	t.Run("Add", func(t *testing.T) {
		s := set.New[int]()
		if s.Has(1) {
			t.Error("expected to not have 1 initially")
		}
		if !s.Add(1) {
			t.Error("expected to add 1")
		}
		if !s.Has(1) {
			t.Error("expected to have 1 after adding")
		}
	})

	t.Run("Remove", func(t *testing.T) {
		type person struct {
			name string
		}
		s := set.New[person]()
		john := person{name: "John"}

		s.Add(john)
		if !s.Has(john) {
			t.Error("expected to have john after adding")
		}
		if !s.Remove(john) {
			t.Error("expected to remove john successfully")
		}
		if s.Has(john) {
			t.Error("expected to not have john after removing")
		}
	})

	t.Run("Len", func(t *testing.T) {
		s := set.New[int]()

		if s.Len() != 0 {
			t.Error("expected length to be 0 initially")
		}
		s.Add(1)
		if s.Len() != 1 {
			t.Error("expected length to be 1 after adding 1")
		}
		s.Add(2)
		if s.Len() != 2 {
			t.Error("expected length to be 2 after adding 2")
		}
		s.Remove(1)
		if s.Len() != 1 {
			t.Error("expected length to be 1 after removing 1")
		}
	})

	t.Run("All", func(t *testing.T) {
		data := []int{1, 2, 3}
		s1 := set.New[int]()
		s2 := set.New[int]()

		for _, v := range data {
			s1.Add(v)
			s2.Add(v)
		}

		expectedLenOfS2 := 3
		for v := range s1.All() {
			s2.Remove(v)
			expectedLenOfS2--
			if s2.Len() != expectedLenOfS2 {
				t.Errorf("expected length of s2 to be %d after removing %d", expectedLenOfS2, v)
			}
		}
	})
}
