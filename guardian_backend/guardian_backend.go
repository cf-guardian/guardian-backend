// a Linux backend for Warden based on libcontainer.
package guardian_backend

import (
	"github.com/cloudfoundry-incubator/garden/warden"
	"github.com/cloudfoundry-incubator/warden-linux/system_info"
)

type GuardianBackend struct {
	systemInfo system_info.Provider
}

func New(systemInfo system_info.Provider) *GuardianBackend {
	return &GuardianBackend{
		systemInfo: systemInfo,
	}
}

func (b *GuardianBackend) Ping() error {
	return nil
}

func (b *GuardianBackend) Capacity() (warden.Capacity, error) {
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
		MaxContainers: 0, // TODO: should be non-zero
	}, nil
}

func (b *GuardianBackend) Create(spec warden.ContainerSpec) (warden.Container, error) {
	panic(`unimplemented`)
}
