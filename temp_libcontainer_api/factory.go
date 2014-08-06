package temp_libcontainer_api

import (
	"github.com/docker/libcontainer"
)

type Name string

type Factory interface {
	// Creates a new container. Generates a unique ID for the container and
	// starts the initial process inside the container.
	//
	// Returns the new container with a running process.
	//
	// Errors:
	// Config is invalid
	// System error
	//
	// On error, any partially created container parts are cleaned up (the operation is atomic).
	Create(config *libcontainer.Config) (Container, error)

}
