// Code generated by MockGen. DO NOT EDIT.
// Source: service.pb.go

// Package mock_vzmgrpb is a generated GoMock package.
package mock_vzmgrpb

import (
	context "context"
	types "github.com/gogo/protobuf/types"
	gomock "github.com/golang/mock/gomock"
	grpc "google.golang.org/grpc"
	vzmgrpb "pixielabs.ai/pixielabs/src/cloud/vzmgr/vzmgrpb"
	proto "pixielabs.ai/pixielabs/src/common/uuid/proto"
	cvmsgspb "pixielabs.ai/pixielabs/src/shared/cvmsgspb"
	reflect "reflect"
)

// MockVZMgrServiceClient is a mock of VZMgrServiceClient interface
type MockVZMgrServiceClient struct {
	ctrl     *gomock.Controller
	recorder *MockVZMgrServiceClientMockRecorder
}

// MockVZMgrServiceClientMockRecorder is the mock recorder for MockVZMgrServiceClient
type MockVZMgrServiceClientMockRecorder struct {
	mock *MockVZMgrServiceClient
}

// NewMockVZMgrServiceClient creates a new mock instance
func NewMockVZMgrServiceClient(ctrl *gomock.Controller) *MockVZMgrServiceClient {
	mock := &MockVZMgrServiceClient{ctrl: ctrl}
	mock.recorder = &MockVZMgrServiceClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockVZMgrServiceClient) EXPECT() *MockVZMgrServiceClientMockRecorder {
	return m.recorder
}

// CreateVizierCluster mocks base method
func (m *MockVZMgrServiceClient) CreateVizierCluster(ctx context.Context, in *vzmgrpb.CreateVizierClusterRequest, opts ...grpc.CallOption) (*proto.UUID, error) {
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "CreateVizierCluster", varargs...)
	ret0, _ := ret[0].(*proto.UUID)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateVizierCluster indicates an expected call of CreateVizierCluster
func (mr *MockVZMgrServiceClientMockRecorder) CreateVizierCluster(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateVizierCluster", reflect.TypeOf((*MockVZMgrServiceClient)(nil).CreateVizierCluster), varargs...)
}

