/*
  @Author:  Trial(Trialpro@gmail.com)
  @Creat:   2021/12/31 23:55
  @Notice:	返回的类
*/

package PCQQ

// QRResp /*二维码相关*/
type (
	QRResp struct {
		Status             uint8
		QRCode             []byte
		sig0x30, sigQRSing []byte
		key                []byte
	}
)
