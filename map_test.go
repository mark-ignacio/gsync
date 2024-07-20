package gsync_test

import (
	"testing"

	gsync "github.com/mark-ignacio/gsyncmap"
)

// Runs tests in parallel in order to race condition the heck out of it all
func TestMap(t *testing.T) {
	var m gsync.Map[string, int]
	t.Run("LoadStore", func(t *testing.T) {
		t.Parallel()
		m.Store("a", 100)
		value, ok := m.Load("a")
		if !ok {
			t.Fatal("expected map load to work")
		}
		if value != 100 {
			t.Fatalf("unexpected map value: %d", value)
		}
	})
	t.Run("LoadDelete", func(t *testing.T) {
		t.Parallel()
		const key = "loaddelete"
		m.Store(key, 500)
		value, loaded := m.LoadAndDelete(key)
		if !loaded {
			t.Fatal("expected value to be loaded")
		}
		if value != 500 {
			t.Fatalf("unexpected map value: %d", value)
		}
		_, ok := m.Load(key)
		if ok {
			t.Fatal("expected load miss")
		}
	})
	t.Run("LoadOrStore", func(t *testing.T) {
		t.Parallel()
		const key = "loaddelete"
		m.Store(key, 500)
		value, loaded := m.LoadAndDelete(key)
		if !loaded {
			t.Fatal("expected value to be loaded")
		}
		if value != 500 {
			t.Fatalf("unexpected map value: %d", value)
		}
		_, ok := m.Load(key)
		if ok {
			t.Fatal("expected load miss")
		}
	})
}
