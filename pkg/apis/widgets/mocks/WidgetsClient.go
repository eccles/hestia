// Code generated by mockery v2.10.0. DO NOT EDIT.

package mocks

import (
	context "context"

	grpc "google.golang.org/grpc"
	emptypb "google.golang.org/protobuf/types/known/emptypb"

	mock "github.com/stretchr/testify/mock"

	widgetsAPI "github.com/eccles/hestia/pkg/apis/widgets"
)

// WidgetsClient is an autogenerated mock type for the WidgetsClient type
type WidgetsClient struct {
	mock.Mock
}

// Create provides a mock function with given fields: ctx, in, opts
func (_m *WidgetsClient) Create(ctx context.Context, in *widgetsAPI.CreateRequest, opts ...grpc.CallOption) (*widgetsAPI.Widget, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, in)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *widgetsAPI.Widget
	if rf, ok := ret.Get(0).(func(context.Context, *widgetsAPI.CreateRequest, ...grpc.CallOption) *widgetsAPI.Widget); ok {
		r0 = rf(ctx, in, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*widgetsAPI.Widget)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *widgetsAPI.CreateRequest, ...grpc.CallOption) error); ok {
		r1 = rf(ctx, in, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Delete provides a mock function with given fields: ctx, in, opts
func (_m *WidgetsClient) Delete(ctx context.Context, in *widgetsAPI.DeleteRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, in)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *emptypb.Empty
	if rf, ok := ret.Get(0).(func(context.Context, *widgetsAPI.DeleteRequest, ...grpc.CallOption) *emptypb.Empty); ok {
		r0 = rf(ctx, in, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*emptypb.Empty)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *widgetsAPI.DeleteRequest, ...grpc.CallOption) error); ok {
		r1 = rf(ctx, in, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// FindByID provides a mock function with given fields: ctx, in, opts
func (_m *WidgetsClient) FindByID(ctx context.Context, in *widgetsAPI.FindRequest, opts ...grpc.CallOption) (*widgetsAPI.Widget, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, in)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *widgetsAPI.Widget
	if rf, ok := ret.Get(0).(func(context.Context, *widgetsAPI.FindRequest, ...grpc.CallOption) *widgetsAPI.Widget); ok {
		r0 = rf(ctx, in, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*widgetsAPI.Widget)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *widgetsAPI.FindRequest, ...grpc.CallOption) error); ok {
		r1 = rf(ctx, in, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// List provides a mock function with given fields: ctx, in, opts
func (_m *WidgetsClient) List(ctx context.Context, in *widgetsAPI.ListRequest, opts ...grpc.CallOption) (*widgetsAPI.ListResponse, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, in)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *widgetsAPI.ListResponse
	if rf, ok := ret.Get(0).(func(context.Context, *widgetsAPI.ListRequest, ...grpc.CallOption) *widgetsAPI.ListResponse); ok {
		r0 = rf(ctx, in, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*widgetsAPI.ListResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *widgetsAPI.ListRequest, ...grpc.CallOption) error); ok {
		r1 = rf(ctx, in, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Update provides a mock function with given fields: ctx, in, opts
func (_m *WidgetsClient) Update(ctx context.Context, in *widgetsAPI.UpdateRequest, opts ...grpc.CallOption) (*widgetsAPI.Widget, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, in)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *widgetsAPI.Widget
	if rf, ok := ret.Get(0).(func(context.Context, *widgetsAPI.UpdateRequest, ...grpc.CallOption) *widgetsAPI.Widget); ok {
		r0 = rf(ctx, in, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*widgetsAPI.Widget)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *widgetsAPI.UpdateRequest, ...grpc.CallOption) error); ok {
		r1 = rf(ctx, in, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}