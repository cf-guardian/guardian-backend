package temp_libcontainer_api

import (
	"github.com/docker/libcontainer"
)

type Id string

type Container interface {
	Id() Id

	// Returns the current run state of the container.
	//
	// Errors: container no longer exists,
	//         system error.
	RunState() (*libcontainer.RunState, error)

	// Returns the current config of the container.
	Config() *libcontainer.Config

	// Start a process inside the container. Returns the PID of the new process (in the caller process's namespace) and a channel that will return the exit status of the process whenever it dies.
	//
	// Errors: container no longer exists,
	//         config is invalid,
	//         container is paused,
	//         system error.
	Start(*libcontainer.ProcessConfig) (pid int, exitChan chan int, err error)

	// Destroys the container after killing all running processes.
	//
	// Any event registrations are removed before the container is destroyed.
	// No error is returned if the container is already destroyed.
	//
	// Errors: system error.
	Destroy() error

	// Returns the PIDs inside this container. The PIDs are in the namespace of the calling process.
	//
	// Errors: container no longer exists,
	//         system error.
	//
	// Some of the returned PIDs may no longer refer to processes in the Container, unless
	// the Container state is PAUSED in which case every PID in the slice is valid.
	Processes() ([]int, error)

	// Returns statistics for the container.
	//
	// Errors: container no longer exists,
	//         system error.
	Stats() (*libcontainer.ContainerStats, error)

	// If the Container state is RUNNING or PAUSING, sets the Container state to PAUSING and pauses
	// the execution of any user processes. Asynchronously, when the container finished being paused the
	// state is changed to PAUSED.
	// If the Container state is PAUSED, do nothing.
	//
	// Errors: container no longer exists,
	//         system error.
	Pause() error

	// If the Container state is PAUSED, resumes the execution of any user processes in the
	// Container before setting the Container state to RUNNING.
	// If the Container state is RUNNING, do nothing.
	//
	// Errors: container no longer exists,
	//         system error.
	Resume() error
}
