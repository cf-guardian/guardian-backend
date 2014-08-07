package config_builder

import (
	"github.com/docker/libcontainer"
	"github.com/cf-guardian/guardian-backend/utils/gerror"
)

type ConfigBuilder interface {

	// Builds a libcontainer configuration.
	Build() *libcontainer.Config
}

type configBuilder struct {
}

func newConfigBuilder() (ConfigBuilder, gerror.Gerror) {
	return &configBuilder{}, nil
}

func (cb *configBuilder) Build() *libcontainer.Config {
	return &libcontainer.Config{}
}
