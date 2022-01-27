/*
  @Author:  Trial(Trialpro@gmail.com)
  @Creat:   2022/1/2 20:19
  @Notice:  消息相关
*/

package model

import "encoding/json"

type MsgTime struct {
	Receive uint64 `json:"Receive"`
	Send    uint64 `json:"Send"`
}
type Font struct {
	Red      uint8  `json:"Red"`
	Blue     uint8  `json:"Blue"`
	Green    uint8  `json:"Green"`
	Size     uint8  `json:"Size"`
	Encoding uint16 `json:"Encoding"`
	FontName string `json:"FontName"`
}

type GroupMsg struct {
	Account   uint64 `json:"Uin"`
	GroupUin  uint64 `json:"GroupUin"` //群号
	SenderUin uint64 `json:"FromUin"`  //发送者QQ
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
