package dataio

import (
	"bytes"
	"encoding/binary"
	"fmt"
)

/*var INT3_MIN_VALUE int = 0xff800000
var INT3_MAX_VALUE int = 0x007fffff
var LONG5_MIN_VALUE int64 = 0xffffff8000000000
var	LONG5_MAX_VALUE int64 = 0x0000007fffffffff */

// A DataOutputX is a output stream which used write various kinds of data.
type DataOutputX struct {
	written int // the wrtten bytes.
	buffer  *bytes.Buffer
}

// NewDataOutputX returns DataOutputX object
func NewDataOutputX() *DataOutputX {
	out := new(DataOutputX)
	out.written = 0
	out.buffer = new(bytes.Buffer)
	return out
}

// WriteInt32 write int32 number to buffer.
func (out *DataOutputX) WriteInt32(value int32) *DataOutputX {
	out.written += 4
	err := binary.Write(out.buffer, binary.BigEndian, value)
	if err != nil {
		fmt.Println("Failed to binary write : ", err)
	}
	return out
}

// WriteInt16 write int16 number to buffer.
func (out *DataOutputX) WriteInt16(value int16) *DataOutputX {
	out.written += 2
	err := binary.Write(out.buffer, binary.BigEndian, value)
	if err != nil {
		fmt.Println("Failed to binary write : ", err)
	}
	return out
}

// WriteInt64 write int64 number to buffer.
func (out *DataOutputX) WriteInt64(value int64) *DataOutputX {
	out.written += 8
	err := binary.Write(out.buffer, binary.BigEndian, value)
	if err != nil {
		fmt.Println("Failed to binary write : ", err)
	}
	return out
}

//ReadInt32 reads int32 number from buffer and assign to value.
func (out *DataOutputX) ReadInt32(value *int32) {
	err := binary.Read(out.buffer, binary.BigEndian, value)
	if err != nil {
		fmt.Println("Faileed to binary read :", err)
		value = nil
	}

}
