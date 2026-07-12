package datanode_services

import "distributed-file-system/internal/cluster/state"

type NodeAgentServices struct {
	Register    RegisterProvider
	Coordinator state.CoordinatorFinder
}

func NewNodeAgentServices(coordinatorFinder state.CoordinatorFinder, registerService RegisterProvider) NodeAgentServices {
	return NodeAgentServices{
		Register:    registerService,
		Coordinator: coordinatorFinder,
	}
}
