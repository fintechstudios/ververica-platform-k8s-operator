// Code generated by mockery v1.0.0. DO NOT EDIT.

package mocks

import (
	context "context"

	mock "github.com/stretchr/testify/mock"

	platformapi "github.com/fintechstudios/ververica-platform-k8s-operator/pkg/vvp/platform-api"
)

// NamespacesService is an autogenerated mock type for the NamespacesService type
type NamespacesService struct {
	mock.Mock
}

// CreateNamespace provides a mock function with given fields: ctx, namespace
func (_m *NamespacesService) CreateNamespace(ctx context.Context, namespace platformapi.Namespace) (*platformapi.Namespace, error) {
	ret := _m.Called(ctx, namespace)

	var r0 *platformapi.Namespace
	if rf, ok := ret.Get(0).(func(context.Context, platformapi.Namespace) *platformapi.Namespace); ok {
		r0 = rf(ctx, namespace)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*platformapi.Namespace)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, platformapi.Namespace) error); ok {
		r1 = rf(ctx, namespace)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DeleteNamespace provides a mock function with given fields: ctx, namespaceName
func (_m *NamespacesService) DeleteNamespace(ctx context.Context, namespaceName string) (*platformapi.Namespace, error) {
	ret := _m.Called(ctx, namespaceName)

	var r0 *platformapi.Namespace
	if rf, ok := ret.Get(0).(func(context.Context, string) *platformapi.Namespace); ok {
		r0 = rf(ctx, namespaceName)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*platformapi.Namespace)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, namespaceName)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetNamespace provides a mock function with given fields: ctx, namespaceName
func (_m *NamespacesService) GetNamespace(ctx context.Context, namespaceName string) (*platformapi.Namespace, error) {
	ret := _m.Called(ctx, namespaceName)

	var r0 *platformapi.Namespace
	if rf, ok := ret.Get(0).(func(context.Context, string) *platformapi.Namespace); ok {
		r0 = rf(ctx, namespaceName)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*platformapi.Namespace)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, namespaceName)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UpdateNamespace provides a mock function with given fields: ctx, namespaceName, namespace
func (_m *NamespacesService) UpdateNamespace(ctx context.Context, namespaceName string, namespace platformapi.Namespace) (*platformapi.Namespace, error) {
	ret := _m.Called(ctx, namespaceName, namespace)

	var r0 *platformapi.Namespace
	if rf, ok := ret.Get(0).(func(context.Context, string, platformapi.Namespace) *platformapi.Namespace); ok {
		r0 = rf(ctx, namespaceName, namespace)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*platformapi.Namespace)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string, platformapi.Namespace) error); ok {
		r1 = rf(ctx, namespaceName, namespace)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
