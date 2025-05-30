// Code generated by mockery v2.43.2. DO NOT EDIT.

package mocks

import (
	io "io"
	time "time"

	mock "github.com/stretchr/testify/mock"
)

// ReportController is an autogenerated mock type for the ReportController type
type ReportController struct {
	mock.Mock
}

type ReportController_Expecter struct {
	mock *mock.Mock
}

func (_m *ReportController) EXPECT() *ReportController_Expecter {
	return &ReportController_Expecter{mock: &_m.Mock}
}

// GenerateReport provides a mock function with given fields: startDate, endDate, intervalMinutes, initialBalance, tradesFile, assetsFiles
func (_m *ReportController) GenerateReport(startDate time.Time, endDate time.Time, intervalMinutes int, initialBalance float64, tradesFile io.Reader, assetsFiles map[string]io.Reader) (string, error) {
	ret := _m.Called(startDate, endDate, intervalMinutes, initialBalance, tradesFile, assetsFiles)

	if len(ret) == 0 {
		panic("no return value specified for GenerateReport")
	}

	var r0 string
	var r1 error
	if rf, ok := ret.Get(0).(func(time.Time, time.Time, int, float64, io.Reader, map[string]io.Reader) (string, error)); ok {
		return rf(startDate, endDate, intervalMinutes, initialBalance, tradesFile, assetsFiles)
	}
	if rf, ok := ret.Get(0).(func(time.Time, time.Time, int, float64, io.Reader, map[string]io.Reader) string); ok {
		r0 = rf(startDate, endDate, intervalMinutes, initialBalance, tradesFile, assetsFiles)
	} else {
		r0 = ret.Get(0).(string)
	}

	if rf, ok := ret.Get(1).(func(time.Time, time.Time, int, float64, io.Reader, map[string]io.Reader) error); ok {
		r1 = rf(startDate, endDate, intervalMinutes, initialBalance, tradesFile, assetsFiles)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ReportController_GenerateReport_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GenerateReport'
type ReportController_GenerateReport_Call struct {
	*mock.Call
}

// GenerateReport is a helper method to define mock.On call
//   - startDate time.Time
//   - endDate time.Time
//   - intervalMinutes int
//   - initialBalance float64
//   - tradesFile io.Reader
//   - assetsFiles map[string]io.Reader
func (_e *ReportController_Expecter) GenerateReport(startDate interface{}, endDate interface{}, intervalMinutes interface{}, initialBalance interface{}, tradesFile interface{}, assetsFiles interface{}) *ReportController_GenerateReport_Call {
	return &ReportController_GenerateReport_Call{Call: _e.mock.On("GenerateReport", startDate, endDate, intervalMinutes, initialBalance, tradesFile, assetsFiles)}
}

func (_c *ReportController_GenerateReport_Call) Run(run func(startDate time.Time, endDate time.Time, intervalMinutes int, initialBalance float64, tradesFile io.Reader, assetsFiles map[string]io.Reader)) *ReportController_GenerateReport_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(time.Time), args[1].(time.Time), args[2].(int), args[3].(float64), args[4].(io.Reader), args[5].(map[string]io.Reader))
	})
	return _c
}

func (_c *ReportController_GenerateReport_Call) Return(_a0 string, _a1 error) *ReportController_GenerateReport_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *ReportController_GenerateReport_Call) RunAndReturn(run func(time.Time, time.Time, int, float64, io.Reader, map[string]io.Reader) (string, error)) *ReportController_GenerateReport_Call {
	_c.Call.Return(run)
	return _c
}

// NewReportController creates a new instance of ReportController. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewReportController(t interface {
	mock.TestingT
	Cleanup(func())
}) *ReportController {
	mock := &ReportController{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
