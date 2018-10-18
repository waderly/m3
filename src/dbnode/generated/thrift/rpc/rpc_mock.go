// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/m3db/m3/src/dbnode/generated/thrift/rpc/tchan-go

// Copyright (c) 2018 Uber Technologies, Inc.
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
// THE SOFTWARE.

// Package rpc is a generated GoMock package.
package rpc

import (
	"reflect"

	"github.com/golang/mock/gomock"
	"github.com/uber/tchannel-go/thrift"
)

// MockTChanCluster is a mock of TChanCluster interface
type MockTChanCluster struct {
	ctrl     *gomock.Controller
	recorder *MockTChanClusterMockRecorder
}

// MockTChanClusterMockRecorder is the mock recorder for MockTChanCluster
type MockTChanClusterMockRecorder struct {
	mock *MockTChanCluster
}

// NewMockTChanCluster creates a new mock instance
func NewMockTChanCluster(ctrl *gomock.Controller) *MockTChanCluster {
	mock := &MockTChanCluster{ctrl: ctrl}
	mock.recorder = &MockTChanClusterMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockTChanCluster) EXPECT() *MockTChanClusterMockRecorder {
	return m.recorder
}

// Fetch mocks base method
func (m *MockTChanCluster) Fetch(ctx thrift.Context, req *FetchRequest) (*FetchResult_, error) {
	ret := m.ctrl.Call(m, "Fetch", ctx, req)
	ret0, _ := ret[0].(*FetchResult_)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Fetch indicates an expected call of Fetch
func (mr *MockTChanClusterMockRecorder) Fetch(ctx, req interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Fetch", reflect.TypeOf((*MockTChanCluster)(nil).Fetch), ctx, req)
}

// FetchTagged mocks base method
func (m *MockTChanCluster) FetchTagged(ctx thrift.Context, req *FetchTaggedRequest) (*FetchTaggedResult_, error) {
	ret := m.ctrl.Call(m, "FetchTagged", ctx, req)
	ret0, _ := ret[0].(*FetchTaggedResult_)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FetchTagged indicates an expected call of FetchTagged
func (mr *MockTChanClusterMockRecorder) FetchTagged(ctx, req interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FetchTagged", reflect.TypeOf((*MockTChanCluster)(nil).FetchTagged), ctx, req)
}

// Health mocks base method
func (m *MockTChanCluster) Health(ctx thrift.Context) (*HealthResult_, error) {
	ret := m.ctrl.Call(m, "Health", ctx)
	ret0, _ := ret[0].(*HealthResult_)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Health indicates an expected call of Health
func (mr *MockTChanClusterMockRecorder) Health(ctx interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Health", reflect.TypeOf((*MockTChanCluster)(nil).Health), ctx)
}

// Query mocks base method
func (m *MockTChanCluster) Query(ctx thrift.Context, req *QueryRequest) (*QueryResult_, error) {
	ret := m.ctrl.Call(m, "Query", ctx, req)
	ret0, _ := ret[0].(*QueryResult_)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Query indicates an expected call of Query
func (mr *MockTChanClusterMockRecorder) Query(ctx, req interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Query", reflect.TypeOf((*MockTChanCluster)(nil).Query), ctx, req)
}

// Truncate mocks base method
func (m *MockTChanCluster) Truncate(ctx thrift.Context, req *TruncateRequest) (*TruncateResult_, error) {
	ret := m.ctrl.Call(m, "Truncate", ctx, req)
	ret0, _ := ret[0].(*TruncateResult_)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Truncate indicates an expected call of Truncate
func (mr *MockTChanClusterMockRecorder) Truncate(ctx, req interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Truncate", reflect.TypeOf((*MockTChanCluster)(nil).Truncate), ctx, req)
}

// Write mocks base method
func (m *MockTChanCluster) Write(ctx thrift.Context, req *WriteRequest) error {
	ret := m.ctrl.Call(m, "Write", ctx, req)
	ret0, _ := ret[0].(error)
	return ret0
}

// Write indicates an expected call of Write
func (mr *MockTChanClusterMockRecorder) Write(ctx, req interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Write", reflect.TypeOf((*MockTChanCluster)(nil).Write), ctx, req)
}

// WriteTagged mocks base method
func (m *MockTChanCluster) WriteTagged(ctx thrift.Context, req *WriteTaggedRequest) error {
	ret := m.ctrl.Call(m, "WriteTagged", ctx, req)
	ret0, _ := ret[0].(error)
	return ret0
}

// WriteTagged indicates an expected call of WriteTagged
func (mr *MockTChanClusterMockRecorder) WriteTagged(ctx, req interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "WriteTagged", reflect.TypeOf((*MockTChanCluster)(nil).WriteTagged), ctx, req)
}

// MockTChanNode is a mock of TChanNode interface
type MockTChanNode struct {
	ctrl     *gomock.Controller
	recorder *MockTChanNodeMockRecorder
}

// MockTChanNodeMockRecorder is the mock recorder for MockTChanNode
type MockTChanNodeMockRecorder struct {
	mock *MockTChanNode
}

// NewMockTChanNode creates a new mock instance
func NewMockTChanNode(ctrl *gomock.Controller) *MockTChanNode {
	mock := &MockTChanNode{ctrl: ctrl}
	mock.recorder = &MockTChanNodeMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockTChanNode) EXPECT() *MockTChanNodeMockRecorder {
	return m.recorder
}

// Bootstrapped mocks base method
func (m *MockTChanNode) Bootstrapped(ctx thrift.Context) (*NodeBootstrappedResult_, error) {
	ret := m.ctrl.Call(m, "Bootstrapped", ctx)
	ret0, _ := ret[0].(*NodeBootstrappedResult_)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Bootstrapped indicates an expected call of Bootstrapped
func (mr *MockTChanNodeMockRecorder) Bootstrapped(ctx interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Bootstrapped", reflect.TypeOf((*MockTChanNode)(nil).Bootstrapped), ctx)
}

// Fetch mocks base method
func (m *MockTChanNode) Fetch(ctx thrift.Context, req *FetchRequest) (*FetchResult_, error) {
	ret := m.ctrl.Call(m, "Fetch", ctx, req)
	ret0, _ := ret[0].(*FetchResult_)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Fetch indicates an expected call of Fetch
func (mr *MockTChanNodeMockRecorder) Fetch(ctx, req interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Fetch", reflect.TypeOf((*MockTChanNode)(nil).Fetch), ctx, req)
}

// FetchBatchRaw mocks base method
func (m *MockTChanNode) FetchBatchRaw(ctx thrift.Context, req *FetchBatchRawRequest) (*FetchBatchRawResult_, error) {
	ret := m.ctrl.Call(m, "FetchBatchRaw", ctx, req)
	ret0, _ := ret[0].(*FetchBatchRawResult_)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FetchBatchRaw indicates an expected call of FetchBatchRaw
func (mr *MockTChanNodeMockRecorder) FetchBatchRaw(ctx, req interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FetchBatchRaw", reflect.TypeOf((*MockTChanNode)(nil).FetchBatchRaw), ctx, req)
}

// FetchBlocksMetadataRawV2 mocks base method
func (m *MockTChanNode) FetchBlocksMetadataRawV2(ctx thrift.Context, req *FetchBlocksMetadataRawV2Request) (*FetchBlocksMetadataRawV2Result_, error) {
	ret := m.ctrl.Call(m, "FetchBlocksMetadataRawV2", ctx, req)
	ret0, _ := ret[0].(*FetchBlocksMetadataRawV2Result_)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FetchBlocksMetadataRawV2 indicates an expected call of FetchBlocksMetadataRawV2
func (mr *MockTChanNodeMockRecorder) FetchBlocksMetadataRawV2(ctx, req interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FetchBlocksMetadataRawV2", reflect.TypeOf((*MockTChanNode)(nil).FetchBlocksMetadataRawV2), ctx, req)
}

// FetchBlocksRaw mocks base method
func (m *MockTChanNode) FetchBlocksRaw(ctx thrift.Context, req *FetchBlocksRawRequest) (*FetchBlocksRawResult_, error) {
	ret := m.ctrl.Call(m, "FetchBlocksRaw", ctx, req)
	ret0, _ := ret[0].(*FetchBlocksRawResult_)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FetchBlocksRaw indicates an expected call of FetchBlocksRaw
func (mr *MockTChanNodeMockRecorder) FetchBlocksRaw(ctx, req interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FetchBlocksRaw", reflect.TypeOf((*MockTChanNode)(nil).FetchBlocksRaw), ctx, req)
}

// FetchTagged mocks base method
func (m *MockTChanNode) FetchTagged(ctx thrift.Context, req *FetchTaggedRequest) (*FetchTaggedResult_, error) {
	ret := m.ctrl.Call(m, "FetchTagged", ctx, req)
	ret0, _ := ret[0].(*FetchTaggedResult_)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FetchTagged indicates an expected call of FetchTagged
func (mr *MockTChanNodeMockRecorder) FetchTagged(ctx, req interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FetchTagged", reflect.TypeOf((*MockTChanNode)(nil).FetchTagged), ctx, req)
}

// GetPersistRateLimit mocks base method
func (m *MockTChanNode) GetPersistRateLimit(ctx thrift.Context) (*NodePersistRateLimitResult_, error) {
	ret := m.ctrl.Call(m, "GetPersistRateLimit", ctx)
	ret0, _ := ret[0].(*NodePersistRateLimitResult_)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetPersistRateLimit indicates an expected call of GetPersistRateLimit
func (mr *MockTChanNodeMockRecorder) GetPersistRateLimit(ctx interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetPersistRateLimit", reflect.TypeOf((*MockTChanNode)(nil).GetPersistRateLimit), ctx)
}

// GetWriteNewSeriesAsync mocks base method
func (m *MockTChanNode) GetWriteNewSeriesAsync(ctx thrift.Context) (*NodeWriteNewSeriesAsyncResult_, error) {
	ret := m.ctrl.Call(m, "GetWriteNewSeriesAsync", ctx)
	ret0, _ := ret[0].(*NodeWriteNewSeriesAsyncResult_)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetWriteNewSeriesAsync indicates an expected call of GetWriteNewSeriesAsync
func (mr *MockTChanNodeMockRecorder) GetWriteNewSeriesAsync(ctx interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetWriteNewSeriesAsync", reflect.TypeOf((*MockTChanNode)(nil).GetWriteNewSeriesAsync), ctx)
}

// GetWriteNewSeriesBackoffDuration mocks base method
func (m *MockTChanNode) GetWriteNewSeriesBackoffDuration(ctx thrift.Context) (*NodeWriteNewSeriesBackoffDurationResult_, error) {
	ret := m.ctrl.Call(m, "GetWriteNewSeriesBackoffDuration", ctx)
	ret0, _ := ret[0].(*NodeWriteNewSeriesBackoffDurationResult_)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetWriteNewSeriesBackoffDuration indicates an expected call of GetWriteNewSeriesBackoffDuration
func (mr *MockTChanNodeMockRecorder) GetWriteNewSeriesBackoffDuration(ctx interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetWriteNewSeriesBackoffDuration", reflect.TypeOf((*MockTChanNode)(nil).GetWriteNewSeriesBackoffDuration), ctx)
}

// GetWriteNewSeriesLimitPerShardPerSecond mocks base method
func (m *MockTChanNode) GetWriteNewSeriesLimitPerShardPerSecond(ctx thrift.Context) (*NodeWriteNewSeriesLimitPerShardPerSecondResult_, error) {
	ret := m.ctrl.Call(m, "GetWriteNewSeriesLimitPerShardPerSecond", ctx)
	ret0, _ := ret[0].(*NodeWriteNewSeriesLimitPerShardPerSecondResult_)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetWriteNewSeriesLimitPerShardPerSecond indicates an expected call of GetWriteNewSeriesLimitPerShardPerSecond
func (mr *MockTChanNodeMockRecorder) GetWriteNewSeriesLimitPerShardPerSecond(ctx interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetWriteNewSeriesLimitPerShardPerSecond", reflect.TypeOf((*MockTChanNode)(nil).GetWriteNewSeriesLimitPerShardPerSecond), ctx)
}

// Health mocks base method
func (m *MockTChanNode) Health(ctx thrift.Context) (*NodeHealthResult_, error) {
	ret := m.ctrl.Call(m, "Health", ctx)
	ret0, _ := ret[0].(*NodeHealthResult_)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Health indicates an expected call of Health
func (mr *MockTChanNodeMockRecorder) Health(ctx interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Health", reflect.TypeOf((*MockTChanNode)(nil).Health), ctx)
}

// Query mocks base method
func (m *MockTChanNode) Query(ctx thrift.Context, req *QueryRequest) (*QueryResult_, error) {
	ret := m.ctrl.Call(m, "Query", ctx, req)
	ret0, _ := ret[0].(*QueryResult_)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Query indicates an expected call of Query
func (mr *MockTChanNodeMockRecorder) Query(ctx, req interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Query", reflect.TypeOf((*MockTChanNode)(nil).Query), ctx, req)
}

// Repair mocks base method
func (m *MockTChanNode) Repair(ctx thrift.Context) error {
	ret := m.ctrl.Call(m, "Repair", ctx)
	ret0, _ := ret[0].(error)
	return ret0
}

// Repair indicates an expected call of Repair
func (mr *MockTChanNodeMockRecorder) Repair(ctx interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Repair", reflect.TypeOf((*MockTChanNode)(nil).Repair), ctx)
}

// SetPersistRateLimit mocks base method
func (m *MockTChanNode) SetPersistRateLimit(ctx thrift.Context, req *NodeSetPersistRateLimitRequest) (*NodePersistRateLimitResult_, error) {
	ret := m.ctrl.Call(m, "SetPersistRateLimit", ctx, req)
	ret0, _ := ret[0].(*NodePersistRateLimitResult_)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SetPersistRateLimit indicates an expected call of SetPersistRateLimit
func (mr *MockTChanNodeMockRecorder) SetPersistRateLimit(ctx, req interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetPersistRateLimit", reflect.TypeOf((*MockTChanNode)(nil).SetPersistRateLimit), ctx, req)
}

// SetWriteNewSeriesAsync mocks base method
func (m *MockTChanNode) SetWriteNewSeriesAsync(ctx thrift.Context, req *NodeSetWriteNewSeriesAsyncRequest) (*NodeWriteNewSeriesAsyncResult_, error) {
	ret := m.ctrl.Call(m, "SetWriteNewSeriesAsync", ctx, req)
	ret0, _ := ret[0].(*NodeWriteNewSeriesAsyncResult_)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SetWriteNewSeriesAsync indicates an expected call of SetWriteNewSeriesAsync
func (mr *MockTChanNodeMockRecorder) SetWriteNewSeriesAsync(ctx, req interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetWriteNewSeriesAsync", reflect.TypeOf((*MockTChanNode)(nil).SetWriteNewSeriesAsync), ctx, req)
}

// SetWriteNewSeriesBackoffDuration mocks base method
func (m *MockTChanNode) SetWriteNewSeriesBackoffDuration(ctx thrift.Context, req *NodeSetWriteNewSeriesBackoffDurationRequest) (*NodeWriteNewSeriesBackoffDurationResult_, error) {
	ret := m.ctrl.Call(m, "SetWriteNewSeriesBackoffDuration", ctx, req)
	ret0, _ := ret[0].(*NodeWriteNewSeriesBackoffDurationResult_)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SetWriteNewSeriesBackoffDuration indicates an expected call of SetWriteNewSeriesBackoffDuration
func (mr *MockTChanNodeMockRecorder) SetWriteNewSeriesBackoffDuration(ctx, req interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetWriteNewSeriesBackoffDuration", reflect.TypeOf((*MockTChanNode)(nil).SetWriteNewSeriesBackoffDuration), ctx, req)
}

// SetWriteNewSeriesLimitPerShardPerSecond mocks base method
func (m *MockTChanNode) SetWriteNewSeriesLimitPerShardPerSecond(ctx thrift.Context, req *NodeSetWriteNewSeriesLimitPerShardPerSecondRequest) (*NodeWriteNewSeriesLimitPerShardPerSecondResult_, error) {
	ret := m.ctrl.Call(m, "SetWriteNewSeriesLimitPerShardPerSecond", ctx, req)
	ret0, _ := ret[0].(*NodeWriteNewSeriesLimitPerShardPerSecondResult_)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SetWriteNewSeriesLimitPerShardPerSecond indicates an expected call of SetWriteNewSeriesLimitPerShardPerSecond
func (mr *MockTChanNodeMockRecorder) SetWriteNewSeriesLimitPerShardPerSecond(ctx, req interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetWriteNewSeriesLimitPerShardPerSecond", reflect.TypeOf((*MockTChanNode)(nil).SetWriteNewSeriesLimitPerShardPerSecond), ctx, req)
}

// Truncate mocks base method
func (m *MockTChanNode) Truncate(ctx thrift.Context, req *TruncateRequest) (*TruncateResult_, error) {
	ret := m.ctrl.Call(m, "Truncate", ctx, req)
	ret0, _ := ret[0].(*TruncateResult_)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Truncate indicates an expected call of Truncate
func (mr *MockTChanNodeMockRecorder) Truncate(ctx, req interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Truncate", reflect.TypeOf((*MockTChanNode)(nil).Truncate), ctx, req)
}

// Write mocks base method
func (m *MockTChanNode) Write(ctx thrift.Context, req *WriteRequest) error {
	ret := m.ctrl.Call(m, "Write", ctx, req)
	ret0, _ := ret[0].(error)
	return ret0
}

// Write indicates an expected call of Write
func (mr *MockTChanNodeMockRecorder) Write(ctx, req interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Write", reflect.TypeOf((*MockTChanNode)(nil).Write), ctx, req)
}

// WriteBatchRaw mocks base method
func (m *MockTChanNode) WriteBatchRaw(ctx thrift.Context, req *WriteBatchRawRequest) error {
	ret := m.ctrl.Call(m, "WriteBatchRaw", ctx, req)
	ret0, _ := ret[0].(error)
	return ret0
}

// WriteBatchRaw indicates an expected call of WriteBatchRaw
func (mr *MockTChanNodeMockRecorder) WriteBatchRaw(ctx, req interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "WriteBatchRaw", reflect.TypeOf((*MockTChanNode)(nil).WriteBatchRaw), ctx, req)
}

// WriteTagged mocks base method
func (m *MockTChanNode) WriteTagged(ctx thrift.Context, req *WriteTaggedRequest) error {
	ret := m.ctrl.Call(m, "WriteTagged", ctx, req)
	ret0, _ := ret[0].(error)
	return ret0
}

// WriteTagged indicates an expected call of WriteTagged
func (mr *MockTChanNodeMockRecorder) WriteTagged(ctx, req interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "WriteTagged", reflect.TypeOf((*MockTChanNode)(nil).WriteTagged), ctx, req)
}

// WriteTaggedBatchRaw mocks base method
func (m *MockTChanNode) WriteTaggedBatchRaw(ctx thrift.Context, req *WriteTaggedBatchRawRequest) error {
	ret := m.ctrl.Call(m, "WriteTaggedBatchRaw", ctx, req)
	ret0, _ := ret[0].(error)
	return ret0
}

// WriteTaggedBatchRaw indicates an expected call of WriteTaggedBatchRaw
func (mr *MockTChanNodeMockRecorder) WriteTaggedBatchRaw(ctx, req interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "WriteTaggedBatchRaw", reflect.TypeOf((*MockTChanNode)(nil).WriteTaggedBatchRaw), ctx, req)
}
