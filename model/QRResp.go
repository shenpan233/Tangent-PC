/*
  @Author:  Trial(Trialpro@gmail.com)
  @Creat:   2022/4/4 8:17
  @Notice:  二维码请求相关
*/

package model

type (
	QRResp struct {
		Status    uint8  `json:"status"`
		QRCode    []byte `json:"bufQRCode"`
		Sig0x30   []byte `json:"bufSig0X30"`
		SigQRSing []byte `json:"bufSigQRSing"`
		BufQRKey  []byte `json:"bufQRKey"`
	}
)
