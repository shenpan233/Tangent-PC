/*
  @Author:  Trial(Trialpro@gmail.com)
  @Creat:   2022/1/28 11:53
  @Notice:  数据校验
*/

package Bytes

import (
	"crypto/md5"
	"hash/crc32"
)

func GetMd5Bytes(data []byte) (ret []byte) {
	tmp := md5.Sum(data)
	ret = tmp[:]
	return
}

func GetCrc32(data []byte) []byte {
	ieee := crc32.ChecksumIEEE(data)
	return []byte{byte(ieee >> 0), byte(ieee >> 8), byte(ieee >> 16), byte(ieee >> 24)}
}
