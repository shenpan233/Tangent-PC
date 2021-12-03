/*
* @Author:  Trial
* @email:   shenpan233@vip.qq.com
* @app:		类
* @Creat:   2021/11/26 0026 22:33
 */
package model

type (
	Version struct {
		DwSSOVersion uint32 //Sso版本
		DwPubNo      uint32
		ServiceId    uint32 //客户端ID
		ClientVer    uint32 //客户端版本
		CMainVer     uint16
		CSubVer      uint16 //同CMainVer
	}
	Information struct {
		LongUin  uint64
		Account  string
		PassWord []byte //md5加密
	}

	TeaKey struct {
		Ping0825Key []byte
		PublicKey   []byte
		ShareKey    []byte
	}
)
