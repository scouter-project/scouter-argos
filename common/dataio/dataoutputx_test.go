package dataio

import "testing"

func TestWriteInt(t *testing.T) {
	out := NewDataOutputX()
	out.WriteInt32(100)
	out.WriteInt32(200)
	var a int32
	out.ReadInt32(&a)
	t.Log("the value of a : %d", a)
	out.ReadInt32(&a)
	t.Log("the value of a : %d", a)
}
