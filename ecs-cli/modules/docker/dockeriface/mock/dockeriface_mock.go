// Copyright 2015-2017 Amazon.com, Inc. or its affiliates. All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License"). You may
// not use this file except in compliance with the License. A copy of the
// License is located at
//
//     http://aws.amazon.com/apache2.0/
//
// or in the "license" file accompanying this file. This file is distributed
// on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either
// express or implied. See the License for the specific language governing
// permissions and limitations under the License.

// Automatically generated by MockGen. DO NOT EDIT!
// Source: github.com/aws/amazon-ecs-cli/ecs-cli/modules/docker/dockeriface (interfaces: DockerAPI)

package mock_dockeriface

import (
	go_dockerclient "github.com/fsouza/go-dockerclient"
	gomock "github.com/golang/mock/gomock"
)

// Mock of DockerAPI interface
type MockDockerAPI struct {
	ctrl     *gomock.Controller
	recorder *_MockDockerAPIRecorder
}

// Recorder for MockDockerAPI (not exported)
type _MockDockerAPIRecorder struct {
	mock *MockDockerAPI
}

func NewMockDockerAPI(ctrl *gomock.Controller) *MockDockerAPI {
	mock := &MockDockerAPI{ctrl: ctrl}
	mock.recorder = &_MockDockerAPIRecorder{mock}
	return mock
}

func (_m *MockDockerAPI) EXPECT() *_MockDockerAPIRecorder {
	return _m.recorder
}

func (_m *MockDockerAPI) PullImage(_param0 go_dockerclient.PullImageOptions, _param1 go_dockerclient.AuthConfiguration) error {
	ret := _m.ctrl.Call(_m, "PullImage", _param0, _param1)
	ret0, _ := ret[0].(error)
	return ret0
}

func (_mr *_MockDockerAPIRecorder) PullImage(arg0, arg1 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "PullImage", arg0, arg1)
}

func (_m *MockDockerAPI) PushImage(_param0 go_dockerclient.PushImageOptions, _param1 go_dockerclient.AuthConfiguration) error {
	ret := _m.ctrl.Call(_m, "PushImage", _param0, _param1)
	ret0, _ := ret[0].(error)
	return ret0
}

func (_mr *_MockDockerAPIRecorder) PushImage(arg0, arg1 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "PushImage", arg0, arg1)
}

func (_m *MockDockerAPI) TagImage(_param0 string, _param1 go_dockerclient.TagImageOptions) error {
	ret := _m.ctrl.Call(_m, "TagImage", _param0, _param1)
	ret0, _ := ret[0].(error)
	return ret0
}

func (_mr *_MockDockerAPIRecorder) TagImage(arg0, arg1 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "TagImage", arg0, arg1)
}
