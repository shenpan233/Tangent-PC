/*
  @Author:  Trial(Trialpro@gmail.com)
  @Creat:   2022/2/4 12:31
  @Notice:  tag4所有httpConn请求都有的客户端信息
*/

package HttpConn

import (
	"github.com/golang/protobuf/proto"
	"github.com/shenpan233/Tangent-PC/model"
)

const (
	csVer = 1
)

func GetHttpConnSig(Uin uint32, subCmd uint32, sdk model.Version, BufSigHttpConnToken []byte) *MsgHttpConnHead {
	return &MsgHttpConnHead{
		Uin:                 &Uin,
		Command:             proto.Uint32(1791),
		SubCommand:          proto.Uint32(3088),
		Seq:                 &subCmd,
		ClientVer:           &sdk.ClientVer,
		DwClientType:        &sdk.DwClientType,
		DwPubNo:             &sdk.DwPubNo,
		ServiceId:           &sdk.ServiceId,
		BufSigHttpConnToken: BufSigHttpConnToken,
		Flag:                proto.Uint32(0),
		CompressType:        proto.Uint32(0),
		MsgOiDbHead: &MsgHttpConnHeadMsgOidbhead{
			OidbCommand: proto.Uint32(2029),
			ServiceType: proto.Uint32(1),
		},
	}
}

func GetReqBody(MsgHttpConnHead *MsgHttpConnHead, Skey string) []byte {
	req, _ := (&ReqBody{
		HeadType:        proto.Uint32(4),
		MsgHttpConnHead: MsgHttpConnHead,
		Skey: &ReqBodyPSkeyBuf{
			Type: proto.Uint32(csVer),
			Sig:  &Skey,
		},
	}).Marshal()
	return req
}
