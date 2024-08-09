package orderedmap

import (
	"testing"
)

func TestOrderedMap(t *testing.T) {
	om := NewOrderedMap()
	om.Set("a", "1")
	om.Set("b", "2")
	om.Set("c", "3")
	om.Delete("b")
	pairs := om.GetAll()
	if len(pairs) != 2 {
        t.Fatalf(`len of pairs = %d, want %d`, len(pairs), 2)
	}

	if pairs[0].Key != "a" {
		t.Fatalf(`pairs[0].Key = %s, want "a"`, pairs[0].Key)
	}

	if pairs[1].Key != "c" {
		t.Fatalf(`pairs[1].Key = %s, want "c"`, pairs[1].Key)
	}

	b := om.Get("b")
	if b != "" {
		t.Fatalf(`om.Get("b") = %s, want ""`, b)
	}

	c := om.Get("c")
	if c != "3" {
		t.Fatalf(`om.Get("c") = %s, want "3"`, c)
	}
}
