package maps

import (
	"github.com/cheekybits/genny/generic"
	"sync"
)

//go:generate genny -in=$GOFILE -out=gen_$GOFILE gen "KeyType=BUILTINS ValueType=BUILTINS"

type KeyType generic.Type
type ValueType generic.Type

type ConMapKeyTypeValueType struct {
	M map[KeyType]ValueType
	L sync.RWMutex
}

func NewConMapKeyTypeValueType() *ConMapKeyTypeValueType {
	return ToConMapKeyTypeValueType(nil)
}

func ToConMapKeyTypeValueType(data map[KeyType]ValueType) *ConMapKeyTypeValueType {
	if data == nil {
		data = make(map[KeyType]ValueType)
	}
	return &ConMapKeyTypeValueType{M: data}
}

func (cm *ConMapKeyTypeValueType) Get(k KeyType) ValueType {
	cm.L.RLock()
	v := cm.M[k]
	cm.L.RUnlock()
	return v
}
func (cm *ConMapKeyTypeValueType) GetOK(k KeyType) (ValueType, bool) {
	cm.L.RLock()
	v, ok := cm.M[k]
	cm.L.RUnlock()
	return v, ok
}

func (cm *ConMapKeyTypeValueType) Set(k KeyType, v ValueType) {
	cm.L.Lock()
	cm.M[k] = v
	cm.L.Unlock()
}

func (cm *ConMapKeyTypeValueType) Delete(k KeyType) (ValueType, bool){
	cm.L.Lock()
	v, ok := cm.M[k]
	if ok {
		delete (cm.M, k)
	}
	cm.L.Unlock()
	return v, ok
}

func (cm *ConMapKeyTypeValueType) Len() int {
	cm.L.RLock()
	l := len(cm.M)
	cm.L.RUnlock()
	return l
}




