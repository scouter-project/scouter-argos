package main

import (
	"fmt"

	"github.com/iwindfree/argos/common/dataio"
)

func main() {
	out := dataio.NewDataOutputX()
	out.WriteInt32(10)
	out.WriteInt32(100)
	out.WriteInt32(200)
	var value int32
	out.ReadInt32(&value)
	fmt.Println("the value of  a : %d", value)
	out.ReadInt32(&value)
	fmt.Println("the value of  a : %d", value)
	out.ReadInt32(&value)
	fmt.Println("the value of  a : %d", value)

}
