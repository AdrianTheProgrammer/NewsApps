// Code generated by mockery v2.43.2. DO NOT EDIT.

package mocks

import (
	articles "newsapps/internal/features/articles"

	mock "github.com/stretchr/testify/mock"
)

// Queries is an autogenerated mock type for the Queries type
type Queries struct {
	mock.Mock
}

// CreateArticle provides a mock function with given fields: newArticle
func (_m *Queries) CreateArticle(newArticle articles.Article) error {
	ret := _m.Called(newArticle)

	if len(ret) == 0 {
		panic("no return value specified for CreateArticle")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(articles.Article) error); ok {
		r0 = rf(newArticle)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// DeleteArticle provides a mock function with given fields: ID
func (_m *Queries) DeleteArticle(ID uint) error {
	ret := _m.Called(ID)

	if len(ret) == 0 {
		panic("no return value specified for DeleteArticle")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(uint) error); ok {
		r0 = rf(ID)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// ReadArticle provides a mock function with given fields: ID
func (_m *Queries) ReadArticle(ID uint) (articles.Article, error) {
	ret := _m.Called(ID)

	if len(ret) == 0 {
		panic("no return value specified for ReadArticle")
	}

	var r0 articles.Article
	var r1 error
	if rf, ok := ret.Get(0).(func(uint) (articles.Article, error)); ok {
		return rf(ID)
	}
	if rf, ok := ret.Get(0).(func(uint) articles.Article); ok {
		r0 = rf(ID)
	} else {
		r0 = ret.Get(0).(articles.Article)
	}

	if rf, ok := ret.Get(1).(func(uint) error); ok {
		r1 = rf(ID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ShowAllArticles provides a mock function with given fields:
func (_m *Queries) ShowAllArticles() ([]articles.Article, error) {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for ShowAllArticles")
	}

	var r0 []articles.Article
	var r1 error
	if rf, ok := ret.Get(0).(func() ([]articles.Article, error)); ok {
		return rf()
	}
	if rf, ok := ret.Get(0).(func() []articles.Article); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]articles.Article)
		}
	}

	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UpdateArticle provides a mock function with given fields: updatedArticle
func (_m *Queries) UpdateArticle(updatedArticle articles.Article) error {
	ret := _m.Called(updatedArticle)

	if len(ret) == 0 {
		panic("no return value specified for UpdateArticle")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(articles.Article) error); ok {
		r0 = rf(updatedArticle)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// NewQueries creates a new instance of Queries. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewQueries(t interface {
	mock.TestingT
	Cleanup(func())
}) *Queries {
	mock := &Queries{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
