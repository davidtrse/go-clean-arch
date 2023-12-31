// Code generated by mockery v2.33.2. DO NOT EDIT.

package mocks

import (
	context "context"

	domain "github.com/davidtrse/go-clean-arch/internal/domain"
	mock "github.com/stretchr/testify/mock"
)

// AccountRepository is an autogenerated mock type for the AccountRepository type
type AccountRepository struct {
	mock.Mock
}

type AccountRepository_Expecter struct {
	mock *mock.Mock
}

func (_m *AccountRepository) EXPECT() *AccountRepository_Expecter {
	return &AccountRepository_Expecter{mock: &_m.Mock}
}

// Create provides a mock function with given fields: ctx, acc
func (_m *AccountRepository) Create(ctx context.Context, acc *domain.Account) error {
	ret := _m.Called(ctx, acc)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *domain.Account) error); ok {
		r0 = rf(ctx, acc)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// AccountRepository_Create_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Create'
type AccountRepository_Create_Call struct {
	*mock.Call
}

// Create is a helper method to define mock.On call
//   - ctx context.Context
//   - acc *domain.Account
func (_e *AccountRepository_Expecter) Create(ctx interface{}, acc interface{}) *AccountRepository_Create_Call {
	return &AccountRepository_Create_Call{Call: _e.mock.On("Create", ctx, acc)}
}

func (_c *AccountRepository_Create_Call) Run(run func(ctx context.Context, acc *domain.Account)) *AccountRepository_Create_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(*domain.Account))
	})
	return _c
}

func (_c *AccountRepository_Create_Call) Return(_a0 error) *AccountRepository_Create_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *AccountRepository_Create_Call) RunAndReturn(run func(context.Context, *domain.Account) error) *AccountRepository_Create_Call {
	_c.Call.Return(run)
	return _c
}

// Get provides a mock function with given fields: ctx
func (_m *AccountRepository) Get(ctx context.Context) ([]*domain.Account, error) {
	ret := _m.Called(ctx)

	var r0 []*domain.Account
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context) ([]*domain.Account, error)); ok {
		return rf(ctx)
	}
	if rf, ok := ret.Get(0).(func(context.Context) []*domain.Account); ok {
		r0 = rf(ctx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*domain.Account)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(ctx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// AccountRepository_Get_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Get'
type AccountRepository_Get_Call struct {
	*mock.Call
}

// Get is a helper method to define mock.On call
//   - ctx context.Context
func (_e *AccountRepository_Expecter) Get(ctx interface{}) *AccountRepository_Get_Call {
	return &AccountRepository_Get_Call{Call: _e.mock.On("Get", ctx)}
}

func (_c *AccountRepository_Get_Call) Run(run func(ctx context.Context)) *AccountRepository_Get_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context))
	})
	return _c
}

func (_c *AccountRepository_Get_Call) Return(_a0 []*domain.Account, _a1 error) *AccountRepository_Get_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *AccountRepository_Get_Call) RunAndReturn(run func(context.Context) ([]*domain.Account, error)) *AccountRepository_Get_Call {
	_c.Call.Return(run)
	return _c
}

// NewAccountRepository creates a new instance of AccountRepository. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewAccountRepository(t interface {
	mock.TestingT
	Cleanup(func())
}) *AccountRepository {
	mock := &AccountRepository{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
