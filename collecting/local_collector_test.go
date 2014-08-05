package collecting

import (
	api "github.com/cf-guardian/guardian-backend/temp_libcontainer_api"
	identity "github.com/cf-guardian/guardian-backend/temp_libcontainer_api/identity"
	"testing"
)

var identifier identity.Identifier = identity.CreateSimpleIdentifier()

type t_container struct {
	id api.Id
}

func (ctr t_container) Id() api.Id {
	return ctr.id
}

func createContainer() api.Container {
	return &t_container{id: identifier.Generate()}
}

// test local_collector implementation of the Collector interface.

func TestLCCreate(t *testing.T) {
	lc := LocalCollection()
	nms := lc.Names()
	if len(nms) > 0 {
		t.Errorf("New collection should be empty.")
		return
	}
}

const (
	name1 = "name one"
	name2 = "name two"
)

func TestLCAdd(t *testing.T) {
	lc := LocalCollection()
	ctr1 := createContainer()
	if err := lc.Add(api.Name(name1), ctr1); err != nil {
		t.Fatalf("Add failed. %s", err)
	}

	if nms := lc.Names(); len(nms) != 1 {
		t.Fatalf("Collection should have one entry.")
	}
	if nm := lc.Names()[0]; name1 != nm {
		t.Fatalf("Collection should have single entry named %q, but %q instead.", name1, nm)
	}

	ctr2 := createContainer()
	if err := lc.Add(api.Name(name2), ctr2); err != nil {
		t.Fatalf("Add failed. %s", err)
	}

	if nms := lc.Names(); len(nms) != 2 {
		t.Fatalf("Collection should have two entries, not %d.", len(nms))
	}
}

func TestLCRemove(t *testing.T) {
	lc := LocalCollection()
	ctr1 := createContainer()
	if err := lc.Add(api.Name(name1), ctr1); err != nil {
		t.Fatalf("Add failed. %s", err)
	}
	if err := lc.Add(api.Name(name2), ctr1); err != nil {
		t.Fatalf("Add failed. %s", err)
	}

	if err := lc.Remove(api.Name(name1)); err != nil {
		t.Fatalf("Could not Remove(name1). %s", err)
	}

	if nms := lc.Names(); len(nms) != 1 {
		t.Fatalf("Entry not removed. Names() = %v", nms)
	}

	if nm := lc.Names()[0]; name2 != nm {
		t.Fatalf("Collection should have single entry named %q, but %q instead.", name2, nm)
	}
}

func TestLCGet(t *testing.T) {
	lc := LocalCollection()
	ctr1 := createContainer()
	ctr2 := createContainer()
	if err := lc.Add(api.Name(name1), ctr1); err != nil {
		t.Fatalf("Add failed. %s", err)
	}
	if err := lc.Add(api.Name(name2), ctr2); err != nil {
		t.Fatalf("Add failed. %s", err)
	}
	gctr1, err1 := lc.Get(api.Name(name1))
	if err1 != nil {
		t.Fatalf("Get(%q) failed with error %s", name1, err1)
	}
	if ctr1.Id() != gctr1.Id() {
		t.Fatalf("Correct container not Get(name=%q), id=%q, id of result = %q.", name1, ctr1.Id(), gctr1.Id())
	}
	gctr2, err2 := lc.Get(api.Name(name2))
	if err2 != nil {
		t.Fatalf("Get(%q) failed with error %s", name2, err2)
	}
	if ctr2.Id() != gctr2.Id() {
		t.Fatalf("Correct container not Get(name=%q), id=%q, id of result = %q.", name2, ctr2.Id(), gctr2.Id())
	}

	if err := lc.Remove(api.Name(name1)); err != nil {
		t.Fatalf("Could not Remove(name1). %s", err)
	}
	_, err3 := lc.Get(api.Name(name1))
	if err3 == nil {
		t.Fatalf("Get(%q) should fail!", name1)
	}
}
