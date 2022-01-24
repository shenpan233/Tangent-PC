/*
  @Author:  Trial(Trialpro@gmail.com)
  @Creat:   2022/1/21 11:38
  @Notice:
*/

package cmd0x3f7

import (
	util "Tangent-PC/utils"
	"Tangent-PC/utils/GuBuffer"
	"github.com/golang/protobuf/proto"
)

func GetBuffer(GroupCode uint64, MsgSeq uint32) []byte {
	return GuBuffer.NewGuPacketFun(func(pack *GuBuffer.GuPacket) {
		targer := Targer{
			WCsCmdNo: proto.Uint32(wCsCmdNo),
			GroupSub: &Info{
				SubCmd:      proto.Uint32(subCmd),
				GroupType:   proto.Uint32(uint32(util.GetRand32())),
				DwRequestId: proto.Uint64(1),
			},
		}
		msg := Msg{
			WCsCmdNo:  proto.Uint32(wCsCmdNo),
			Tag2:      zeroTag,
			GroupCode: &GroupCode,
			Cs: &Cs{
				MsgSeq:    &MsgSeq,
				MsgRandom: proto.Uint32(uint32(util.GetRand32())),
			},
			Tag5: &Tag5{
				Tag1: zeroTag,
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
