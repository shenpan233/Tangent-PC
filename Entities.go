/*
  @Author:  Trial(Trialpro@gmail.com)
  @Creat:   2021/12/31 23:55
  @Notice:  返回的类
*/

package Tangent_PC

import "github.com/shenpan233/Tangent-PC/model"

// QRResp /*二维码相关*/
type (
	QRResp struct {
		Status    uint8  `json:"status"`
		QRCode    []byte `json:"bufQRCode"`
		Sig0x30   []byte `json:"bufSig0X30"`
		SigQRSing []byte `json:"bufSigQRSing"`
		BufQRKey  []byte `json:"bufQRKey"`
	}

	// HOOK 回调钩子
	HOOK struct {
		GroupMsg func(Msg model.GroupMsg) //群消息接收
	}
)
