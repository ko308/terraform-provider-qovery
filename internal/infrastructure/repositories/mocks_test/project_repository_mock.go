// Code generated by mockery v2.14.0. DO NOT EDIT.

package mocks_test

import (
	context "context"

	project "github.com/qovery/terraform-provider-qovery/internal/domain/project"
	mock "github.com/stretchr/testify/mock"
)

// ProjectRepository is an autogenerated mock type for the Repository type
type ProjectRepository struct {
	mock.Mock
}

type ProjectRepository_Expecter struct {
	mock *mock.Mock
}

func (_m *ProjectRepository) EXPECT() *ProjectRepository_Expecter {
	return &ProjectRepository_Expecter{mock: &_m.Mock}
}

// Create provides a mock function with given fields: ctx, organizationID, request
func (_m *ProjectRepository) Create(ctx context.Context, organizationID string, request project.UpsertRepositoryRequest) (*project.Project, error) {
	ret := _m.Called(ctx, organizationID, request)

	var r0 *project.Project
	if rf, ok := ret.Get(0).(func(context.Context, string, project.UpsertRepositoryRequest) *project.Project); ok {
		r0 = rf(ctx, organizationID, request)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*project.Project)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string, project.UpsertRepositoryRequest) error); ok {
		r1 = rf(ctx, organizationID, request)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ProjectRepository_Create_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Create'
type ProjectRepository_Create_Call struct {
	*mock.Call
}

// Create is a helper method to define mock.On call
//   - ctx context.Context
//   - organizationID string
//   - request project.UpsertRepositoryRequest
func (_e *ProjectRepository_Expecter) Create(ctx interface{}, organizationID interface{}, request interface{}) *ProjectRepository_Create_Call {
	return &ProjectRepository_Create_Call{Call: _e.mock.On("Create", ctx, organizationID, request)}
}

func (_c *ProjectRepository_Create_Call) Run(run func(ctx context.Context, organizationID string, request project.UpsertRepositoryRequest)) *ProjectRepository_Create_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string), args[2].(project.UpsertRepositoryRequest))
	})
	return _c
}

func (_c *ProjectRepository_Create_Call) Return(_a0 *project.Project, _a1 error) *ProjectRepository_Create_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

// Delete provides a mock function with given fields: ctx, projectID
func (_m *ProjectRepository) Delete(ctx context.Context, projectID string) error {
	ret := _m.Called(ctx, projectID)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string) error); ok {
		r0 = rf(ctx, projectID)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// ProjectRepository_Delete_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Delete'
type ProjectRepository_Delete_Call struct {
	*mock.Call
}

// Delete is a helper method to define mock.On call
//   - ctx context.Context
//   - projectID string
func (_e *ProjectRepository_Expecter) Delete(ctx interface{}, projectID interface{}) *ProjectRepository_Delete_Call {
	return &ProjectRepository_Delete_Call{Call: _e.mock.On("Delete", ctx, projectID)}
}

func (_c *ProjectRepository_Delete_Call) Run(run func(ctx context.Context, projectID string)) *ProjectRepository_Delete_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string))
	})
	return _c
}

func (_c *ProjectRepository_Delete_Call) Return(_a0 error) *ProjectRepository_Delete_Call {
	_c.Call.Return(_a0)
	return _c
}

// Get provides a mock function with given fields: ctx, projectID
func (_m *ProjectRepository) Get(ctx context.Context, projectID string) (*project.Project, error) {
	ret := _m.Called(ctx, projectID)

	var r0 *project.Project
	if rf, ok := ret.Get(0).(func(context.Context, string) *project.Project); ok {
		r0 = rf(ctx, projectID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*project.Project)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, projectID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ProjectRepository_Get_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Get'
type ProjectRepository_Get_Call struct {
	*mock.Call
}

// Get is a helper method to define mock.On call
//   - ctx context.Context
//   - projectID string
func (_e *ProjectRepository_Expecter) Get(ctx interface{}, projectID interface{}) *ProjectRepository_Get_Call {
	return &ProjectRepository_Get_Call{Call: _e.mock.On("Get", ctx, projectID)}
}

func (_c *ProjectRepository_Get_Call) Run(run func(ctx context.Context, projectID string)) *ProjectRepository_Get_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string))
	})
	return _c
}

func (_c *ProjectRepository_Get_Call) Return(_a0 *project.Project, _a1 error) *ProjectRepository_Get_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

// Update provides a mock function with given fields: ctx, projectID, request
func (_m *ProjectRepository) Update(ctx context.Context, projectID string, request project.UpsertRepositoryRequest) (*project.Project, error) {
	ret := _m.Called(ctx, projectID, request)

	var r0 *project.Project
	if rf, ok := ret.Get(0).(func(context.Context, string, project.UpsertRepositoryRequest) *project.Project); ok {
		r0 = rf(ctx, projectID, request)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*project.Project)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string, project.UpsertRepositoryRequest) error); ok {
		r1 = rf(ctx, projectID, request)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ProjectRepository_Update_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Update'
type ProjectRepository_Update_Call struct {
	*mock.Call
}

// Update is a helper method to define mock.On call
//   - ctx context.Context
//   - projectID string
//   - request project.UpsertRepositoryRequest
func (_e *ProjectRepository_Expecter) Update(ctx interface{}, projectID interface{}, request interface{}) *ProjectRepository_Update_Call {
	return &ProjectRepository_Update_Call{Call: _e.mock.On("Update", ctx, projectID, request)}
}

func (_c *ProjectRepository_Update_Call) Run(run func(ctx context.Context, projectID string, request project.UpsertRepositoryRequest)) *ProjectRepository_Update_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string), args[2].(project.UpsertRepositoryRequest))
	})
	return _c
}

func (_c *ProjectRepository_Update_Call) Return(_a0 *project.Project, _a1 error) *ProjectRepository_Update_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

type mockConstructorTestingTNewProjectRepository interface {
	mock.TestingT
	Cleanup(func())
}

// NewProjectRepository creates a new instance of ProjectRepository. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewProjectRepository(t mockConstructorTestingTNewProjectRepository) *ProjectRepository {
	mock := &ProjectRepository{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
