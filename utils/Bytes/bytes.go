/*
  @Author:  Trial(Trialpro@gmail.com)
  @Creat:   2022/1/31 19:21
  @Notice:  Bytes转其他类型
*/

package Bytes

func Bytes2Uint8(bin []byte) uint8 {
	return bin[0] & 255
}
