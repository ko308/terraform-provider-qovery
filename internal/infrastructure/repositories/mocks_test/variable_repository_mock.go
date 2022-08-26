// Code generated by mockery v2.14.0. DO NOT EDIT.

package mocks_test

import (
	context "context"

	variable "github.com/qovery/terraform-provider-qovery/internal/domain/variable"
	mock "github.com/stretchr/testify/mock"
)

// VariableRepository is an autogenerated mock type for the Repository type
type VariableRepository struct {
	mock.Mock
}

type VariableRepository_Expecter struct {
	mock *mock.Mock
}

func (_m *VariableRepository) EXPECT() *VariableRepository_Expecter {
	return &VariableRepository_Expecter{mock: &_m.Mock}
}

// Create provides a mock function with given fields: ctx, scopeResourceID, request
func (_m *VariableRepository) Create(ctx context.Context, scopeResourceID string, request variable.UpsertRequest) (*variable.Variable, error) {
	ret := _m.Called(ctx, scopeResourceID, request)

	var r0 *variable.Variable
	if rf, ok := ret.Get(0).(func(context.Context, string, variable.UpsertRequest) *variable.Variable); ok {
		r0 = rf(ctx, scopeResourceID, request)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*variable.Variable)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string, variable.UpsertRequest) error); ok {
		r1 = rf(ctx, scopeResourceID, request)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// VariableRepository_Create_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Create'
type VariableRepository_Create_Call struct {
	*mock.Call
}

// Create is a helper method to define mock.On call
//   - ctx context.Context
//   - scopeResourceID string
//   - request variable.UpsertRequest
func (_e *VariableRepository_Expecter) Create(ctx interface{}, scopeResourceID interface{}, request interface{}) *VariableRepository_Create_Call {
	return &VariableRepository_Create_Call{Call: _e.mock.On("Create", ctx, scopeResourceID, request)}
}

func (_c *VariableRepository_Create_Call) Run(run func(ctx context.Context, scopeResourceID string, request variable.UpsertRequest)) *VariableRepository_Create_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string), args[2].(variable.UpsertRequest))
	})
	return _c
}

func (_c *VariableRepository_Create_Call) Return(_a0 *variable.Variable, _a1 error) *VariableRepository_Create_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

// Delete provides a mock function with given fields: ctx, scopeResourceID, variableID
func (_m *VariableRepository) Delete(ctx context.Context, scopeResourceID string, variableID string) error {
	ret := _m.Called(ctx, scopeResourceID, variableID)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string, string) error); ok {
		r0 = rf(ctx, scopeResourceID, variableID)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// VariableRepository_Delete_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Delete'
type VariableRepository_Delete_Call struct {
	*mock.Call
}

// Delete is a helper method to define mock.On call
//   - ctx context.Context
//   - scopeResourceID string
//   - variableID string
func (_e *VariableRepository_Expecter) Delete(ctx interface{}, scopeResourceID interface{}, variableID interface{}) *VariableRepository_Delete_Call {
	return &VariableRepository_Delete_Call{Call: _e.mock.On("Delete", ctx, scopeResourceID, variableID)}
}

func (_c *VariableRepository_Delete_Call) Run(run func(ctx context.Context, scopeResourceID string, variableID string)) *VariableRepository_Delete_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string), args[2].(string))
	})
	return _c
}

func (_c *VariableRepository_Delete_Call) Return(_a0 error) *VariableRepository_Delete_Call {
	_c.Call.Return(_a0)
	return _c
}

// List provides a mock function with given fields: ctx, scopeResourceID
func (_m *VariableRepository) List(ctx context.Context, scopeResourceID string) (variable.Variables, error) {
	ret := _m.Called(ctx, scopeResourceID)

	var r0 variable.Variables
	if rf, ok := ret.Get(0).(func(context.Context, string) variable.Variables); ok {
		r0 = rf(ctx, scopeResourceID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(variable.Variables)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, scopeResourceID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// VariableRepository_List_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'List'
type VariableRepository_List_Call struct {
	*mock.Call
}

// List is a helper method to define mock.On call
//   - ctx context.Context
//   - scopeResourceID string
func (_e *VariableRepository_Expecter) List(ctx interface{}, scopeResourceID interface{}) *VariableRepository_List_Call {
	return &VariableRepository_List_Call{Call: _e.mock.On("List", ctx, scopeResourceID)}
}

func (_c *VariableRepository_List_Call) Run(run func(ctx context.Context, scopeResourceID string)) *VariableRepository_List_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string))
	})
	return _c
}

func (_c *VariableRepository_List_Call) Return(_a0 variable.Variables, _a1 error) *VariableRepository_List_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

// Update provides a mock function with given fields: ctx, scopeResourceID, variableID, request
func (_m *VariableRepository) Update(ctx context.Context, scopeResourceID string, variableID string, request variable.UpsertRequest) (*variable.Variable, error) {
	ret := _m.Called(ctx, scopeResourceID, variableID, request)

	var r0 *variable.Variable
	if rf, ok := ret.Get(0).(func(context.Context, string, string, variable.UpsertRequest) *variable.Variable); ok {
		r0 = rf(ctx, scopeResourceID, variableID, request)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*variable.Variable)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string, string, variable.UpsertRequest) error); ok {
		r1 = rf(ctx, scopeResourceID, variableID, request)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// VariableRepository_Update_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Update'
type VariableRepository_Update_Call struct {
	*mock.Call
}

// Update is a helper method to define mock.On call
//   - ctx context.Context
//   - scopeResourceID string
//   - variableID string
//   - request variable.UpsertRequest
func (_e *VariableRepository_Expecter) Update(ctx interface{}, scopeResourceID interface{}, variableID interface{}, request interface{}) *VariableRepository_Update_Call {
	return &VariableRepository_Update_Call{Call: _e.mock.On("Update", ctx, scopeResourceID, variableID, request)}
}

func (_c *VariableRepository_Update_Call) Run(run func(ctx context.Context, scopeResourceID string, variableID string, request variable.UpsertRequest)) *VariableRepository_Update_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string), args[2].(string), args[3].(variable.UpsertRequest))
	})
	return _c
}

func (_c *VariableRepository_Update_Call) Return(_a0 *variable.Variable, _a1 error) *VariableRepository_Update_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

type mockConstructorTestingTNewVariableRepository interface {
	mock.TestingT
	Cleanup(func())
}

// NewVariableRepository creates a new instance of VariableRepository. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewVariableRepository(t mockConstructorTestingTNewVariableRepository) *VariableRepository {
	mock := &VariableRepository{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
