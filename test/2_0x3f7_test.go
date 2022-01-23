/*
  @Author:  Trial(Trialpro@gmail.com)
  @Creat:   2022/1/19 16:28
  @Notice:
*/

package test

import (
	"Tangent-PC/protocal/Protobuf/im/cs/cmd0x3f7"
	util "Tangent-PC/utils"
	"fmt"
	"testing"
)

func Test0x3f7(t *testing.T) {
	fmt.Println(util.BinToHex(cmd0x3f7.GetBuffer(959103636, 2471)))

}
