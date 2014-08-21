package collecting

// This collector implementation is the real collection management.
// Only processes with access to the collection state and the container states will use
// this implementation directly.

import (
	"fmt"
	api "github.com/cf-guardian/guardian-backend/temp_libcontainer_api"
)

// local error type
type c_error struct {
	e_msg  string
	e_name api.Name
}

func (cerr *c_error) Error() string {
	return fmt.Sprintf("name: %q. %s", cerr.e_name, cerr.e_msg)
}

type l_collection struct {
	coll map[api.Name]api.Container
}

func LocalCollection() Collector {
	return &l_collection{coll: make(map[api.Name]api.Container)}
}

func (coll *l_collection) Add(name api.Name, cont api.Container) error {
	if _, ispresent := coll.coll[name]; ispresent {
		return &c_error{"name is already used", name}
	} else {
		coll.coll[name] = cont
	}
	return nil
}

func (coll *l_collection) Remove(name api.Name) error {
	if _, ispresent := coll.coll[name]; ispresent {
		delete(coll.coll, name)
	} else {
		return &c_error{"name not in collection", name}
	}
	return nil
}

func (coll *l_collection) Get(name api.Name) (api.Container, error) {
	if id, ispresent := coll.coll[name]; ispresent {
		return id, nil
	}
	return nil, &c_error{"name not in collection", name}
}

func (coll *l_collection) Names() []api.Name {
	i, names := 0, make([]api.Name, len(coll.coll))
	for name, _ := range coll.coll {
		names[i] = name
		i++
	}
	return names
}
