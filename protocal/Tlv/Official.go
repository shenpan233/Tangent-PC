/*
* @Author:  Trial
* @email:   shenpan233@vip.qq.com
* @app:		Official算法
* @Creat:   2021/12/11 20:27
 */
package Tlv

import (
	util "Tangent-PC/utils"
	"bytes"
)

const (
	MD5InfoCount   = 4
	Round          = 256
	TmOffMod       = 19
	TmOffModAdd    = 5
	TxpTEANKeySize = 16
)

/*
未完成 => 8轮的tea还有问题
*/

func CreateOfficial(PasswordMd5, OfficialSig, OfficialKey []byte) []byte {
	MD5Info := util.ToMd5Bytes(OfficialKey)
	MD5Info = append(MD5Info, util.ToMd5Bytes(OfficialSig)...)

	ls, off, seq := make([]byte, Round), make([]byte, Round), make([]byte, Round)
	keyRound := 480%TmOffMod + TmOffModAdd
	for i := 0; i < Round; i++ {
		thisSeq := byte(i - 1)
		seq[i] = thisSeq
		ls[i] = MD5Info[TxpTEANKeySize+(thisSeq%TxpTEANKeySize)]
	}
	id := byte(0) /*一个索引*/
	for i := 0; i < Round; i++ {
		id = uint8(int(id+seq[i]+ls[i]) % Round)
		seq[id+1], seq[i] = seq[i], seq[id+1]
	}
	id = 0
	for i := 0; i < TxpTEANKeySize; i++ {
		id = uint8(int(id+seq[i+1]) % Round)
		seq[id+1], seq[i+1] = seq[i+1], seq[id+1]
		childZid := uint8(int(seq[id+1]+seq[i+1])%Round + 1)
		MD5Info = append(MD5Info, seq[childZid]|MD5Info[i])
	}

	/*最后加密准备工作*/
	MD5Info = append(MD5Info, util.ToMd5Bytes(PasswordMd5)...)
	MD5InfoTwice := util.ToMd5Bytes(MD5Info)
	MO := MD5InfoTwice
	for i := 0; i < keyRound; i++ {
		MO = util.ToMd5Bytes(MO)
	}
	MD5Info = append(MO, MD5Info[TxpTEANKeySize:]...)
	bodyLeft := MD5InfoTwice[:TxpTEANKeySize/2]
	bodyRight := MD5InfoTwice[TxpTEANKeySize/2:]
	for i := 1; i < MD5InfoCount; i++ {
		prekey := MD5Info[(i-1)*TxpTEANKeySize:]
		prekey = prekey[:TxpTEANKeySize+1]
		key := make([]byte, 0)
		for c := 0; c < TxpTEANKeySize; c += 4 {
			key = append(key, prekey[c+3], prekey[c+2], prekey[c+1], prekey[c])
		}
		Md5InfoEncode := bytes.NewBuffer(nil)
		Md5InfoEncode.Write(util.Encrypt(key, bodyLeft))
		Md5InfoEncode.Write(util.Encrypt(key, bodyRight))
		bufMd5InfoEncode := Md5InfoEncode.Bytes()
		for id := i; id < 16; id++ {
			off[id] = off[id] | bufMd5InfoEncode[id]
		}
	}
	return util.ToMd5Bytes(off)
}
