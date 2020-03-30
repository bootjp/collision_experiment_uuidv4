package main

import (
	mapset "github.com/deckarep/golang-set"
	"testing"

	"github.com/satori/go.uuid"
)

func TestRun(t *testing.T) {
	testData := []struct {
		store   uuid.UUID
		compare uuid.UUID
		exist   bool
	}{
		{
			uuid.Must(uuid.FromString("6ba7b810-9dad-11d1-80b4-00c04fd430c8")),
			uuid.Must(uuid.FromString("6ba7b810-9dad-11d1-80b4-00c04fd430c8")),
			true,
		},
		{
			uuid.Must(uuid.FromString("6ba7b810-9dad-11d1-80b4-00c04fd430c1")),
			uuid.Must(uuid.FromString("6ba7b810-9dad-11d1-80b4-00c04fd430c8")),
			false,
		},
	}

	for _, v := range testData {
		// s := sync.Map{}
		store := mapset.NewThreadUnsafeSet()
		store.Add(v.store)
		if store.Contains(v.compare) != v.exist {
			t.Fatalf("store: %v, compare: %v exist: %v ", v.store, v.compare, v.exist)
		}
	}

}
