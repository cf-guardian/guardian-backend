package config_builder

import "github.com/cf-guardian/guardian-backend/utils/gerror"

func Wire() (ConfigBuilder, gerror.Gerror) {
	return WireWith()
}

func WireWith() (ConfigBuilder, gerror.Gerror) {
	return newConfigBuilder()
}
