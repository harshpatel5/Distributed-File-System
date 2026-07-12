package datanode_controllers

import (
	"distributed-file-system/internal/cluster/state"
	"distributed-file-system/internal/common"
	"github.com/stretchr/testify/mock"
)

type MockHeartbeatController struct {
	mock.Mock
}

func (m *MockHeartbeatController) Run(info *common.NodeInfo, csm state.ClusterStateManager, cf state.CoordinatorFinder) error {
	return m.Called(info, csm, cf).Error(0)
}
