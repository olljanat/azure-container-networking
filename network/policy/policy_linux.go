package policy

import "encoding/json"

type Policy struct {
	Type CNIPolicyType
	Data json.RawMessage
}
