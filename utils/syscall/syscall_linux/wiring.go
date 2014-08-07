package syscall_linux

import (
	"github.com/cf-guardian/guardian-backend/utils/syscall"
	"github.com/cf-guardian/guardian-backend/utils/gerror"
)

func WireFS() (syscall.SyscallFS, gerror.Gerror) {
	return WireFSWith()
}

func WireFSWith() (syscall.SyscallFS, gerror.Gerror) {
	return newFS()
}
