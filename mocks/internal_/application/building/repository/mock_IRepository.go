// Code generated by mockery v2.38.0. DO NOT EDIT.

package repository

import (
	context "context"

	model "github.com/radityacandra/go-project-mongodb/internal/application/building/model"
	mock "github.com/stretchr/testify/mock"
)

// MockIRepository is an autogenerated mock type for the IRepository type
type MockIRepository struct {
	mock.Mock
}

type MockIRepository_Expecter struct {
	mock *mock.Mock
}

func (_m *MockIRepository) EXPECT() *MockIRepository_Expecter {
	return &MockIRepository_Expecter{mock: &_m.Mock}
}

// Create provides a mock function with given fields: ctx, building
func (_m *MockIRepository) Create(ctx context.Context, building *model.Building) error {
	ret := _m.Called(ctx, building)

	if len(ret) == 0 {
		panic("no return value specified for Create")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *model.Building) error); ok {
		r0 = rf(ctx, building)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockIRepository_Create_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Create'
type MockIRepository_Create_Call struct {
	*mock.Call
}

// Create is a helper method to define mock.On call
//   - ctx context.Context
//   - building *model.Building
func (_e *MockIRepository_Expecter) Create(ctx interface{}, building interface{}) *MockIRepository_Create_Call {
	return &MockIRepository_Create_Call{Call: _e.mock.On("Create", ctx, building)}
}

func (_c *MockIRepository_Create_Call) Run(run func(ctx context.Context, building *model.Building)) *MockIRepository_Create_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(*model.Building))
	})
	return _c
}

func (_c *MockIRepository_Create_Call) Return(_a0 error) *MockIRepository_Create_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockIRepository_Create_Call) RunAndReturn(run func(context.Context, *model.Building) error) *MockIRepository_Create_Call {
	_c.Call.Return(run)
	return _c
}

// FindBuildingByUserId provides a mock function with given fields: ctx, userId, page, pageSize, sortBy, order, keyword
func (_m *MockIRepository) FindBuildingByUserId(ctx context.Context, userId string, page int, pageSize int, sortBy string, order string, keyword *string) ([]model.Building, int) {
	ret := _m.Called(ctx, userId, page, pageSize, sortBy, order, keyword)

	if len(ret) == 0 {
		panic("no return value specified for FindBuildingByUserId")
	}

	var r0 []model.Building
	var r1 int
	if rf, ok := ret.Get(0).(func(context.Context, string, int, int, string, string, *string) ([]model.Building, int)); ok {
		return rf(ctx, userId, page, pageSize, sortBy, order, keyword)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string, int, int, string, string, *string) []model.Building); ok {
		r0 = rf(ctx, userId, page, pageSize, sortBy, order, keyword)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]model.Building)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string, int, int, string, string, *string) int); ok {
		r1 = rf(ctx, userId, page, pageSize, sortBy, order, keyword)
	} else {
		r1 = ret.Get(1).(int)
	}

	return r0, r1
}

// MockIRepository_FindBuildingByUserId_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'FindBuildingByUserId'
type MockIRepository_FindBuildingByUserId_Call struct {
	*mock.Call
}

// FindBuildingByUserId is a helper method to define mock.On call
//   - ctx context.Context
//   - userId string
//   - page int
//   - pageSize int
//   - sortBy string
//   - order string
//   - keyword *string
func (_e *MockIRepository_Expecter) FindBuildingByUserId(ctx interface{}, userId interface{}, page interface{}, pageSize interface{}, sortBy interface{}, order interface{}, keyword interface{}) *MockIRepository_FindBuildingByUserId_Call {
	return &MockIRepository_FindBuildingByUserId_Call{Call: _e.mock.On("FindBuildingByUserId", ctx, userId, page, pageSize, sortBy, order, keyword)}
}

func (_c *MockIRepository_FindBuildingByUserId_Call) Run(run func(ctx context.Context, userId string, page int, pageSize int, sortBy string, order string, keyword *string)) *MockIRepository_FindBuildingByUserId_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string), args[2].(int), args[3].(int), args[4].(string), args[5].(string), args[6].(*string))
	})
	return _c
}

func (_c *MockIRepository_FindBuildingByUserId_Call) Return(_a0 []model.Building, _a1 int) *MockIRepository_FindBuildingByUserId_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockIRepository_FindBuildingByUserId_Call) RunAndReturn(run func(context.Context, string, int, int, string, string, *string) ([]model.Building, int)) *MockIRepository_FindBuildingByUserId_Call {
	_c.Call.Return(run)
	return _c
}

// FindBuildingByUserIdAndName provides a mock function with given fields: ctx, userId, name
func (_m *MockIRepository) FindBuildingByUserIdAndName(ctx context.Context, userId string, name string) *model.Building {
	ret := _m.Called(ctx, userId, name)

	if len(ret) == 0 {
		panic("no return value specified for FindBuildingByUserIdAndName")
	}

	var r0 *model.Building
	if rf, ok := ret.Get(0).(func(context.Context, string, string) *model.Building); ok {
		r0 = rf(ctx, userId, name)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.Building)
		}
	}

	return r0
}

// MockIRepository_FindBuildingByUserIdAndName_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'FindBuildingByUserIdAndName'
type MockIRepository_FindBuildingByUserIdAndName_Call struct {
	*mock.Call
}

// FindBuildingByUserIdAndName is a helper method to define mock.On call
//   - ctx context.Context
//   - userId string
//   - name string
func (_e *MockIRepository_Expecter) FindBuildingByUserIdAndName(ctx interface{}, userId interface{}, name interface{}) *MockIRepository_FindBuildingByUserIdAndName_Call {
	return &MockIRepository_FindBuildingByUserIdAndName_Call{Call: _e.mock.On("FindBuildingByUserIdAndName", ctx, userId, name)}
}

func (_c *MockIRepository_FindBuildingByUserIdAndName_Call) Run(run func(ctx context.Context, userId string, name string)) *MockIRepository_FindBuildingByUserIdAndName_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string), args[2].(string))
	})
	return _c
}

func (_c *MockIRepository_FindBuildingByUserIdAndName_Call) Return(_a0 *model.Building) *MockIRepository_FindBuildingByUserIdAndName_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockIRepository_FindBuildingByUserIdAndName_Call) RunAndReturn(run func(context.Context, string, string) *model.Building) *MockIRepository_FindBuildingByUserIdAndName_Call {
	_c.Call.Return(run)
	return _c
}

// NewMockIRepository creates a new instance of MockIRepository. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockIRepository(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockIRepository {
	mock := &MockIRepository{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}