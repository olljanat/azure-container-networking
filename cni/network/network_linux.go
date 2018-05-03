package network

import (
	"github.com/Azure/azure-container-networking/network"
	"github.com/Azure/azure-container-networking/network/policy"
	cniTypesCurr "github.com/containernetworking/cni/pkg/types/current"
)

// handleConsecutiveAdd is a dummy function for Linux platform.
func handleConsecutiveAdd(containerId, endpointId string, nwInfo *NetworkInfo, nwCfg *NetworkConfig) (*cniTypesCurr.Result, error) {
	return nil, nil
}

func setPolicies(epInfo *network.EndpointInfo, policies []policy.Policy) {}
