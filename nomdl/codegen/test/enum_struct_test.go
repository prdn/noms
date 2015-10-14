package test

import (
	"testing"

	"github.com/attic-labs/noms/Godeps/_workspace/src/github.com/stretchr/testify/assert"
	"github.com/attic-labs/noms/chunks"
	"github.com/attic-labs/noms/types"
)

func TestEnum(t *testing.T) {
	assert := assert.New(t)

	def := EnumStructDef{Right}
	st := def.New()

	def2 := st.Def()
	st2 := def.New()

	assert.Equal(def, def2)
	assert.True(st.Equals(st2))

	st3 := NewEnumStruct()
	assert.Equal(Right, st3.Hand())
	st3 = st3.SetHand(Left)
	assert.Equal(Left, st3.Hand())
}

func TestEnumValue(t *testing.T) {
	assert := assert.New(t)

	def := EnumStructDef{Switch}
	st := def.New()
	val := st.NomsValue()
	st2 := EnumStructFromVal(val)
	assert.True(st.Equals(st2))
}

func TestEnumIsValue(t *testing.T) {
	cs := chunks.NewMemoryStore()
	var v types.Value = NewEnumStruct()
	ref := types.WriteValue(v, cs)
	v2 := types.ReadValue(ref, cs)
	assert.True(t, v.Equals(v2))
}