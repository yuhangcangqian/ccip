// Code generated by mockery v2.43.2. DO NOT EDIT.

package mocks

import (
	types "github.com/smartcontractkit/chainlink/v2/core/services/ccipcapability/types"
	mock "github.com/stretchr/testify/mock"
)

// OracleCreator is an autogenerated mock type for the OracleCreator type
type OracleCreator struct {
	mock.Mock
}

// CreateBootstrapOracle provides a mock function with given fields: config
func (_m *OracleCreator) CreateBootstrapOracle(config types.OCR3ConfigWithMeta) (types.CCIPOracle, error) {
	ret := _m.Called(config)

	if len(ret) == 0 {
		panic("no return value specified for CreateBootstrapOracle")
	}

	var r0 types.CCIPOracle
	var r1 error
	if rf, ok := ret.Get(0).(func(types.OCR3ConfigWithMeta) (types.CCIPOracle, error)); ok {
		return rf(config)
	}
	if rf, ok := ret.Get(0).(func(types.OCR3ConfigWithMeta) types.CCIPOracle); ok {
		r0 = rf(config)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(types.CCIPOracle)
		}
	}

	if rf, ok := ret.Get(1).(func(types.OCR3ConfigWithMeta) error); ok {
		r1 = rf(config)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CreatePluginOracle provides a mock function with given fields: pluginType, config
func (_m *OracleCreator) CreatePluginOracle(pluginType types.PluginType, config types.OCR3ConfigWithMeta) (types.CCIPOracle, error) {
	ret := _m.Called(pluginType, config)

	if len(ret) == 0 {
		panic("no return value specified for CreatePluginOracle")
	}

	var r0 types.CCIPOracle
	var r1 error
	if rf, ok := ret.Get(0).(func(types.PluginType, types.OCR3ConfigWithMeta) (types.CCIPOracle, error)); ok {
		return rf(pluginType, config)
	}
	if rf, ok := ret.Get(0).(func(types.PluginType, types.OCR3ConfigWithMeta) types.CCIPOracle); ok {
		r0 = rf(pluginType, config)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(types.CCIPOracle)
		}
	}

	if rf, ok := ret.Get(1).(func(types.PluginType, types.OCR3ConfigWithMeta) error); ok {
		r1 = rf(pluginType, config)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NewOracleCreator creates a new instance of OracleCreator. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewOracleCreator(t interface {
	mock.TestingT
	Cleanup(func())
}) *OracleCreator {
	mock := &OracleCreator{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
