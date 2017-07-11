// Copyright 2017 Northern.tech AS
//
//    Licensed under the Apache License, Version 2.0 (the "License");
//    you may not use this file except in compliance with the License.
//    You may obtain a copy of the License at
//
//        http://www.apache.org/licenses/LICENSE-2.0
//
//    Unless required by applicable law or agreed to in writing, software
//    distributed under the License is distributed on an "AS IS" BASIS,
//    WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//    See the License for the specific language governing permissions and
//    limitations under the License.
// Code generated by mockery v1.0.0
package mocks

import context "context"
import mock "github.com/stretchr/testify/mock"
import model "github.com/mendersoftware/useradm/model"

// DataStore is an autogenerated mock type for the DataStore type
type DataStore struct {
	mock.Mock
}

// CreateUser provides a mock function with given fields: ctx, u
func (_m *DataStore) CreateUser(ctx context.Context, u *model.User) error {
	ret := _m.Called(ctx, u)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *model.User) error); ok {
		r0 = rf(ctx, u)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// DeleteUser provides a mock function with given fields: ctx, id
func (_m *DataStore) DeleteUser(ctx context.Context, id string) error {
	ret := _m.Called(ctx, id)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string) error); ok {
		r0 = rf(ctx, id)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// GetUserByEmail provides a mock function with given fields: ctx, email
func (_m *DataStore) GetUserByEmail(ctx context.Context, email string) (*model.User, error) {
	ret := _m.Called(ctx, email)

	var r0 *model.User
	if rf, ok := ret.Get(0).(func(context.Context, string) *model.User); ok {
		r0 = rf(ctx, email)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.User)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, email)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetUserById provides a mock function with given fields: ctx, id
func (_m *DataStore) GetUserById(ctx context.Context, id string) (*model.User, error) {
	ret := _m.Called(ctx, id)

	var r0 *model.User
	if rf, ok := ret.Get(0).(func(context.Context, string) *model.User); ok {
		r0 = rf(ctx, id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.User)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetUsers provides a mock function with given fields: ctx
func (_m *DataStore) GetUsers(ctx context.Context) ([]model.User, error) {
	ret := _m.Called(ctx)

	var r0 []model.User
	if rf, ok := ret.Get(0).(func(context.Context) []model.User); ok {
		r0 = rf(ctx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]model.User)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(ctx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UpdateUser provides a mock function with given fields: ctx, id, u
func (_m *DataStore) UpdateUser(ctx context.Context, id string, u *model.UserUpdate) error {
	ret := _m.Called(ctx, id, u)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string, *model.UserUpdate) error); ok {
		r0 = rf(ctx, id, u)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}
