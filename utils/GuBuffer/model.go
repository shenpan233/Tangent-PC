package GuBuffer

import "bytes"

type Tlv struct {
	Tag   uint
	Len   uint
	Value []byte
}
type (
	GuUnPacket struct {
		r *bytes.Reader
	}

	GuPacket struct {
		w   *bytes.Buffer
		tmp []byte
	}
)

type (
	GuPackFun   func(pack *GuPacket)
	GuUnPackFun func(pack *GuUnPacket)
)
