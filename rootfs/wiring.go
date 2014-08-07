package rootfs

import (
	"github.com/cf-guardian/guardian-backend/utils/fileutils"
	"github.com/cf-guardian/guardian-backend/utils/syscall"
	"github.com/cf-guardian/guardian-backend/utils/syscall/syscall_linux"
	"github.com/cf-guardian/guardian-backend/utils/gerror"
)

func Wire(rwBaseDir string) (RootFS, gerror.Gerror) {
	sc, err := syscall_linux.WireFS()
	if err != nil {
		return nil, err
	}

	f, err := fileutils.Wire()
	if err != nil {
		return nil, err
	}

	return WireWith(rwBaseDir, sc, f)
}

func WireWith(rwBaseDir string, sc syscall.SyscallFS, f fileutils.Fileutils) (RootFS, gerror.Gerror) {
	return newRootFS(sc, f, rwBaseDir)
}
