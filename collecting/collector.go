package collecting

import (
	api "github.com/cf-guardian/guardian-backend/temp_libcontainer_api"
)

// Manage a collection of containers by name.
// In the collection each container is associated with an external name.
type Collector interface {

	// Add a Container to this collection with the given name.
	Add(name api.Name, ctr api.Container) error

	// Remove the given name from this collection.
	Remove(name api.Name) error

	// Return the Container associated with the given name from this collection.
	Get(name api.Name) (api.Container, error)

	// Return the (external) names in this collection.
	Names() []api.Name
}
