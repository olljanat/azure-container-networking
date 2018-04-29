// Copyright Microsoft Corp.
// All rights reserved.

package network

import (
	"encoding/json"
)

type CNIPolicyType string

const (
	OutBoundNatPolicy CNIPolicyType = "OutBoundNatPolicy"
)

type Policy struct {
	Type CNIPolicyType
	Data json.RawMessage
}
