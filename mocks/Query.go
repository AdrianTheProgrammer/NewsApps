// Code generated by mockery v2.43.2. DO NOT EDIT.

package mocks

import mock "github.com/stretchr/testify/mock"

// Query is an autogenerated mock type for the Query type
type Query struct {
	mock.Mock
}

// NewQuery creates a new instance of Query. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewQuery(t interface {
	mock.TestingT
	Cleanup(func())
}) *Query {
	mock := &Query{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
