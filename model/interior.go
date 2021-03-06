/*
* @Author:  Trial
* @email:   shenpan233@vip.qq.com
* @app:     内部的类
* @Creat:   2021/11/26 0026 22:33
 */

package model

import "container/list"

type (
	Version struct {
		SAppName     string
		DwSSOConfig  uint32
		DwSSOVersion uint32 //Sso版本
		DwPubNo      uint32
		ServiceId    uint32 //客户端ID
		DwAppVer     uint32 //客户端版本
		CMainVer     uint16
		CSubVer      uint16 //同CMainVer
		DwQdVersion  uint32
		DwClientType uint32
		ClientMd5    []byte //客户端MD5

	}

	Information struct {
		LongUin    uint64
		Account    string
		PassWord   []byte //md5加密
		PingTime   uint32
		SelfWebKey *WebKey
		Computer
	}

	TeaKey struct {
		Ping0825Key []byte
		Ping0818Key []byte
		PublicKey   []byte
		ShareKey    []byte
		SessionKey  []byte
		HttpConn    []byte
	}

	// Sig 一些token/sign
	Sig struct {
		BufSigClientAddr    []byte /*0825返回*/
		BufTgTGTKey         []byte
		BufQR303            []byte
		BufSession          []byte
		BufPwdForConn       []byte
		BufSigHttpConnToken []byte
		Buf0102             []byte
		Buf0202             []byte
		BufTgt              []byte
	}

	//Computer 硬件信息
	Computer struct {
		ComputerId   []byte `json:"ComputerId"`
		ComputerIdEx []byte `json:"ComputerIdEx"`
		DeviceID     []byte `json:"DeviceID"`
		ComputerName string `json:"ComputerName"`
		MacGuid      []byte `json:"MacGuid"`
		RedirectIp   *list.List
		ConnectIp    string `json:"ConnectIp"`
		WlanIp       string
	}
)
