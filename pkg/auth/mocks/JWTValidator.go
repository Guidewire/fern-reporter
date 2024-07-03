// Code generated by mockery v2.43.0. DO NOT EDIT.

package mocks

import (
	context "context"

	jwk "github.com/lestrrat-go/jwx/v2/jwk"
	jwt "github.com/lestrrat-go/jwx/v2/jwt"

	mock "github.com/stretchr/testify/mock"
)

// JWTValidator is an autogenerated mock type for the JWTValidator type
type JWTValidator struct {
	mock.Mock
}

// ParseAndValidateToken provides a mock function with given fields: ctx, tokenString, set
func (_m *JWTValidator) ParseAndValidateToken(ctx context.Context, tokenString string, set jwk.Set) (jwt.Token, error) {
	ret := _m.Called(ctx, tokenString, set)

	if len(ret) == 0 {
		panic("no return value specified for ParseAndValidateToken")
	}

	var r0 jwt.Token
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string, jwk.Set) (jwt.Token, error)); ok {
		return rf(ctx, tokenString, set)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string, jwk.Set) jwt.Token); ok {
		r0 = rf(ctx, tokenString, set)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(jwt.Token)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string, jwk.Set) error); ok {
		r1 = rf(ctx, tokenString, set)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NewJWTValidator creates a new instance of JWTValidator. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewJWTValidator(t interface {
	mock.TestingT
	Cleanup(func())
}) *JWTValidator {
	mock := &JWTValidator{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
