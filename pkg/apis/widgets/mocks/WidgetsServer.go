// Code generated by mockery v2.10.0. DO NOT EDIT.

package mocks

import (
	context "context"

	mock "github.com/stretchr/testify/mock"
	emptypb "google.golang.org/protobuf/types/known/emptypb"

	widgetsAPI "github.com/eccles/hestia/pkg/apis/widgets"
)

// WidgetsServer is an autogenerated mock type for the WidgetsServer type
type WidgetsServer struct {
	mock.Mock
}

// Create provides a mock function with given fields: _a0, _a1
func (_m *WidgetsServer) Create(_a0 context.Context, _a1 *widgetsAPI.CreateRequest) (*widgetsAPI.Widget, error) {
	ret := _m.Called(_a0, _a1)

	var r0 *widgetsAPI.Widget
	if rf, ok := ret.Get(0).(func(context.Context, *widgetsAPI.CreateRequest) *widgetsAPI.Widget); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*widgetsAPI.Widget)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *widgetsAPI.CreateRequest) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Delete provides a mock function with given fields: _a0, _a1
func (_m *WidgetsServer) Delete(_a0 context.Context, _a1 *widgetsAPI.DeleteRequest) (*emptypb.Empty, error) {
	ret := _m.Called(_a0, _a1)

	var r0 *emptypb.Empty
	if rf, ok := ret.Get(0).(func(context.Context, *widgetsAPI.DeleteRequest) *emptypb.Empty); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*emptypb.Empty)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *widgetsAPI.DeleteRequest) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// FindByID provides a mock function with given fields: _a0, _a1
func (_m *WidgetsServer) FindByID(_a0 context.Context, _a1 *widgetsAPI.FindRequest) (*widgetsAPI.Widget, error) {
	ret := _m.Called(_a0, _a1)

	var r0 *widgetsAPI.Widget
	if rf, ok := ret.Get(0).(func(context.Context, *widgetsAPI.FindRequest) *widgetsAPI.Widget); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*widgetsAPI.Widget)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *widgetsAPI.FindRequest) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// List provides a mock function with given fields: _a0, _a1
func (_m *WidgetsServer) List(_a0 context.Context, _a1 *widgetsAPI.ListRequest) (*widgetsAPI.ListResponse, error) {
	ret := _m.Called(_a0, _a1)

	var r0 *widgetsAPI.ListResponse
	if rf, ok := ret.Get(0).(func(context.Context, *widgetsAPI.ListRequest) *widgetsAPI.ListResponse); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*widgetsAPI.ListResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *widgetsAPI.ListRequest) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Update provides a mock function with given fields: _a0, _a1
func (_m *WidgetsServer) Update(_a0 context.Context, _a1 *widgetsAPI.UpdateRequest) (*widgetsAPI.Widget, error) {
	ret := _m.Called(_a0, _a1)

	var r0 *widgetsAPI.Widget
	if rf, ok := ret.Get(0).(func(context.Context, *widgetsAPI.UpdateRequest) *widgetsAPI.Widget); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*widgetsAPI.Widget)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *widgetsAPI.UpdateRequest) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// mustEmbedUnimplementedWidgetsServer provides a mock function with given fields:
func (_m *WidgetsServer) mustEmbedUnimplementedWidgetsServer() {
	_m.Called()
}