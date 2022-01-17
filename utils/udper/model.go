/*
* @Author:  Trial
* @email:   shenpan233@vip.qq.com
* @app:		类定义
* @Creat:   2021/11/27 0027 12:03
 */
package udper

import (
	"Tangent-PC/utils/GuBuffer"
	"context"
	"net"
	"sync"
)

type (
	Udper struct {
		conn net.Conn
		seq  uint32
		*Set
		pull *sync.Map
		context.Context
		context.CancelFunc
	}

	Set struct {
		BuffMaxSize int
		UdpRecv
	}

	UdpRecv func(Cmd int16, seq uint16, pack *GuBuffer.GuUnPacket)
)
