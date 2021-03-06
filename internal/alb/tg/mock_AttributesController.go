// Code generated by mockery v1.0.0. DO NOT EDIT.

package tg

import context "context"
import elbv2 "github.com/aws/aws-sdk-go/service/elbv2"
import mock "github.com/stretchr/testify/mock"

// MockAttributesController is an autogenerated mock type for the AttributesController type
type MockAttributesController struct {
	mock.Mock
}

// Reconcile provides a mock function with given fields: ctx, tgArn, attributes
func (_m *MockAttributesController) Reconcile(ctx context.Context, tgArn string, attributes []*elbv2.TargetGroupAttribute) error {
	ret := _m.Called(ctx, tgArn, attributes)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string, []*elbv2.TargetGroupAttribute) error); ok {
		r0 = rf(ctx, tgArn, attributes)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}
