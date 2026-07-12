package datanode_services

import (
	"context"

	"distributed-file-system/internal/cluster/state"
	"distributed-file-system/internal/common"
	"github.com/stretchr/testify/mock"
)

type MockRegisterService struct {
	mock.Mock
}

func (m *MockRegisterService) RegisterWithCoordinator(ctx context.Context, info *common.NodeInfo, csm state.ClusterStateManager, cf state.CoordinatorFinder) error {
	return m.Called(ctx, info, csm, cf).Error(0)
}
