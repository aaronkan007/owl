// Code generated by MockGen. DO NOT EDIT.
// Source: ./service/db_info/cluster.go

// Package mock is a generated GoMock package.
package mock

import (
	gomock "github.com/golang/mock/gomock"
	db_info "gitlab.pri.ibanyu.com/middleware/dbinjection/service/db_info"
	reflect "reflect"
)

// MockClusterDao is a mock of ClusterDao interface
type MockClusterDao struct {
	ctrl     *gomock.Controller
	recorder *MockClusterDaoMockRecorder
}

// MockClusterDaoMockRecorder is the mock recorder for MockClusterDao
type MockClusterDaoMockRecorder struct {
	mock *MockClusterDao
}

// NewMockClusterDao creates a new mock instance
func NewMockClusterDao(ctrl *gomock.Controller) *MockClusterDao {
	mock := &MockClusterDao{ctrl: ctrl}
	mock.recorder = &MockClusterDaoMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockClusterDao) EXPECT() *MockClusterDaoMockRecorder {
	return m.recorder
}

// AddCluster mocks base method
func (m *MockClusterDao) AddCluster(cluster *db_info.DbInjectionCluster) (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddCluster", cluster)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AddCluster indicates an expected call of AddCluster
func (mr *MockClusterDaoMockRecorder) AddCluster(cluster interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddCluster", reflect.TypeOf((*MockClusterDao)(nil).AddCluster), cluster)
}

// UpdateCluster mocks base method
func (m *MockClusterDao) UpdateCluster(cluster *db_info.DbInjectionCluster) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateCluster", cluster)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateCluster indicates an expected call of UpdateCluster
func (mr *MockClusterDaoMockRecorder) UpdateCluster(cluster interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateCluster", reflect.TypeOf((*MockClusterDao)(nil).UpdateCluster), cluster)
}

// DelCluster mocks base method
func (m *MockClusterDao) DelCluster(id int64) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DelCluster", id)
	ret0, _ := ret[0].(error)
	return ret0
}

// DelCluster indicates an expected call of DelCluster
func (mr *MockClusterDaoMockRecorder) DelCluster(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DelCluster", reflect.TypeOf((*MockClusterDao)(nil).DelCluster), id)
}

// GetClusterByName mocks base method
func (m *MockClusterDao) GetClusterByName(clusterName string) (*db_info.DbInjectionCluster, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetClusterByName", clusterName)
	ret0, _ := ret[0].(*db_info.DbInjectionCluster)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetClusterByName indicates an expected call of GetClusterByName
func (mr *MockClusterDaoMockRecorder) GetClusterByName(clusterName interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetClusterByName", reflect.TypeOf((*MockClusterDao)(nil).GetClusterByName), clusterName)
}

// ListCluster mocks base method
func (m *MockClusterDao) ListCluster() ([]db_info.DbInjectionCluster, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListCluster")
	ret0, _ := ret[0].([]db_info.DbInjectionCluster)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListCluster indicates an expected call of ListCluster
func (mr *MockClusterDaoMockRecorder) ListCluster() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListCluster", reflect.TypeOf((*MockClusterDao)(nil).ListCluster))
}
