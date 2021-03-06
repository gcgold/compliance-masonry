package mocks

import (
	"github.com/opencontrol/compliance-masonry/pkg/lib/common"
	"github.com/stretchr/testify/mock"
)

// SchemaParser is an autogenerated mock type for the SchemaParser type
type SchemaParser struct {
	mock.Mock
}

// Parse provides a mock function with given fields: data
func (_m *SchemaParser) Parse(data []byte) (common.OpenControl, error) {
	ret := _m.Called(data)

	var r0 common.OpenControl
	if rf, ok := ret.Get(0).(func([]byte) common.OpenControl); ok {
		r0 = rf(data)
	} else {
		r0 = ret.Get(0).(common.OpenControl)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func([]byte) error); ok {
		r1 = rf(data)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
