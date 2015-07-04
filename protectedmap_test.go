// Provides a map that can be used in a protected and concurrent fashion.
// Map key must be a string, but the data can be anything.
package protectedmap

import (
	"fmt"
	"reflect"
	"strconv"
	"testing"
)

type TestData struct {
	ID   int
	Name string
}

func TestNewProtectedMap(t *testing.T) {
	m := New()
	if reflect.TypeOf(m).Name() != "ProtectedMap" {

		t.Error("Map is not the correct type")
	}

	if m.Count() != 0 {
		t.Error("New map is not empty")
	}
}

func TestSet(t *testing.T) {
	m := New()
	one := TestData{ID: 1, Name: "one"}
	two := TestData{ID: 2, Name: "two"}

	m.Set("one", one)
	m.Set("two", two)

	if m.Count() != 2 {
		t.Error("Map does not contain two items")
	}
	if _, ok := m.Get("one"); !ok {
		t.Error("Key 'one' did not exist.")
	}
}

func TestGet(t *testing.T) {
	m := New()

	m.Set("one", TestData{ID: 1, Name: "one"})
	m.Set("two", TestData{ID: 2, Name: "two"})
	m.Set("three", TestData{ID: 3, Name: "three"})

	test, ok := m.Get("two")
	gotten := test.(TestData)

	if !ok {
		t.Error("Unable to get item from map by key")
	}

	if gotten.ID != 2 || gotten.Name != "two" {
		t.Error("Wrong item was returned from map")
	}
}

func TestDelete(t *testing.T) {
	m := New()
	m.Set("one", TestData{ID: 1, Name: "one"})
	m.Set("two", TestData{ID: 2, Name: "two"})
	m.Set("three", TestData{ID: 3, Name: "three"})

	m.Delete("two")
	_, ok := m.Get("two")
	if ok {
		t.Error("Deleted key still exists")
	}
	if m.Count() != 2 {
		t.Error("Size of map was wrong")
	}
}

func TestCount(t *testing.T) {
	m := New()
	m.Set("one", TestData{ID: 1, Name: "one"})
	if m.Count() != 1 {
		t.Error("First count was wrong")
	}
	m.Set("two", TestData{ID: 2, Name: "two"})
	if m.Count() != 2 {
		t.Error("Second count was wrong")
	}
	m.Set("three", TestData{ID: 3, Name: "three"})
	if m.Count() != 3 {
		t.Error("Third count was wrong")
	}
}

func TestIterator(t *testing.T) {
	m := New()
	for i := 0; i < 100; i++ {
		str := strconv.Itoa(i)
		m.Set(str, TestData{ID: i, Name: str})
	}

	itr := m.Iterator()
	j := 0
	for item := range itr.Loop() {
		t.Logf("Iterating through %v.", item)
		j++
	}
	if j != 100 {
		t.Errorf("Did not loop through 100 items, found %v instead.", j)
	}
}

func TestIteratorBreak(t *testing.T) {
	m := New()
	for i := 0; i < 1000; i++ {
		str := strconv.Itoa(i)
		m.Set(str, TestData{ID: i, Name: str})
	}

	itr := m.Iterator()
	j := 0
	for _ = range itr.Loop() {
		if j == 60 {
			itr.Break()
			break
		}
		j++
	}
}

func ExampleIterator_full() {
	m := New()

	m.Set("1", "one")
	m.Set("2", "two")
	m.Set("3", "three")
	m.Set("4", "four")
	m.Set("5", "five")

	iter := m.Iterator()
	for data := range iter.Loop() {
		fmt.Printf("%s > %v\n", data.Key, data.Val)
	}
}

func ExampleIterator_break() {
	m := New()

	m.Set("1", "one")
	m.Set("2", "two")
	m.Set("3", "three")
	m.Set("4", "four")
	m.Set("5", "five")

	iter := m.Iterator()
	for data := range iter.Loop() {
		if data.Key == "3" {
			iter.Break()
			break
		}
		fmt.Printf("%s > %v\n", data.Key, data.Val)
	}
}
