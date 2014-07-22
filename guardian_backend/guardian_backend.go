// a Linux backend for Warden based on libcontainer.
package guardian_backend

import (
	"github.com/cf-guardian/guardian-backend/container"
	"github.com/cf-guardian/guardian-backend/system_info"
	"github.com/cloudfoundry-incubator/garden/warden"
	"time"
)

type guardianBackend struct {
	systemInfo system_info.Provider
}

func New(depotPath string) warden.Backend {
	systemInfo := system_info.NewProvider(depotPath)
	return &guardianBackend{
		systemInfo: systemInfo,
	}
}

func (b *guardianBackend) Ping() error {
	return nil
}

func (b *guardianBackend) Capacity() (warden.Capacity, error) {
	totalMemory, err := b.systemInfo.TotalMemory()
	if err != nil {
		return warden.Capacity{}, err
	}

	totalDisk, err := b.systemInfo.TotalDisk()
	if err != nil {
		return warden.Capacity{}, err
	}

	return warden.Capacity{
		MemoryInBytes: totalMemory,
		DiskInBytes:   totalDisk,
		MaxContainers: 0, // TODO[gn]: should this be non-zero?
	}, nil
}

func (b *guardianBackend) Create(spec warden.ContainerSpec) (warden.Container, error) {
	_, err := container.New(spec)
	if err != nil {
		return nil, err
	}
	// TODO[sp]: container management
	panic(`unimplemented`)
}

func (b *guardianBackend) Destroy(handle string) error {
	panic(`unimplemented`)
}

func (b *guardianBackend) Containers(warden.Properties) ([]warden.Container, error) {
	panic(`unimplemented`)
}

func (b *guardianBackend) Lookup(handle string) (warden.Container, error) {
	panic(`unimplemented`)
}

// Start the backend.
func (b *guardianBackend) Start() error {
	// TODO: is recovery required?
	return nil
}

func (b *guardianBackend) Stop() {
	panic(`unimplemented`)
}

func (b *guardianBackend) GraceTime(warden.Container) time.Duration {
	panic(`unimplemented`)
}
