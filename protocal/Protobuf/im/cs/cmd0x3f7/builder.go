/*
  @Author:  Trial(Trialpro@gmail.com)
  @Creat:   2022/1/21 11:38
  @Notice:  构造器
*/

package cmd0x3f7

import (
	"github.com/golang/protobuf/proto"
	"github.com/shenpan233/Tangent-PC/utils/GuBuffer"
)

func GetBuffer(GroupCode uint64, MsgSeq, MsgID uint32) []byte {
	return GuBuffer.NewGuPacketFun(func(pack *GuBuffer.GuPacket) {
		targer := Targer{
			WCsCmdNo: proto.Uint32(wCsCmdNo),
			GroupSub: &Info{
				SubCmd: proto.Uint32(subCmd),
			},
		}
		msg := Msg{
			WCsCmdNo:  proto.Uint32(wCsCmdNo),
			Tag2:      zeroTag,
			GroupCode: &GroupCode,
			Cs: &Cs{
				MsgSeq: &MsgSeq,
				MsgID:  &MsgID,
			},
			Tag5: &Tag5{
				Tag1: proto.Uint32(subCmd),
				Tag2: &Tag5Tag2{
					MsgSeq: &MsgSeq,
					Tag2:   zeroTag,
					Tag3:   proto.Uint32(1),
					Tag4:   zeroTag,
				},
			},
		}
		if bin1, err := targer.Marshal(); err == nil {
			if bin2, err := msg.Marshal(); err == nil {
				pack.SetUint32(uint32(len(bin1)))
				pack.SetUint32(uint32(len(bin2)))
				pack.SetBytes(bin1)
				pack.SetBytes(bin2)
			}
		}

	})
}

func Decode(bin []byte) *Resp {
	response := new(Resp)
	if err := proto.Unmarshal(bin, response); err != nil {
		return nil
	} else {
		return response
	}
}
