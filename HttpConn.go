/*
  @Author:  Trial(Trialpro@gmail.com)
  @Creat:   2022/2/25 22:59
  @Notice:  HttpConn操作
*/

package Tangent_PC

import "github.com/shenpan233/Tangent-PC/protocal/Protobuf/HttpConn"

func (this *TangentPC) GetCommonHttpConnSig(subCmd uint32) []byte {
	return HttpConn.GetReqBody(HttpConn.GetHttpConnSig(uint32(this.info.LongUin), subCmd, *this.sdk, this.sig.BufSigHttpConnToken), this.info.SelfWebKey.Common.Skey)
}
