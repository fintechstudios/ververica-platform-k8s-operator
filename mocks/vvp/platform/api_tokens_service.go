// Code generated by mockery v1.0.0. DO NOT EDIT.

package mocks

import (
	context "context"

	mock "github.com/stretchr/testify/mock"

	platformapi "github.com/fintechstudios/ververica-platform-k8s-operator/pkg/vvp/platform-api"
)

// APITokensService is an autogenerated mock type for the APITokensService type
type APITokensService struct {
	mock.Mock
}

// CreateAPIToken provides a mock function with given fields: ctx, namespaceName, token
func (_m *APITokensService) CreateAPIToken(ctx context.Context, namespaceName string, token platformapi.ApiToken) (*platformapi.ApiToken, error) {
	ret := _m.Called(ctx, namespaceName, token)

	var r0 *platformapi.ApiToken
	if rf, ok := ret.Get(0).(func(context.Context, string, platformapi.ApiToken) *platformapi.ApiToken); ok {
		r0 = rf(ctx, namespaceName, token)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*platformapi.ApiToken)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string, platformapi.ApiToken) error); ok {
		r1 = rf(ctx, namespaceName, token)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DeleteAPIToken provides a mock function with given fields: ctx, namespaceName, name
func (_m *APITokensService) DeleteAPIToken(ctx context.Context, namespaceName string, name string) error {
	ret := _m.Called(ctx, namespaceName, name)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string, string) error); ok {
		r0 = rf(ctx, namespaceName, name)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// GetAPIToken provides a mock function with given fields: ctx, namespaceName, name
func (_m *APITokensService) GetAPIToken(ctx context.Context, namespaceName string, name string) (*platformapi.ApiToken, error) {
	ret := _m.Called(ctx, namespaceName, name)

	var r0 *platformapi.ApiToken
	if rf, ok := ret.Get(0).(func(context.Context, string, string) *platformapi.ApiToken); ok {
		r0 = rf(ctx, namespaceName, name)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*platformapi.ApiToken)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string, string) error); ok {
		r1 = rf(ctx, namespaceName, name)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
