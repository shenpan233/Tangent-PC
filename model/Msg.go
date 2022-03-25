/*
  @Author:  Trial(Trialpro@gmail.com)
  @Creat:   2022/1/2 20:19
  @Notice:  消息相关的类
*/

package model

import (
	json "github.com/json-iterator/go"
	"github.com/shenpan233/Tangent-PC/utils/GuBuffer"
)

type MsgTime struct {
	Receive uint64 `json:"Receive"`
	Send    uint64 `json:"Send"`
}

const (
	FontNameMicrosoftYaHei = "微软雅黑"
	FontEncodingUTF8       = 0x86_22
	FontEncodingEN         = 0x00_00
	FontEncodingGBK        = 0x86_00
)

type Font struct {
	Red      uint8  `json:"Red"`
	Blue     uint8  `json:"Blue"`
	Green    uint8  `json:"Green"`
	Size     uint8  `json:"Size"`
	Style    uint8  `json:"Style"`
	Encoding uint16 `json:"Encoding"`
	FontName string `json:"FontName"`
}

type GroupMsg struct {
	Account   uint64 `json:"Account"`
	GroupName string
	GroupUin  uint64 `json:"GroupUin"` //群号
	SenderUin uint64 `json:"FromUin"`  //发送者QQ
	FromName  string
	MsgID     uint32 `json:"MsgID"`
	MsgSeq    uint32 `json:"MsgSeq"`
	Msg       string `json:"Msg"`
	MsgTime   `json:"Time"`
	Font      `json:"Font"`
	dataJson
}

func Msg2Json(msg *GroupMsg) string {
	if msg != nil {
		marshal, err := json.Marshal(msg)
		if err != nil {
			return ""
		}
		return string(marshal)
	} else {
		return ""
	}
}

func (this *Font) ToBytes() []byte {
	return GuBuffer.NewGuPacketFun(func(pack *GuBuffer.GuPacket) {
		pack.SetUint8(this.Red)
		pack.SetUint8(this.Blue)
		pack.SetUint8(this.Green)
		pack.SetUint8(this.Size)
		pack.SetUint8(this.Style)
		pack.SetUint16(this.Encoding)
		pack.SetSToken(this.FontName)
	})
}
