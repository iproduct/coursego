package maps_test

import (
	"testing"

	"github.com/iproduct/coursego/generationlab/genny/maps"
	"github.com/cheekybits/is"
)

func TestNew(t *testing.T) {
	is := is.New(t)

	k1 := new(maps.KeyType)
	k2 := new(maps.KeyType)
	v1 := new(maps.ValueType)
	v2 := new(maps.ValueType)

	m := maps.NewConMapKeyTypeValueType()
	_, ok := m.GetOK(k1)
	is.Equal(ok, false)

	m.Set(k1, v1)
	is.Equal(m.Get(k1), v1)
	m.Set(k2, v2)
	is.Equal(m.Get(k2), v2)

	v, ok := m.GetOK(k1)
	is.Equal(ok, true)
	is.Equal(v, v1)

	// delete k2
	v, ok = m.Delete(k2)
	is.Equal(ok, true)
	is.Equal(v, v2)

	_, ok = m.GetOK(k2)
	is.Equal(ok, false)

}

func TestTo(t *testing.T) {
	is := is.New(t)

	k1 := new(maps.KeyType)
	k2 := new(maps.KeyType)
	v1 := new(maps.ValueType)
	v2 := new(maps.ValueType)

	m := map[maps.KeyType]maps.ValueType{
		k1: v1,
		k2: v2,
	}

	cm := maps.ToConMapKeyTypeValueType(m)
	is.OK(cm)
	is.Equal(cm.Get(k1), v1)
	is.Equal(cm.Get(k2), v2)

}
