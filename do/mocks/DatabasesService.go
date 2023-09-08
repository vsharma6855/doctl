// Code generated by MockGen. DO NOT EDIT.
// Source: databases.go

// Package mocks is a generated GoMock package.
package mocks

import (
	reflect "reflect"

	do "github.com/digitalocean/doctl/do"
	godo "github.com/digitalocean/godo"
	gomock "go.uber.org/mock/gomock"
)

// MockDatabasesService is a mock of DatabasesService interface.
type MockDatabasesService struct {
	ctrl     *gomock.Controller
	recorder *MockDatabasesServiceMockRecorder
}

// MockDatabasesServiceMockRecorder is the mock recorder for MockDatabasesService.
type MockDatabasesServiceMockRecorder struct {
	mock *MockDatabasesService
}

// NewMockDatabasesService creates a new mock instance.
func NewMockDatabasesService(ctrl *gomock.Controller) *MockDatabasesService {
	mock := &MockDatabasesService{ctrl: ctrl}
	mock.recorder = &MockDatabasesServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockDatabasesService) EXPECT() *MockDatabasesServiceMockRecorder {
	return m.recorder
}

// Create mocks base method.
func (m *MockDatabasesService) Create(arg0 *godo.DatabaseCreateRequest) (*do.Database, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", arg0)
	ret0, _ := ret[0].(*do.Database)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Create indicates an expected call of Create.
func (mr *MockDatabasesServiceMockRecorder) Create(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockDatabasesService)(nil).Create), arg0)
}

