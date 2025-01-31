package indelible

import (
	"go.k6.io/k6/js/modules"
)
type IndelibleArray struct{}
func init() {
	modules.Register("k6/x/xk6-indelible-array", new(IndelibleArray))
}

var indelibleArray []interface{} // global array

func (m *IndelibleArray) InitIndelibleArray() {
	indelibleArray = indelibleArray[:0]
}

func (m *IndelibleArray) PushToIndelibleArray(value interface{}) {
	indelibleArray = append(indelibleArray, value) // Add value any format
}

func (m *IndelibleArray) GetIndelibleArray() []interface{}{
	return indelibleArray
}
