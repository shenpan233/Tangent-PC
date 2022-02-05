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

func GetHttpConnSig(Uin uint32, subCmd uint32, sdk model.Version, BufSigHttpConnToken []byte) *HttpConnSig {
	return &HttpConnSig{
		Uin:                 &Uin,
		Tag2:                proto.Uint32(1791),
		Tag3:                proto.Uint32(3088),
		SubCmd:              &subCmd,
		ClientVer:           &sdk.ClientVer,
		DwClientType:        &sdk.DwClientType,
		DwPubNo:             &sdk.DwPubNo,
		ServiceId:           &sdk.ServiceId,
		BufSigHttpConnToken: BufSigHttpConnToken,
		Tag19:               proto.Uint32(0),
		Tag24:               proto.Uint32(0),
		Tag25: &HttpConnSigUnknowTag25{
			Tag1: proto.Uint32(2029),
			Tag2: proto.Uint32(1),
		},
	}
}

func GetReqBody(CommonSig *HttpConnSig, Skey string) []byte {
	req, _ := (&ReqBody{
		HttpConnVer: proto.Uint32(4),
		Sig:         CommonSig,
		Skey: &ReqBody_SigSkey{
			CsVer: proto.Uint32(csVer),
			Skey:  &Skey,
		},
	}).Marshal()
	return req
}
