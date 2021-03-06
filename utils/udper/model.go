/*
* @Author:  Trial
* @email:   shenpan233@vip.qq.com
* @app:		类定义
* @Creat:   2021/11/27 0027 12:03
 */
package udper

import (
	"context"
	"github.com/shenpan233/Tangent-PC/utils/GuBuffer"
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

	UdpRecv func(Cmd uint16, seq uint16, pack *GuBuffer.GuUnPacket)
)
