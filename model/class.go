/*
* @Author:  Trial
* @email:   shenpan233@vip.qq.com
* @app:		类
* @Creat:   2021/11/26 0026 22:33
 */
package model

import "container/list"

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
		PingTime uint32
		Computer
	}

	TeaKey struct {
		Ping0825Key []byte
		Ping0818Key []byte
		PublicKey   []byte
		ShareKey    []byte
	}

	/*一些token/sign*/
	Sig struct {
		BufSigClientAddr []byte /*0825返回*/
		BufTgTGTKey      []byte
		BufQR303         []byte
	}

	/*Computer*/
	Computer struct {
		ComputerId   []byte
		ComputerIdEx []byte
		ComputerName string
		MacGuid      []byte
		RedirectIp   *list.List
		ConnectIp    string
		WlanIp       string
	}
)
