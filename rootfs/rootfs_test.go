/*
   Copyright 2014 GoPivotal (UK) Limited.

   Licensed under the Apache License, Version 2.0 (the "License");
   you may not use this file except in compliance with the License.
   You may obtain a copy of the License at

       http://www.apache.org/licenses/LICENSE-2.0

   Unless required by applicable law or agreed to in writing, software
   distributed under the License is distributed on an "AS IS" BASIS,
   WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
   See the License for the specific language governing permissions and
   limitations under the License.
*/

package rootfs_test

import (
	"code.google.com/p/gomock/gomock"
	"errors"
	"github.com/cf-guardian/guardian-backend/utils/gerror"
	"github.com/cf-guardian/guardian-backend/utils/fileutils"
	"github.com/cf-guardian/guardian-backend/utils/fileutils/mock_fileutils"
	"github.com/cf-guardian/guardian-backend/rootfs"
	"github.com/cf-guardian/guardian-backend/utils/syscall/mock_syscall"
	"github.com/cf-guardian/guardian-backend/test_support"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func TestNilSyscallFS(t *testing.T) {
	mockCtrl, mockFileUtils, _ := setupMocks(t)
	defer mockCtrl.Finish()

	rfs, gerr := rootfs.WireWith("", nil, mockFileUtils)
	if rfs != nil || !gerr.EqualTag(rootfs.ErrNilSyscallFS) {
		t.Errorf("Incorrect return values (%s, %s)", rfs, gerr)
		return
	}
}

func TestNonExistentReadWriteBaseDir(t *testing.T) {
	mockCtrl, mockFileUtils, mockSyscallFS := setupMocks(t)
	defer mockCtrl.Finish()

	mockFileUtils.EXPECT().Filemode("/nosuch").Return(os.FileMode(0), gerror.New(fileutils.ErrFileNotFound, "an error"))

	rfs, gerr := rootfs.WireWith("/nosuch", mockSyscallFS, mockFileUtils)
	if rfs != nil || !gerr.EqualTag(rootfs.ErrRwBaseDirMissing) {
		t.Errorf("Incorrect return values (%s, %s)", rfs, gerr)
		return
	}
}

func TestNonDirReadWriteBaseDir(t *testing.T) {
	mockCtrl, mockFileUtils, mockSyscallFS := setupMocks(t)
	defer mockCtrl.Finish()

	tempDir := test_support.CreateTempDir()
	defer test_support.CleanupDirs(t, tempDir)

	filePath := test_support.CreateFile(tempDir, "testFile")
	mockFileUtils.EXPECT().Filemode(filePath).Return(os.FileMode(0700), nil)

	rfs, gerr := rootfs.WireWith(filePath, mockSyscallFS, mockFileUtils)
	if rfs != nil || !gerr.EqualTag(rootfs.ErrRwBaseDirIsFile) {
		t.Errorf("Incorrect return values (%s, %s)", rfs, gerr)
		return
	}
}

func TestReadOnlyReadWriteBaseDir(t *testing.T) {
	mockCtrl, mockFileUtils, mockSyscallFS := setupMocks(t)
	defer mockCtrl.Finish()

	tempDir := test_support.CreateTempDir()
	defer test_support.CleanupDirs(t, tempDir)

	dirPath := test_support.CreateDirWithMode(tempDir, "test-rootfs", os.FileMode(0400))
	mockFileUtils.EXPECT().Filemode(dirPath).Return(os.ModeDir|os.FileMode(0100), nil)

	rfs, gerr := rootfs.WireWith(dirPath, mockSyscallFS, mockFileUtils)
	if rfs != nil || !gerr.EqualTag(rootfs.ErrRwBaseDirNotRw) {
		t.Errorf("Incorrect return values (%s, %s)", rfs, gerr)
		return
	}
}

func TestGenerate(t *testing.T) {
	mockCtrl, mockFileUtils, mockSyscallFS := setupMocks(t)
	defer mockCtrl.Finish()

	tempDir := test_support.CreateTempDir()
	defer test_support.CleanupDirs(t, tempDir)

	mockFileUtils.EXPECT().Filemode(tempDir).Return(os.ModeDir|os.FileMode(0700), nil)
	rfs, gerr := rootfs.WireWith(tempDir, mockSyscallFS, mockFileUtils)
	if gerr != nil {
		t.Errorf("%s", gerr)
		return
	}
	prototypeDir := test_support.CreatePrototype(tempDir)

	mockSyscallFS.EXPECT().BindMountReadOnly(prototypeDir, test_support.NewStringPrefixMatcher(filepath.Join(tempDir, "mnt")))

	for _, dir := range test_support.RootFSDirs() {
		srcMatcher := test_support.NewStringRegexMatcher(filepath.Join(tempDir, "tmp-rootfs-[^/]*", dir))
		mntMatcher := test_support.NewStringRegexMatcher(filepath.Join(tempDir, "mnt-[^/]*", dir))
		mockFileUtils.EXPECT().Exists(srcMatcher).Return(true).AnyTimes()
		mockFileUtils.EXPECT().Exists(mntMatcher).Return(true).AnyTimes()
		mockSyscallFS.EXPECT().BindMountReadWrite(srcMatcher, mntMatcher)
	}

	root, gerr := rfs.Generate(prototypeDir)
	if gerr != nil {
		t.Errorf("%s", gerr)
		return
	}

	rootPrefix := filepath.Join(tempDir, "mnt-")
	if !strings.HasPrefix(root, rootPrefix) {
		t.Errorf("root was %s, but expected it to have prefix %s", root, rootPrefix)
		return
	}
}

func TestGenerateBackoutAfterBindMountReadWriteError(t *testing.T) {
	numDirs := len(test_support.RootFSDirs())
	for i := 0; i <= numDirs-1; i++ {
		testGenerateBackoutAfterBindMountReadWriteError(i, t)
	}
}

func testGenerateBackoutAfterBindMountReadWriteError(i int, t *testing.T) {
	mockCtrl, mockFileUtils, mockSyscallFS := setupMocks(t)
	defer mockCtrl.Finish()

	tempDir := test_support.CreateTempDir()
	defer test_support.CleanupDirs(t, tempDir)

	mockFileUtils.EXPECT().Filemode(tempDir).Return(os.ModeDir|os.FileMode(0700), nil)
	rfs, gerr := rootfs.WireWith(tempDir, mockSyscallFS, mockFileUtils)
	if gerr != nil {
		t.Errorf("%s", gerr)
		return
	}
	prototypeDir := filepath.Join(tempDir, "test-prototype")

	mainMountPointMatcher := test_support.NewStringRegexMatcher(filepath.Join(tempDir, `mnt-[\d]*$`))
	mockSyscallFS.EXPECT().BindMountReadOnly(prototypeDir, mainMountPointMatcher)
	mockSyscallFS.EXPECT().Unmount(mainMountPointMatcher)

	dirs := test_support.RootFSDirs()
	for j := 0; j < i; j++ {
		dir := dirs[j]

		srcMatcher := test_support.NewStringRegexMatcher(filepath.Join(tempDir, "tmp-rootfs-[^/]*", dir))
		mntMatcher := test_support.NewStringRegexMatcher(filepath.Join(tempDir, "mnt-[^/]*", dir))
		mockFileUtils.EXPECT().Exists(srcMatcher).Return(true).AnyTimes()
		mockFileUtils.EXPECT().Exists(mntMatcher).Return(true).AnyTimes()
		mockSyscallFS.EXPECT().BindMountReadWrite(srcMatcher, mntMatcher)
		mockSyscallFS.EXPECT().Unmount(mntMatcher)
	}

	failingDir := dirs[i]
	srcMatcher := test_support.NewStringRegexMatcher(filepath.Join(tempDir, "tmp-rootfs-[^/]*", failingDir))
	mntMatcher := test_support.NewStringRegexMatcher(filepath.Join(tempDir, "mnt-[^/]*", failingDir))
	mockFileUtils.EXPECT().Exists(srcMatcher).Return(true).AnyTimes()
	mockFileUtils.EXPECT().Exists(mntMatcher).Return(true).AnyTimes()
	mockSyscallFS.EXPECT().BindMountReadWrite(srcMatcher, mntMatcher).Return(errors.New("an error"))

	root, gerr := rfs.Generate(prototypeDir)
	if gerr == nil {
		t.Errorf("Unexpected return values %s, %s", root, gerr)
		return
	}
}

func TestRemove(t *testing.T) {
	mockCtrl, mockFileUtils, mockSyscallFS := setupMocks(t)
	defer mockCtrl.Finish()

	tempDir := test_support.CreateTempDir()
	defer test_support.CleanupDirs(t, tempDir)

	mockFileUtils.EXPECT().Filemode(tempDir).Return(os.ModeDir|os.FileMode(0700), nil)
	rfs, gerr := rootfs.WireWith(tempDir, mockSyscallFS, mockFileUtils)
	if gerr != nil {
		t.Errorf("%s", gerr)
		return
	}

	root := "/test-rootfs"

	for _, dir := range test_support.RootFSDirs() {
		mockSyscallFS.EXPECT().Unmount(filepath.Join(root, dir)).Return(nil)
	}
	mockSyscallFS.EXPECT().Unmount(root).Return(nil)


	gerr = rfs.Remove(root)
	if gerr != nil {
		t.Errorf("%s", gerr)
		return
	}
}

func TestRemoveUnmountSubdirFailure(t *testing.T) {
	numDirs := len(test_support.RootFSDirs())
	for i := 0; i <= numDirs-1; i++ {
		testRemoveUnmountSubdirFailure(i, t)
	}
}

func testRemoveUnmountSubdirFailure(i int, t *testing.T) {
	mockCtrl, mockFileUtils, mockSyscallFS := setupMocks(t)
	defer mockCtrl.Finish()

	tempDir := test_support.CreateTempDir()
	defer test_support.CleanupDirs(t, tempDir)

	mockFileUtils.EXPECT().Filemode(tempDir).Return(os.ModeDir|os.FileMode(0700), nil)
	rfs, gerr := rootfs.WireWith(tempDir, mockSyscallFS, mockFileUtils)
	if gerr != nil {
		t.Errorf("%s", gerr)
		return
	}

	root := "/test-rootfs"

	dirs := test_support.RootFSDirs()
	for j, dir := range dirs {
		var err error
		if j == i {
			err = errors.New("an error")
		}
		mockSyscallFS.EXPECT().Unmount(filepath.Join(root, dir)).Return(err)
	}

	gerr = rfs.Remove(root)
	if gerr == nil {
		t.Errorf("%s", gerr)
		return
	}
	if gerr == nil || !gerr.EqualTag(rootfs.ErrUnmountSubdir) {
		t.Errorf("Incorrect error %s", gerr)
		return
	}
}

func TestRemoveUnmountRootFailure(t *testing.T) {
	mockCtrl, mockFileUtils, mockSyscallFS := setupMocks(t)
	defer mockCtrl.Finish()

	tempDir := test_support.CreateTempDir()
	defer test_support.CleanupDirs(t, tempDir)

	mockFileUtils.EXPECT().Filemode(tempDir).Return(os.ModeDir|os.FileMode(0700), nil)
	rfs, gerr := rootfs.WireWith(tempDir, mockSyscallFS, mockFileUtils)
	if gerr != nil {
		t.Errorf("%s", gerr)
		return
	}

	root := "/test-rootfs"

	for _, dir := range test_support.RootFSDirs() {
		mockSyscallFS.EXPECT().Unmount(filepath.Join(root, dir)).Return(nil)
	}
	mockSyscallFS.EXPECT().Unmount(root).Return(errors.New("an error"))

	gerr = rfs.Remove(root)
	if gerr == nil {
		t.Errorf("%s", gerr)
		return
	}
	if gerr == nil || !gerr.EqualTag(rootfs.ErrUnmountRoot) {
		t.Errorf("Incorrect error %s", gerr)
		return
	}
}
func setupMocks(t *testing.T) (*gomock.Controller, *mock_fileutils.MockFileutils, *mock_syscall.MockSyscallFS) {
	mockCtrl := gomock.NewController(t)
	mockFileUtils := mock_fileutils.NewMockFileutils(mockCtrl)
	mockSyscallFS := mock_syscall.NewMockSyscallFS(mockCtrl)
	return mockCtrl, mockFileUtils, mockSyscallFS
}
