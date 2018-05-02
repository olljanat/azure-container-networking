// Copyright 2017 Microsoft. All rights reserved.
// MIT License

package cni

import (
	"encoding/json"

	cniTypes "github.com/containernetworking/cni/pkg/types"
)

// KVPair represents a K-V pair of a json object.
type KVPair struct {
	Name  string          `json:"name"`
	Value json.RawMessage `json:"value"`
}

// NetworkConfig represents Azure CNI plugin network configuration.
type NetworkConfig struct {
	CNIVersion string `json:"cniVersion"`
	Name       string `json:"name"`
	Type       string `json:"type"`
	Mode       string `json:"mode"`
	Master     string `json:"master"`
	Bridge     string `json:"bridge,omitempty"`
	LogLevel   string `json:"logLevel,omitempty"`
	LogTarget  string `json:"logTarget,omitempty"`
	Ipam       struct {
		Type          string `json:"type"`
		Environment   string `json:"environment,omitempty"`
		AddrSpace     string `json:"addressSpace,omitempty"`
		Subnet        string `json:"subnet,omitempty"`
		Address       string `json:"ipAddress,omitempty"`
		QueryInterval string `json:"queryInterval,omitempty"`
	}
	DNS            cniTypes.DNS `json:"dns"`
	AdditionalArgs []KVPair
}

type K8SPodEnvArgs struct {
	cniTypes.CommonArgs
	K8S_POD_NAMESPACE          cniTypes.UnmarshallableString `json:"K8S_POD_NAMESPACE,omitempty"`
	K8S_POD_NAME               cniTypes.UnmarshallableString `json:"K8S_POD_NAME,omitempty"`
	K8S_POD_INFRA_CONTAINER_ID cniTypes.UnmarshallableString `json:"K8S_POD_INFRA_CONTAINER_ID,omitempty"`
}

// ParseCniArgs unmarshals cni arguments.
func ParseCniArgs(args string) (*K8SPodEnvArgs, error) {
	podCfg := K8SPodEnvArgs{}
	err := cniTypes.LoadArgs(args, &podCfg)
	if err != nil {
		return nil, err
	}

	return &podCfg, nil
}

// ParseNetworkConfig unmarshals network configuration from bytes.
func ParseNetworkConfig(b []byte) (*NetworkConfig, error) {
	nwCfg := NetworkConfig{}

	err := json.Unmarshal(b, &nwCfg)
	if err != nil {
		return nil, err
	}

	if nwCfg.CNIVersion == "" {
		nwCfg.CNIVersion = defaultVersion
	}

	return &nwCfg, nil
}

// Serialize marshals a network configuration to bytes.
func (nwcfg *NetworkConfig) Serialize() []byte {
	bytes, _ := json.Marshal(nwcfg)
	return bytes
}
