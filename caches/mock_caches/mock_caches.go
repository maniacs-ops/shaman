// Automatically generated by MockGen. DO NOT EDIT!
// Source: caches/caches.go

package mock_caches

import (
	gomock "github.com/golang/mock/gomock"
)

// Mock of Cacher interface
type MockCacher struct {
	ctrl     *gomock.Controller
	recorder *_MockCacherRecorder
}

// Recorder for MockCacher (not exported)
type _MockCacherRecorder struct {
	mock *MockCacher
}

func NewMockCacher(ctrl *gomock.Controller) *MockCacher {
	mock := &MockCacher{ctrl: ctrl}
	mock.recorder = &_MockCacherRecorder{mock}
	return mock
}

func (_m *MockCacher) EXPECT() *_MockCacherRecorder {
	return _m.recorder
}

func (_m *MockCacher) InitializeDatabase() error {
	ret := _m.ctrl.Call(_m, "InitializeDatabase")
	ret0, _ := ret[0].(error)
	return ret0
}

func (_mr *_MockCacherRecorder) InitializeDatabase() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "InitializeDatabase")
}

func (_m *MockCacher) GetRecord(_param0 string) (string, error) {
	ret := _m.ctrl.Call(_m, "GetRecord", _param0)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockCacherRecorder) GetRecord(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "GetRecord", arg0)
}

func (_m *MockCacher) SetRecord(_param0 string, _param1 string) error {
	ret := _m.ctrl.Call(_m, "SetRecord", _param0, _param1)
	ret0, _ := ret[0].(error)
	return ret0
}

func (_mr *_MockCacherRecorder) SetRecord(arg0, arg1 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "SetRecord", arg0, arg1)
}

func (_m *MockCacher) ReviseRecord(_param0 string, _param1 string) error {
	ret := _m.ctrl.Call(_m, "ReviseRecord", _param0, _param1)
	ret0, _ := ret[0].(error)
	return ret0
}

func (_mr *_MockCacherRecorder) ReviseRecord(arg0, arg1 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "ReviseRecord", arg0, arg1)
}

func (_m *MockCacher) DeleteRecord(_param0 string) error {
	ret := _m.ctrl.Call(_m, "DeleteRecord", _param0)
	ret0, _ := ret[0].(error)
	return ret0
}

func (_mr *_MockCacherRecorder) DeleteRecord(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "DeleteRecord", arg0)
}

func (_m *MockCacher) ListRecords() ([]string, error) {
	ret := _m.ctrl.Call(_m, "ListRecords")
	ret0, _ := ret[0].([]string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockCacherRecorder) ListRecords() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "ListRecords")
}
