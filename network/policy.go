package network

import (
	"encoding/json"
	"strings"

	"github.com/Azure/azure-container-networking/cni"
)

type CNIPolicyType string

const (
	PolicyStr         string        = "Policy"
	NetworkPolicy     CNIPolicyType = "NetworkPolicy"
	EndpointPolicy    CNIPolicyType = "EndpointPolicy"
	OutBoundNatPolicy CNIPolicyType = "OutBoundNatPolicy"
)

type Policy struct {
	Type CNIPolicyType
	Data json.RawMessage
}

// SerializePolicies serializes policies to json.
func SerializePolicies(policyType CNIPolicyType, policies []Policy) []json.RawMessage {
	var jsonPolicies []json.RawMessage
	for _, policy := range policies {
		if policy.Type == policyType {
			jsonPolicies = append(jsonPolicies, policy.Data)
		}
	}
	return jsonPolicies
}

// GetPoliciesFromNwCfg returns network policies from network config.
func GetPoliciesFromNwCfg(kvp []cni.KVPair) []Policy {
	var policies []Policy
	for _, pair := range kvp {
		if strings.Contains(pair.Name, PolicyStr) {
			policy := Policy{
				Type: CNIPolicyType(pair.Name),
				Data: pair.Value,
			}
			policies = append(policies, policy)
		}
	}

	return policies
}
