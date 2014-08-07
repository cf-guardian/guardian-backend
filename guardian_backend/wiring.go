package guardian_backend

import (
	"github.com/cloudfoundry-incubator/garden/warden"
	"github.com/cf-guardian/guardian-backend/rootfs"
	"github.com/cf-guardian/guardian-backend/config_builder"
	"github.com/cf-guardian/guardian-backend/utils/gerror"
)

func Wire(depotPath string, rwBaseDir string) (warden.Backend, gerror.Gerror) {
	rootfs, err := rootfs.Wire(rwBaseDir)
	if err != nil {
		return nil, err
	}

	configBuilder, err := config_builder.Wire()
	if err != nil {
		return nil, err
	}

	return WireWith(depotPath, rootfs, configBuilder)
}

func WireWith(depotPath string, rootfs rootfs.RootFS, configBuilder config_builder.ConfigBuilder) (warden.Backend, gerror.Gerror) {
	return newGuardianBackend(depotPath, rootfs, configBuilder)
}
