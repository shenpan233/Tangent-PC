/*
  @Author:  Trial(Trialpro@gmail.com)
  @Creat:   2022/1/23 22:01
  @Notice:
*/

package PCQQ

import (
	"Tangent-PC/protocal/Protobuf/im/cs/cmd0x3f7"
	"fmt"
)

func (this *TangentPC) pack0x3f7(GroupCode uint64, MsgSeq uint32) []byte {
	buffer := cmd0x3f7.GetBuffer(GroupCode, MsgSeq)
	fmt.Println(buffer)
	return buffer
}
