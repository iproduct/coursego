package maps

import (
	"sync"

	"github.com/cheekybits/genny/generic"
)
//genny get maps/concurrentmap.go "KeyType=BUILTINS ValueType=BUILTINS"

//go:generate genny -in=$GOFILE -out=gen-$GOFILE gen "KeyType=BUILTINS ValueType=BUILTINS"

// KeyType is the key type for the map.
type KeyType generic.Type

// ValueType is the value type for the map.
type ValueType generic.Type

// ConMapKeyTypeValueType is a concurrent safe wrapper around a
// map[KeyType]ValueType.
type ConMapKeyTypeValueType struct {
	// M is the underlying map[KeyType]ValueType.
	M map[KeyType]ValueType
	// L is the sync.RWMutex used to control access to M.
	L sync.RWMutex
}

// NewConMapKeyTypeValueType creates a new concurrent safe map storing ValueType values,
// for KeyType keys.
func NewConMapKeyTypeValueType() *ConMapKeyTypeValueType {
	return ToConMapKeyTypeValueType(nil)
}

// ToConMapKeyTypeValueType creates a new ConMapKeyTypeValueType prepopulated
// with the data from the specified map[KeyType]ValueType.
func ToConMapKeyTypeValueType(data map[KeyType]ValueType) *ConMapKeyTypeValueType {
	if data == nil {
		data = make(map[KeyType]ValueType)
	}
	return &ConMapKeyTypeValueType{M: data}
}

//var nilValueType ValueType

// Get gets the ValueType value for the given KeyType key.
func (cm *ConMapKeyTypeValueType) Get(k KeyType) ValueType {
	cm.L.RLock()
	v := cm.M[k]
	cm.L.RUnlock()
	return v
}

// GetOK gets the ValueType value for the given KeyType key, and a
// bool indicating whether the key was present or not.
func (cm *ConMapKeyTypeValueType) GetOK(k KeyType) (ValueType, bool) {
	cm.L.RLock()
	v, ok := cm.M[k]
	cm.L.RUnlock()
	return v, ok
}

// Set sets the ValueType value for the specified KeyType key.
func (cm *ConMapKeyTypeValueType) Set(k KeyType, v ValueType) {
	cm.L.Lock()
	cm.M[k] = v
	cm.L.Unlock()
}

// Delete removes the specified KeyType key, returning its ValueType
// value and a bool indicating whether the delete occurred or not.
func (cm *ConMapKeyTypeValueType) Delete(k KeyType) (ValueType, bool) {
	cm.L.Lock()
	v, ok := cm.M[k]
	if ok {
		delete(cm.M, k)
	}
	cm.L.Unlock()
	return v, ok
}

// Get gets the ValueType value for the given KeyType key.
func (cm *ConMapKeyTypeValueType) Len() int {
	cm.L.RLock()
	l := len(cm.M)
	cm.L.RUnlock()
	return l
}