// GetViziersByOrg mocks base method
func (m *MockVZMgrServiceClient) GetViziersByOrg(ctx context.Context, in *proto.UUID, opts ...grpc.CallOption) (*vzmgrpb.GetViziersByOrgResponse, error) {
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetViziersByOrg", varargs...)
	ret0, _ := ret[0].(*vzmgrpb.GetViziersByOrgResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetViziersByOrg indicates an expected call of GetViziersByOrg
func (mr *MockVZMgrServiceClientMockRecorder) GetViziersByOrg(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetViziersByOrg", reflect.TypeOf((*MockVZMgrServiceClient)(nil).GetViziersByOrg), varargs...)
}

// GetVizierInfo mocks base method
func (m *MockVZMgrServiceClient) GetVizierInfo(ctx context.Context, in *proto.UUID, opts ...grpc.CallOption) (*cvmsgspb.VizierInfo, error) {
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetVizierInfo", varargs...)
	ret0, _ := ret[0].(*cvmsgspb.VizierInfo)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetVizierInfo indicates an expected call of GetVizierInfo
func (mr *MockVZMgrServiceClientMockRecorder) GetVizierInfo(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetVizierInfo", reflect.TypeOf((*MockVZMgrServiceClient)(nil).GetVizierInfo), varargs...)
}

// GetViziersByShard mocks base method
func (m *MockVZMgrServiceClient) GetViziersByShard(ctx context.Context, in *vzmgrpb.GetViziersByShardRequest, opts ...grpc.CallOption) (*vzmgrpb.GetViziersByShardResponse, error) {
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetViziersByShard", varargs...)
	ret0, _ := ret[0].(*vzmgrpb.GetViziersByShardResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetViziersByShard indicates an expected call of GetViziersByShard
func (mr *MockVZMgrServiceClientMockRecorder) GetViziersByShard(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetViziersByShard", reflect.TypeOf((*MockVZMgrServiceClient)(nil).GetViziersByShard), varargs...)
}

// GetVizierConnectionInfo mocks base method
func (m *MockVZMgrServiceClient) GetVizierConnectionInfo(ctx context.Context, in *proto.UUID, opts ...grpc.CallOption) (*cvmsgspb.VizierConnectionInfo, error) {
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetVizierConnectionInfo", varargs...)
	ret0, _ := ret[0].(*cvmsgspb.VizierConnectionInfo)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetVizierConnectionInfo indicates an expected call of GetVizierConnectionInfo
func (mr *MockVZMgrServiceClientMockRecorder) GetVizierConnectionInfo(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetVizierConnectionInfo", reflect.TypeOf((*MockVZMgrServiceClient)(nil).GetVizierConnectionInfo), varargs...)
}

// VizierConnected mocks base method
func (m *MockVZMgrServiceClient) VizierConnected(ctx context.Context, in *cvmsgspb.RegisterVizierRequest, opts ...grpc.CallOption) (*cvmsgspb.RegisterVizierAck, error) {
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "VizierConnected", varargs...)
	ret0, _ := ret[0].(*cvmsgspb.RegisterVizierAck)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// VizierConnected indicates an expected call of VizierConnected
func (mr *MockVZMgrServiceClientMockRecorder) VizierConnected(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "VizierConnected", reflect.TypeOf((*MockVZMgrServiceClient)(nil).VizierConnected), varargs...)
}

// UpdateVizierConfig mocks base method
func (m *MockVZMgrServiceClient) UpdateVizierConfig(ctx context.Context, in *cvmsgspb.UpdateVizierConfigRequest, opts ...grpc.CallOption) (*cvmsgspb.UpdateVizierConfigResponse, error) {
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "UpdateVizierConfig", varargs...)
	ret0, _ := ret[0].(*cvmsgspb.UpdateVizierConfigResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateVizierConfig indicates an expected call of UpdateVizierConfig
func (mr *MockVZMgrServiceClientMockRecorder) UpdateVizierConfig(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateVizierConfig", reflect.TypeOf((*MockVZMgrServiceClient)(nil).UpdateVizierConfig), varargs...)
}

// UpdateOrInstallVizier mocks base method
func (m *MockVZMgrServiceClient) UpdateOrInstallVizier(ctx context.Context, in *cvmsgspb.UpdateOrInstallVizierRequest, opts ...grpc.CallOption) (*cvmsgspb.UpdateOrInstallVizierResponse, error) {
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "UpdateOrInstallVizier", varargs...)
	ret0, _ := ret[0].(*cvmsgspb.UpdateOrInstallVizierResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateOrInstallVizier indicates an expected call of UpdateOrInstallVizier
func (mr *MockVZMgrServiceClientMockRecorder) UpdateOrInstallVizier(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateOrInstallVizier", reflect.TypeOf((*MockVZMgrServiceClient)(nil).UpdateOrInstallVizier), varargs...)
}

// MockVZMgrServiceServer is a mock of VZMgrServiceServer interface
type MockVZMgrServiceServer struct {
	ctrl     *gomock.Controller
	recorder *MockVZMgrServiceServerMockRecorder
}

// MockVZMgrServiceServerMockRecorder is the mock recorder for MockVZMgrServiceServer
type MockVZMgrServiceServerMockRecorder struct {
	mock *MockVZMgrServiceServer
}

// NewMockVZMgrServiceServer creates a new mock instance
func NewMockVZMgrServiceServer(ctrl *gomock.Controller) *MockVZMgrServiceServer {
	mock := &MockVZMgrServiceServer{ctrl: ctrl}
	mock.recorder = &MockVZMgrServiceServerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockVZMgrServiceServer) EXPECT() *MockVZMgrServiceServerMockRecorder {
	return m.recorder
}

// CreateVizierCluster mocks base method
func (m *MockVZMgrServiceServer) CreateVizierCluster(arg0 context.Context, arg1 *vzmgrpb.CreateVizierClusterRequest) (*proto.UUID, error) {
	ret := m.ctrl.Call(m, "CreateVizierCluster", arg0, arg1)
	ret0, _ := ret[0].(*proto.UUID)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateVizierCluster indicates an expected call of CreateVizierCluster
func (mr *MockVZMgrServiceServerMockRecorder) CreateVizierCluster(arg0, arg1 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateVizierCluster", reflect.TypeOf((*MockVZMgrServiceServer)(nil).CreateVizierCluster), arg0, arg1)
}

// GetViziersByOrg mocks base method
func (m *MockVZMgrServiceServer) GetViziersByOrg(arg0 context.Context, arg1 *proto.UUID) (*vzmgrpb.GetViziersByOrgResponse, error) {
	ret := m.ctrl.Call(m, "GetViziersByOrg", arg0, arg1)
	ret0, _ := ret[0].(*vzmgrpb.GetViziersByOrgResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetViziersByOrg indicates an expected call of GetViziersByOrg
func (mr *MockVZMgrServiceServerMockRecorder) GetViziersByOrg(arg0, arg1 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetViziersByOrg", reflect.TypeOf((*MockVZMgrServiceServer)(nil).GetViziersByOrg), arg0, arg1)
}

// GetVizierInfo mocks base method
func (m *MockVZMgrServiceServer) GetVizierInfo(arg0 context.Context, arg1 *proto.UUID) (*cvmsgspb.VizierInfo, error) {
	ret := m.ctrl.Call(m, "GetVizierInfo", arg0, arg1)
	ret0, _ := ret[0].(*cvmsgspb.VizierInfo)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetVizierInfo indicates an expected call of GetVizierInfo
func (mr *MockVZMgrServiceServerMockRecorder) GetVizierInfo(arg0, arg1 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetVizierInfo", reflect.TypeOf((*MockVZMgrServiceServer)(nil).GetVizierInfo), arg0, arg1)
}

// GetViziersByShard mocks base method
func (m *MockVZMgrServiceServer) GetViziersByShard(arg0 context.Context, arg1 *vzmgrpb.GetViziersByShardRequest) (*vzmgrpb.GetViziersByShardResponse, error) {
	ret := m.ctrl.Call(m, "GetViziersByShard", arg0, arg1)
	ret0, _ := ret[0].(*vzmgrpb.GetViziersByShardResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetViziersByShard indicates an expected call of GetViziersByShard
func (mr *MockVZMgrServiceServerMockRecorder) GetViziersByShard(arg0, arg1 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetViziersByShard", reflect.TypeOf((*MockVZMgrServiceServer)(nil).GetViziersByShard), arg0, arg1)
}

// GetVizierConnectionInfo mocks base method
func (m *MockVZMgrServiceServer) GetVizierConnectionInfo(arg0 context.Context, arg1 *proto.UUID) (*cvmsgspb.VizierConnectionInfo, error) {
	ret := m.ctrl.Call(m, "GetVizierConnectionInfo", arg0, arg1)
	ret0, _ := ret[0].(*cvmsgspb.VizierConnectionInfo)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetVizierConnectionInfo indicates an expected call of GetVizierConnectionInfo
func (mr *MockVZMgrServiceServerMockRecorder) GetVizierConnectionInfo(arg0, arg1 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetVizierConnectionInfo", reflect.TypeOf((*MockVZMgrServiceServer)(nil).GetVizierConnectionInfo), arg0, arg1)
}

// VizierConnected mocks base method
func (m *MockVZMgrServiceServer) VizierConnected(arg0 context.Context, arg1 *cvmsgspb.RegisterVizierRequest) (*cvmsgspb.RegisterVizierAck, error) {
	ret := m.ctrl.Call(m, "VizierConnected", arg0, arg1)
	ret0, _ := ret[0].(*cvmsgspb.RegisterVizierAck)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// VizierConnected indicates an expected call of VizierConnected
func (mr *MockVZMgrServiceServerMockRecorder) VizierConnected(arg0, arg1 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "VizierConnected", reflect.TypeOf((*MockVZMgrServiceServer)(nil).VizierConnected), arg0, arg1)
}

// UpdateVizierConfig mocks base method
func (m *MockVZMgrServiceServer) UpdateVizierConfig(arg0 context.Context, arg1 *cvmsgspb.UpdateVizierConfigRequest) (*cvmsgspb.UpdateVizierConfigResponse, error) {
	ret := m.ctrl.Call(m, "UpdateVizierConfig", arg0, arg1)
	ret0, _ := ret[0].(*cvmsgspb.UpdateVizierConfigResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateVizierConfig indicates an expected call of UpdateVizierConfig
func (mr *MockVZMgrServiceServerMockRecorder) UpdateVizierConfig(arg0, arg1 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateVizierConfig", reflect.TypeOf((*MockVZMgrServiceServer)(nil).UpdateVizierConfig), arg0, arg1)
}

// UpdateOrInstallVizier mocks base method
func (m *MockVZMgrServiceServer) UpdateOrInstallVizier(arg0 context.Context, arg1 *cvmsgspb.UpdateOrInstallVizierRequest) (*cvmsgspb.UpdateOrInstallVizierResponse, error) {
	ret := m.ctrl.Call(m, "UpdateOrInstallVizier", arg0, arg1)
	ret0, _ := ret[0].(*cvmsgspb.UpdateOrInstallVizierResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateOrInstallVizier indicates an expected call of UpdateOrInstallVizier
func (mr *MockVZMgrServiceServerMockRecorder) UpdateOrInstallVizier(arg0, arg1 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateOrInstallVizier", reflect.TypeOf((*MockVZMgrServiceServer)(nil).UpdateOrInstallVizier), arg0, arg1)
}

// MockVZDeploymentKeyServiceClient is a mock of VZDeploymentKeyServiceClient interface
type MockVZDeploymentKeyServiceClient struct {
	ctrl     *gomock.Controller
	recorder *MockVZDeploymentKeyServiceClientMockRecorder
}

// MockVZDeploymentKeyServiceClientMockRecorder is the mock recorder for MockVZDeploymentKeyServiceClient
type MockVZDeploymentKeyServiceClientMockRecorder struct {
	mock *MockVZDeploymentKeyServiceClient
}

// NewMockVZDeploymentKeyServiceClient creates a new mock instance
func NewMockVZDeploymentKeyServiceClient(ctrl *gomock.Controller) *MockVZDeploymentKeyServiceClient {
	mock := &MockVZDeploymentKeyServiceClient{ctrl: ctrl}
	mock.recorder = &MockVZDeploymentKeyServiceClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockVZDeploymentKeyServiceClient) EXPECT() *MockVZDeploymentKeyServiceClientMockRecorder {
	return m.recorder
}

// Create mocks base method
func (m *MockVZDeploymentKeyServiceClient) Create(ctx context.Context, in *vzmgrpb.CreateDeploymentKeyRequest, opts ...grpc.CallOption) (*vzmgrpb.DeploymentKey, error) {
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Create", varargs...)
	ret0, _ := ret[0].(*vzmgrpb.DeploymentKey)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Create indicates an expected call of Create
func (mr *MockVZDeploymentKeyServiceClientMockRecorder) Create(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockVZDeploymentKeyServiceClient)(nil).Create), varargs...)
}

// List mocks base method
func (m *MockVZDeploymentKeyServiceClient) List(ctx context.Context, in *vzmgrpb.ListDeploymentKeyRequest, opts ...grpc.CallOption) (*vzmgrpb.ListDeploymentKeyResponse, error) {
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "List", varargs...)
	ret0, _ := ret[0].(*vzmgrpb.ListDeploymentKeyResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// List indicates an expected call of List
func (mr *MockVZDeploymentKeyServiceClientMockRecorder) List(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "List", reflect.TypeOf((*MockVZDeploymentKeyServiceClient)(nil).List), varargs...)
}

// Get mocks base method
func (m *MockVZDeploymentKeyServiceClient) Get(ctx context.Context, in *vzmgrpb.GetDeploymentKeyRequest, opts ...grpc.CallOption) (*vzmgrpb.GetDeploymentKeyResponse, error) {
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Get", varargs...)
	ret0, _ := ret[0].(*vzmgrpb.GetDeploymentKeyResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Get indicates an expected call of Get
func (mr *MockVZDeploymentKeyServiceClientMockRecorder) Get(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockVZDeploymentKeyServiceClient)(nil).Get), varargs...)
}

// Delete mocks base method
func (m *MockVZDeploymentKeyServiceClient) Delete(ctx context.Context, in *proto.UUID, opts ...grpc.CallOption) (*types.Empty, error) {
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Delete", varargs...)
	ret0, _ := ret[0].(*types.Empty)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Delete indicates an expected call of Delete
func (mr *MockVZDeploymentKeyServiceClientMockRecorder) Delete(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockVZDeploymentKeyServiceClient)(nil).Delete), varargs...)
}

// MockVZDeploymentKeyServiceServer is a mock of VZDeploymentKeyServiceServer interface
type MockVZDeploymentKeyServiceServer struct {
	ctrl     *gomock.Controller
	recorder *MockVZDeploymentKeyServiceServerMockRecorder
}

// MockVZDeploymentKeyServiceServerMockRecorder is the mock recorder for MockVZDeploymentKeyServiceServer
type MockVZDeploymentKeyServiceServerMockRecorder struct {
	mock *MockVZDeploymentKeyServiceServer
}

// NewMockVZDeploymentKeyServiceServer creates a new mock instance
func NewMockVZDeploymentKeyServiceServer(ctrl *gomock.Controller) *MockVZDeploymentKeyServiceServer {
	mock := &MockVZDeploymentKeyServiceServer{ctrl: ctrl}
	mock.recorder = &MockVZDeploymentKeyServiceServerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockVZDeploymentKeyServiceServer) EXPECT() *MockVZDeploymentKeyServiceServerMockRecorder {
	return m.recorder
}

// Create mocks base method
func (m *MockVZDeploymentKeyServiceServer) Create(arg0 context.Context, arg1 *vzmgrpb.CreateDeploymentKeyRequest) (*vzmgrpb.DeploymentKey, error) {
	ret := m.ctrl.Call(m, "Create", arg0, arg1)
	ret0, _ := ret[0].(*vzmgrpb.DeploymentKey)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Create indicates an expected call of Create
func (mr *MockVZDeploymentKeyServiceServerMockRecorder) Create(arg0, arg1 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockVZDeploymentKeyServiceServer)(nil).Create), arg0, arg1)
}

// List mocks base method
func (m *MockVZDeploymentKeyServiceServer) List(arg0 context.Context, arg1 *vzmgrpb.ListDeploymentKeyRequest) (*vzmgrpb.ListDeploymentKeyResponse, error) {
	ret := m.ctrl.Call(m, "List", arg0, arg1)
	ret0, _ := ret[0].(*vzmgrpb.ListDeploymentKeyResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// List indicates an expected call of List
func (mr *MockVZDeploymentKeyServiceServerMockRecorder) List(arg0, arg1 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "List", reflect.TypeOf((*MockVZDeploymentKeyServiceServer)(nil).List), arg0, arg1)
}

// Get mocks base method
func (m *MockVZDeploymentKeyServiceServer) Get(arg0 context.Context, arg1 *vzmgrpb.GetDeploymentKeyRequest) (*vzmgrpb.GetDeploymentKeyResponse, error) {
	ret := m.ctrl.Call(m, "Get", arg0, arg1)
	ret0, _ := ret[0].(*vzmgrpb.GetDeploymentKeyResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Get indicates an expected call of Get
func (mr *MockVZDeploymentKeyServiceServerMockRecorder) Get(arg0, arg1 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockVZDeploymentKeyServiceServer)(nil).Get), arg0, arg1)
}

// Delete mocks base method
func (m *MockVZDeploymentKeyServiceServer) Delete(arg0 context.Context, arg1 *proto.UUID) (*types.Empty, error) {
	ret := m.ctrl.Call(m, "Delete", arg0, arg1)
	ret0, _ := ret[0].(*types.Empty)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Delete indicates an expected call of Delete
func (mr *MockVZDeploymentKeyServiceServerMockRecorder) Delete(arg0, arg1 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockVZDeploymentKeyServiceServer)(nil).Delete), arg0, arg1)
}