// CreateDB mocks base method.
func (m *MockDatabasesService) CreateDB(arg0 string, arg1 *godo.DatabaseCreateDBRequest) (*do.DatabaseDB, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateDB", arg0, arg1)
	ret0, _ := ret[0].(*do.DatabaseDB)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateDB indicates an expected call of CreateDB.
func (mr *MockDatabasesServiceMockRecorder) CreateDB(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateDB", reflect.TypeOf((*MockDatabasesService)(nil).CreateDB), arg0, arg1)
}

// CreatePool mocks base method.
func (m *MockDatabasesService) CreatePool(arg0 string, arg1 *godo.DatabaseCreatePoolRequest) (*do.DatabasePool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreatePool", arg0, arg1)
	ret0, _ := ret[0].(*do.DatabasePool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreatePool indicates an expected call of CreatePool.
func (mr *MockDatabasesServiceMockRecorder) CreatePool(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreatePool", reflect.TypeOf((*MockDatabasesService)(nil).CreatePool), arg0, arg1)
}

// CreateReplica mocks base method.
func (m *MockDatabasesService) CreateReplica(arg0 string, arg1 *godo.DatabaseCreateReplicaRequest) (*do.DatabaseReplica, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateReplica", arg0, arg1)
	ret0, _ := ret[0].(*do.DatabaseReplica)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateReplica indicates an expected call of CreateReplica.
func (mr *MockDatabasesServiceMockRecorder) CreateReplica(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateReplica", reflect.TypeOf((*MockDatabasesService)(nil).CreateReplica), arg0, arg1)
}

// CreateUser mocks base method.
func (m *MockDatabasesService) CreateUser(arg0 string, arg1 *godo.DatabaseCreateUserRequest) (*do.DatabaseUser, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateUser", arg0, arg1)
	ret0, _ := ret[0].(*do.DatabaseUser)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateUser indicates an expected call of CreateUser.
func (mr *MockDatabasesServiceMockRecorder) CreateUser(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateUser", reflect.TypeOf((*MockDatabasesService)(nil).CreateUser), arg0, arg1)
}

// Delete mocks base method.
func (m *MockDatabasesService) Delete(arg0 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// Delete indicates an expected call of Delete.
func (mr *MockDatabasesServiceMockRecorder) Delete(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockDatabasesService)(nil).Delete), arg0)
}

// DeleteDB mocks base method.
func (m *MockDatabasesService) DeleteDB(arg0, arg1 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteDB", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteDB indicates an expected call of DeleteDB.
func (mr *MockDatabasesServiceMockRecorder) DeleteDB(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteDB", reflect.TypeOf((*MockDatabasesService)(nil).DeleteDB), arg0, arg1)
}

// DeletePool mocks base method.
func (m *MockDatabasesService) DeletePool(arg0, arg1 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeletePool", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeletePool indicates an expected call of DeletePool.
func (mr *MockDatabasesServiceMockRecorder) DeletePool(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeletePool", reflect.TypeOf((*MockDatabasesService)(nil).DeletePool), arg0, arg1)
}

// DeleteReplica mocks base method.
func (m *MockDatabasesService) DeleteReplica(arg0, arg1 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteReplica", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteReplica indicates an expected call of DeleteReplica.
func (mr *MockDatabasesServiceMockRecorder) DeleteReplica(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteReplica", reflect.TypeOf((*MockDatabasesService)(nil).DeleteReplica), arg0, arg1)
}

// DeleteUser mocks base method.
func (m *MockDatabasesService) DeleteUser(arg0, arg1 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteUser", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteUser indicates an expected call of DeleteUser.
func (mr *MockDatabasesServiceMockRecorder) DeleteUser(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteUser", reflect.TypeOf((*MockDatabasesService)(nil).DeleteUser), arg0, arg1)
}

// Get mocks base method.
func (m *MockDatabasesService) Get(arg0 string) (*do.Database, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get", arg0)
	ret0, _ := ret[0].(*do.Database)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Get indicates an expected call of Get.
func (mr *MockDatabasesServiceMockRecorder) Get(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockDatabasesService)(nil).Get), arg0)
}

// GetConnection mocks base method.
func (m *MockDatabasesService) GetConnection(arg0 string, arg1 bool) (*do.DatabaseConnection, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetConnection", arg0, arg1)
	ret0, _ := ret[0].(*do.DatabaseConnection)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetConnection indicates an expected call of GetConnection.
func (mr *MockDatabasesServiceMockRecorder) GetConnection(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetConnection", reflect.TypeOf((*MockDatabasesService)(nil).GetConnection), arg0, arg1)
}

// GetDB mocks base method.
func (m *MockDatabasesService) GetDB(arg0, arg1 string) (*do.DatabaseDB, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetDB", arg0, arg1)
	ret0, _ := ret[0].(*do.DatabaseDB)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetDB indicates an expected call of GetDB.
func (mr *MockDatabasesServiceMockRecorder) GetDB(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetDB", reflect.TypeOf((*MockDatabasesService)(nil).GetDB), arg0, arg1)
}

// GetFirewallRules mocks base method.
func (m *MockDatabasesService) GetFirewallRules(arg0 string) (do.DatabaseFirewallRules, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetFirewallRules", arg0)
	ret0, _ := ret[0].(do.DatabaseFirewallRules)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetFirewallRules indicates an expected call of GetFirewallRules.
func (mr *MockDatabasesServiceMockRecorder) GetFirewallRules(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetFirewallRules", reflect.TypeOf((*MockDatabasesService)(nil).GetFirewallRules), arg0)
}

// GetMaintenance mocks base method.
func (m *MockDatabasesService) GetMaintenance(arg0 string) (*do.DatabaseMaintenanceWindow, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetMaintenance", arg0)
	ret0, _ := ret[0].(*do.DatabaseMaintenanceWindow)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetMaintenance indicates an expected call of GetMaintenance.
func (mr *MockDatabasesServiceMockRecorder) GetMaintenance(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetMaintenance", reflect.TypeOf((*MockDatabasesService)(nil).GetMaintenance), arg0)
}

// GetPool mocks base method.
func (m *MockDatabasesService) GetPool(arg0, arg1 string) (*do.DatabasePool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetPool", arg0, arg1)
	ret0, _ := ret[0].(*do.DatabasePool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetPool indicates an expected call of GetPool.
func (mr *MockDatabasesServiceMockRecorder) GetPool(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetPool", reflect.TypeOf((*MockDatabasesService)(nil).GetPool), arg0, arg1)
}

// GetReplica mocks base method.
func (m *MockDatabasesService) GetReplica(arg0, arg1 string) (*do.DatabaseReplica, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetReplica", arg0, arg1)
	ret0, _ := ret[0].(*do.DatabaseReplica)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetReplica indicates an expected call of GetReplica.
func (mr *MockDatabasesServiceMockRecorder) GetReplica(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetReplica", reflect.TypeOf((*MockDatabasesService)(nil).GetReplica), arg0, arg1)
}

// GetReplicaConnection mocks base method.
func (m *MockDatabasesService) GetReplicaConnection(arg0, arg1 string) (*do.DatabaseConnection, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetReplicaConnection", arg0, arg1)
	ret0, _ := ret[0].(*do.DatabaseConnection)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetReplicaConnection indicates an expected call of GetReplicaConnection.
func (mr *MockDatabasesServiceMockRecorder) GetReplicaConnection(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetReplicaConnection", reflect.TypeOf((*MockDatabasesService)(nil).GetReplicaConnection), arg0, arg1)
}

// GetSQLMode mocks base method.
func (m *MockDatabasesService) GetSQLMode(arg0 string) ([]string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetSQLMode", arg0)
	ret0, _ := ret[0].([]string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetSQLMode indicates an expected call of GetSQLMode.
func (mr *MockDatabasesServiceMockRecorder) GetSQLMode(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetSQLMode", reflect.TypeOf((*MockDatabasesService)(nil).GetSQLMode), arg0)
}

// GetUser mocks base method.
func (m *MockDatabasesService) GetUser(arg0, arg1 string) (*do.DatabaseUser, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUser", arg0, arg1)
	ret0, _ := ret[0].(*do.DatabaseUser)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUser indicates an expected call of GetUser.
func (mr *MockDatabasesServiceMockRecorder) GetUser(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUser", reflect.TypeOf((*MockDatabasesService)(nil).GetUser), arg0, arg1)
}

// List mocks base method.
func (m *MockDatabasesService) List() (do.Databases, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "List")
	ret0, _ := ret[0].(do.Databases)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// List indicates an expected call of List.
func (mr *MockDatabasesServiceMockRecorder) List() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "List", reflect.TypeOf((*MockDatabasesService)(nil).List))
}

// ListBackups mocks base method.
func (m *MockDatabasesService) ListBackups(arg0 string) (do.DatabaseBackups, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListBackups", arg0)
	ret0, _ := ret[0].(do.DatabaseBackups)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListBackups indicates an expected call of ListBackups.
func (mr *MockDatabasesServiceMockRecorder) ListBackups(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListBackups", reflect.TypeOf((*MockDatabasesService)(nil).ListBackups), arg0)
}

// ListDBs mocks base method.
func (m *MockDatabasesService) ListDBs(arg0 string) (do.DatabaseDBs, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListDBs", arg0)
	ret0, _ := ret[0].(do.DatabaseDBs)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListDBs indicates an expected call of ListDBs.
func (mr *MockDatabasesServiceMockRecorder) ListDBs(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListDBs", reflect.TypeOf((*MockDatabasesService)(nil).ListDBs), arg0)
}

// ListOptions mocks base method.
func (m *MockDatabasesService) ListOptions() (*do.DatabaseOptions, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListOptions")
	ret0, _ := ret[0].(*do.DatabaseOptions)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListOptions indicates an expected call of ListOptions.
func (mr *MockDatabasesServiceMockRecorder) ListOptions() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListOptions", reflect.TypeOf((*MockDatabasesService)(nil).ListOptions))
}

// ListPools mocks base method.
func (m *MockDatabasesService) ListPools(arg0 string) (do.DatabasePools, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListPools", arg0)
	ret0, _ := ret[0].(do.DatabasePools)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListPools indicates an expected call of ListPools.
func (mr *MockDatabasesServiceMockRecorder) ListPools(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListPools", reflect.TypeOf((*MockDatabasesService)(nil).ListPools), arg0)
}

// ListReplicas mocks base method.
func (m *MockDatabasesService) ListReplicas(arg0 string) (do.DatabaseReplicas, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListReplicas", arg0)
	ret0, _ := ret[0].(do.DatabaseReplicas)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListReplicas indicates an expected call of ListReplicas.
func (mr *MockDatabasesServiceMockRecorder) ListReplicas(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListReplicas", reflect.TypeOf((*MockDatabasesService)(nil).ListReplicas), arg0)
}

// ListUsers mocks base method.
func (m *MockDatabasesService) ListUsers(arg0 string) (do.DatabaseUsers, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListUsers", arg0)
	ret0, _ := ret[0].(do.DatabaseUsers)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListUsers indicates an expected call of ListUsers.
func (mr *MockDatabasesServiceMockRecorder) ListUsers(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListUsers", reflect.TypeOf((*MockDatabasesService)(nil).ListUsers), arg0)
}

// Migrate mocks base method.
func (m *MockDatabasesService) Migrate(arg0 string, arg1 *godo.DatabaseMigrateRequest) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Migrate", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// Migrate indicates an expected call of Migrate.
func (mr *MockDatabasesServiceMockRecorder) Migrate(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Migrate", reflect.TypeOf((*MockDatabasesService)(nil).Migrate), arg0, arg1)
}

// PromoteReplica mocks base method.
func (m *MockDatabasesService) PromoteReplica(arg0, arg1 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PromoteReplica", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// PromoteReplica indicates an expected call of PromoteReplica.
func (mr *MockDatabasesServiceMockRecorder) PromoteReplica(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PromoteReplica", reflect.TypeOf((*MockDatabasesService)(nil).PromoteReplica), arg0, arg1)
}

// ResetUserAuth mocks base method.
func (m *MockDatabasesService) ResetUserAuth(arg0, arg1 string, arg2 *godo.DatabaseResetUserAuthRequest) (*do.DatabaseUser, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ResetUserAuth", arg0, arg1, arg2)
	ret0, _ := ret[0].(*do.DatabaseUser)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ResetUserAuth indicates an expected call of ResetUserAuth.
func (mr *MockDatabasesServiceMockRecorder) ResetUserAuth(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ResetUserAuth", reflect.TypeOf((*MockDatabasesService)(nil).ResetUserAuth), arg0, arg1, arg2)
}

// Resize mocks base method.
func (m *MockDatabasesService) Resize(arg0 string, arg1 *godo.DatabaseResizeRequest) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Resize", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// Resize indicates an expected call of Resize.
func (mr *MockDatabasesServiceMockRecorder) Resize(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Resize", reflect.TypeOf((*MockDatabasesService)(nil).Resize), arg0, arg1)
}

// SetSQLMode mocks base method.
func (m *MockDatabasesService) SetSQLMode(arg0 string, arg1 ...string) error {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0}
	for _, a := range arg1 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "SetSQLMode", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// SetSQLMode indicates an expected call of SetSQLMode.
func (mr *MockDatabasesServiceMockRecorder) SetSQLMode(arg0 interface{}, arg1 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0}, arg1...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetSQLMode", reflect.TypeOf((*MockDatabasesService)(nil).SetSQLMode), varargs...)
}

// UpdateFirewallRules mocks base method.
func (m *MockDatabasesService) UpdateFirewallRules(databaseID string, req *godo.DatabaseUpdateFirewallRulesRequest) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateFirewallRules", databaseID, req)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateFirewallRules indicates an expected call of UpdateFirewallRules.
func (mr *MockDatabasesServiceMockRecorder) UpdateFirewallRules(databaseID, req interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateFirewallRules", reflect.TypeOf((*MockDatabasesService)(nil).UpdateFirewallRules), databaseID, req)
}

// UpdateMaintenance mocks base method.
func (m *MockDatabasesService) UpdateMaintenance(arg0 string, arg1 *godo.DatabaseUpdateMaintenanceRequest) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateMaintenance", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateMaintenance indicates an expected call of UpdateMaintenance.
func (mr *MockDatabasesServiceMockRecorder) UpdateMaintenance(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateMaintenance", reflect.TypeOf((*MockDatabasesService)(nil).UpdateMaintenance), arg0, arg1)
}
