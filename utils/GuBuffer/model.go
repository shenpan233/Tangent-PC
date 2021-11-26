package GuBuffer

import "bytes"

type Tlv struct {
	Tag   uint
	Len   uint
	Value []byte
}
type GuPacket struct {
	w   *bytes.Buffer
	tmp []byte
}
type GuUnPacket struct {
	r *bytes.Reader
}
