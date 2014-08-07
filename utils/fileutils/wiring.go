package fileutils

import "github.com/cf-guardian/guardian-backend/utils/gerror"


func Wire() (Fileutils, gerror.Gerror) {
	return WireWith()
}

func WireWith() (Fileutils, gerror.Gerror) {
	return newFileutils()
}
