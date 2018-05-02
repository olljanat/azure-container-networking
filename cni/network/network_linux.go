package network

import (
	"github.com/Azure/azure-container-networking/cni"
	cniTypesCurr "github.com/containernetworking/cni/pkg/types/current"
)

// HandleConsecutiveAdd is a dummy function for Linux platform.
func HandleConsecutiveAdd(containerId, endpointId string, nwInfo *NetworkInfo, nwCfg *cni.NetworkConfig) (*cniTypesCurr.Result, error) {
	return nil, nil
}
