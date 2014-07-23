package container

import (
	"fmt"
	"github.com/cloudfoundry-incubator/garden/warden"
	"github.com/docker/libcontainer"
	"io"
)

var lf libcontainer.Factory

func New(spec warden.ContainerSpec) (warden.Container, error) {
	cfg, err := config(spec)
	if err != nil {
		return nil, err
	}
	lc, err := lf.Create(cfg)
	if err != nil {
		return nil, fmt.Errorf("libcontainer Create failed: %s", err)
	}
	return &container{
		libcontainer: lc,
	}, nil
}

type container struct {
	libcontainer libcontainer.Container
}

func config(spec warden.ContainerSpec) (*libcontainer.Config, error) {
	panic("unimplemented")
}

func (c *container) Handle() string {
	panic("unimplemented")
}

func (c *container) Stop(kill bool) error {
	panic("unimplemented")
}

func (c *container) Info() (warden.ContainerInfo, error) {
	panic("unimplemented")
}

func (c *container) StreamIn(dstPath string, tarStream io.Reader) error {
	panic("unimplemented")
}

func (c *container) StreamOut(srcPath string) (io.ReadCloser, error) {
	panic("unimplemented")
}

func (c *container) LimitBandwidth(limits warden.BandwidthLimits) error {
	panic("unimplemented")
}

func (c *container) CurrentBandwidthLimits() (warden.BandwidthLimits, error) {
	panic("unimplemented")
}

func (c *container) LimitCPU(limits warden.CPULimits) error {
	panic("unimplemented")
}

func (c *container) CurrentCPULimits() (warden.CPULimits, error) {
	panic("unimplemented")
}

func (c *container) LimitDisk(limits warden.DiskLimits) error {
	panic("unimplemented")
}

func (c *container) CurrentDiskLimits() (warden.DiskLimits, error) {
	panic("unimplemented")
}

func (c *container) LimitMemory(limits warden.MemoryLimits) error {
	panic("unimplemented")
}

func (c *container) CurrentMemoryLimits() (warden.MemoryLimits, error) {
	panic("unimplemented")
}

func (c *container) NetIn(hostPort, containerPort uint32) (uint32, uint32, error) {
	panic("unimplemented")
}

func (c *container) NetOut(network string, port uint32) error {
	panic("unimplemented")
}

func (c *container) Run(warden.ProcessSpec, warden.ProcessIO) (warden.Process, error) {
	panic("unimplemented")
}

func (c *container) Attach(uint32, warden.ProcessIO) (warden.Process, error) {
	panic("unimplemented")
}
