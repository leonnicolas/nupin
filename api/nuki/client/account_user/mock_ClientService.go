// Code generated by mockery v2.20.0. DO NOT EDIT.

package account_user

import (
	runtime "github.com/go-openapi/runtime"
	mock "github.com/stretchr/testify/mock"
)

// MockClientService is an autogenerated mock type for the ClientService type
type MockClientService struct {
	mock.Mock
}

// AccountUserResourceDeleteDelete provides a mock function with given fields: params, authInfo, opts
func (_m *MockClientService) AccountUserResourceDeleteDelete(params *AccountUserResourceDeleteDeleteParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*AccountUserResourceDeleteDeleteNoContent, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, params, authInfo)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *AccountUserResourceDeleteDeleteNoContent
	var r1 error
	if rf, ok := ret.Get(0).(func(*AccountUserResourceDeleteDeleteParams, runtime.ClientAuthInfoWriter, ...ClientOption) (*AccountUserResourceDeleteDeleteNoContent, error)); ok {
		return rf(params, authInfo, opts...)
	}
	if rf, ok := ret.Get(0).(func(*AccountUserResourceDeleteDeleteParams, runtime.ClientAuthInfoWriter, ...ClientOption) *AccountUserResourceDeleteDeleteNoContent); ok {
		r0 = rf(params, authInfo, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*AccountUserResourceDeleteDeleteNoContent)
		}
	}

	if rf, ok := ret.Get(1).(func(*AccountUserResourceDeleteDeleteParams, runtime.ClientAuthInfoWriter, ...ClientOption) error); ok {
		r1 = rf(params, authInfo, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// AccountUserResourceGetGet provides a mock function with given fields: params, authInfo, opts
func (_m *MockClientService) AccountUserResourceGetGet(params *AccountUserResourceGetGetParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*AccountUserResourceGetGetOK, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, params, authInfo)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *AccountUserResourceGetGetOK
	var r1 error
	if rf, ok := ret.Get(0).(func(*AccountUserResourceGetGetParams, runtime.ClientAuthInfoWriter, ...ClientOption) (*AccountUserResourceGetGetOK, error)); ok {
		return rf(params, authInfo, opts...)
	}
	if rf, ok := ret.Get(0).(func(*AccountUserResourceGetGetParams, runtime.ClientAuthInfoWriter, ...ClientOption) *AccountUserResourceGetGetOK); ok {
		r0 = rf(params, authInfo, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*AccountUserResourceGetGetOK)
		}
	}

	if rf, ok := ret.Get(1).(func(*AccountUserResourceGetGetParams, runtime.ClientAuthInfoWriter, ...ClientOption) error); ok {
		r1 = rf(params, authInfo, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// AccountUserResourcePostPost provides a mock function with given fields: params, authInfo, opts
func (_m *MockClientService) AccountUserResourcePostPost(params *AccountUserResourcePostPostParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*AccountUserResourcePostPostOK, *AccountUserResourcePostPostNoContent, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, params, authInfo)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *AccountUserResourcePostPostOK
	var r1 *AccountUserResourcePostPostNoContent
	var r2 error
	if rf, ok := ret.Get(0).(func(*AccountUserResourcePostPostParams, runtime.ClientAuthInfoWriter, ...ClientOption) (*AccountUserResourcePostPostOK, *AccountUserResourcePostPostNoContent, error)); ok {
		return rf(params, authInfo, opts...)
	}
	if rf, ok := ret.Get(0).(func(*AccountUserResourcePostPostParams, runtime.ClientAuthInfoWriter, ...ClientOption) *AccountUserResourcePostPostOK); ok {
		r0 = rf(params, authInfo, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*AccountUserResourcePostPostOK)
		}
	}

	if rf, ok := ret.Get(1).(func(*AccountUserResourcePostPostParams, runtime.ClientAuthInfoWriter, ...ClientOption) *AccountUserResourcePostPostNoContent); ok {
		r1 = rf(params, authInfo, opts...)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*AccountUserResourcePostPostNoContent)
		}
	}

	if rf, ok := ret.Get(2).(func(*AccountUserResourcePostPostParams, runtime.ClientAuthInfoWriter, ...ClientOption) error); ok {
		r2 = rf(params, authInfo, opts...)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// AccountUsersResourceGetGet provides a mock function with given fields: params, authInfo, opts
func (_m *MockClientService) AccountUsersResourceGetGet(params *AccountUsersResourceGetGetParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*AccountUsersResourceGetGetOK, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, params, authInfo)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *AccountUsersResourceGetGetOK
	var r1 error
	if rf, ok := ret.Get(0).(func(*AccountUsersResourceGetGetParams, runtime.ClientAuthInfoWriter, ...ClientOption) (*AccountUsersResourceGetGetOK, error)); ok {
		return rf(params, authInfo, opts...)
	}
	if rf, ok := ret.Get(0).(func(*AccountUsersResourceGetGetParams, runtime.ClientAuthInfoWriter, ...ClientOption) *AccountUsersResourceGetGetOK); ok {
		r0 = rf(params, authInfo, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*AccountUsersResourceGetGetOK)
		}
	}

	if rf, ok := ret.Get(1).(func(*AccountUsersResourceGetGetParams, runtime.ClientAuthInfoWriter, ...ClientOption) error); ok {
		r1 = rf(params, authInfo, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// AccountUsersResourcePutPut provides a mock function with given fields: params, authInfo, opts
func (_m *MockClientService) AccountUsersResourcePutPut(params *AccountUsersResourcePutPutParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*AccountUsersResourcePutPutOK, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, params, authInfo)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *AccountUsersResourcePutPutOK
	var r1 error
	if rf, ok := ret.Get(0).(func(*AccountUsersResourcePutPutParams, runtime.ClientAuthInfoWriter, ...ClientOption) (*AccountUsersResourcePutPutOK, error)); ok {
		return rf(params, authInfo, opts...)
	}
	if rf, ok := ret.Get(0).(func(*AccountUsersResourcePutPutParams, runtime.ClientAuthInfoWriter, ...ClientOption) *AccountUsersResourcePutPutOK); ok {
		r0 = rf(params, authInfo, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*AccountUsersResourcePutPutOK)
		}
	}

	if rf, ok := ret.Get(1).(func(*AccountUsersResourcePutPutParams, runtime.ClientAuthInfoWriter, ...ClientOption) error); ok {
		r1 = rf(params, authInfo, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// SetTransport provides a mock function with given fields: transport
func (_m *MockClientService) SetTransport(transport runtime.ClientTransport) {
	_m.Called(transport)
}

type mockConstructorTestingTNewMockClientService interface {
	mock.TestingT
	Cleanup(func())
}

// NewMockClientService creates a new instance of MockClientService. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewMockClientService(t mockConstructorTestingTNewMockClientService) *MockClientService {
	mock := &MockClientService{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}