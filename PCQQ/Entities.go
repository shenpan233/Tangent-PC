/*
* @Author:  Trial
* @email:   shenpan233@vip.qq.com
* @app:		返回的类
* @Creat:   2021/12/4 0004 16:24
 */
package PCQQ

/*二维码相关*/
type (
	QRResp struct {
		Status uint8
		QRCode []byte
	}
)
