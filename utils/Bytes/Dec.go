/*
  @Author:  Trial(Trialpro@gmail.com)
  @Creat:   2022/1/28 11:47
  @Notice:  数字类型转hex
*/

package Bytes

//Uint32ToBytes 拆出来的函数
func Uint32ToBytes(i uint32) []byte {
	return []byte{(byte)(i >> 24), (byte)(i >> 16), (byte)(i >> 8), (byte)(i >> 0)}
}
