/*
  @Author:  Trial(Trialpro@gmail.com)
  @Creat:   2022/1/29 15:53
  @Notice:  消息撤回的测试
*/

package test

import (
	"fmt"
	"github.com/shenpan233/Tangent-PC/protocal/Protobuf/im/cs/cmd0x3f7"
	util "github.com/shenpan233/Tangent-PC/utils"
	"testing"
)

func TestCmd0x3f7(t *testing.T) {
	fmt.Println(util.BinToHex(cmd0x3f7.GetBuffer(959103636, 123, 456)))
}
