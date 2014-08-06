// Automatically generated by MockGen. DO NOT EDIT!
// Source: /Users/spowell/dev/go/src/github.com/cf-guardian/guardian-backend/development/scripts/../../utils/fileutils/fileutils.go

package mock_fileutils

import (
	gomock "code.google.com/p/gomock/gomock"
	gerror "github.com/cf-guardian/guardian-backend/utils/gerror"
	os "os"
)

// Mock of Fileutils interface
type MockFileutils struct {
	ctrl     *gomock.Controller
	recorder *_MockFileutilsRecorder
}

// Recorder for MockFileutils (not exported)
type _MockFileutilsRecorder struct {
	mock *MockFileutils
}

func NewMockFileutils(ctrl *gomock.Controller) *MockFileutils {
	mock := &MockFileutils{ctrl: ctrl}
	mock.recorder = &_MockFileutilsRecorder{mock}
	return mock
}

func (_m *MockFileutils) EXPECT() *_MockFileutilsRecorder {
	return _m.recorder
}

func (_m *MockFileutils) Copy(destPath string, srcPath string) gerror.Gerror {
	ret := _m.ctrl.Call(_m, "Copy", destPath, srcPath)
	ret0, _ := ret[0].(gerror.Gerror)
	return ret0
}

func (_mr *_MockFileutilsRecorder) Copy(arg0, arg1 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Copy", arg0, arg1)
}

func (_m *MockFileutils) Exists(path string) bool {
	ret := _m.ctrl.Call(_m, "Exists", path)
	ret0, _ := ret[0].(bool)
	return ret0
}

func (_mr *_MockFileutilsRecorder) Exists(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Exists", arg0)
}

func (_m *MockFileutils) Filemode(path string) (os.FileMode, gerror.Gerror) {
	ret := _m.ctrl.Call(_m, "Filemode", path)
	ret0, _ := ret[0].(os.FileMode)
	ret1, _ := ret[1].(gerror.Gerror)
	return ret0, ret1
}

func (_mr *_MockFileutilsRecorder) Filemode(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Filemode", arg0)
}