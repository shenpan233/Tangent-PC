/*
  @Author:  Trial(Trialpro@gmail.com)
  @Creat:   2022/4/16 19:51
  @Notice:  Friend.Builder
*/

package Msg

type Builder interface {
	Generate() []byte
	Decode() string
}
