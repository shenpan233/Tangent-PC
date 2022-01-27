/*
  @Author:  Trial(Trialpro@gmail.com)
  @Creat:   2022/1/8 10:10
  @Notice:  接口函数
*/

package model

type dataJson interface {
	Encode() string
	Decode(data string) bool
}
