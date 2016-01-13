// Automatically generated by MockGen. DO NOT EDIT!
// Source: github.com/pivotal-cf/cf-watch/watch (interfaces: CLI)

package mocks

import (
	gomock "github.com/golang/mock/gomock"
)

// Mock of CLI interface
type MockCLI struct {
	ctrl     *gomock.Controller
	recorder *_MockCLIRecorder
}

// Recorder for MockCLI (not exported)
type _MockCLIRecorder struct {
	mock *MockCLI
}

func NewMockCLI(ctrl *gomock.Controller) *MockCLI {
	mock := &MockCLI{ctrl: ctrl}
	mock.recorder = &_MockCLIRecorder{mock}
	return mock
}

func (_m *MockCLI) EXPECT() *_MockCLIRecorder {
	return _m.recorder
}

func (_m *MockCLI) CliCommandWithoutTerminalOutput(_param0 ...string) ([]string, error) {
	_s := []interface{}{}
	for _, _x := range _param0 {
		_s = append(_s, _x)
	}
	ret := _m.ctrl.Call(_m, "CliCommandWithoutTerminalOutput", _s...)
	ret0, _ := ret[0].([]string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockCLIRecorder) CliCommandWithoutTerminalOutput(arg0 ...interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "CliCommandWithoutTerminalOutput", arg0...)
}
