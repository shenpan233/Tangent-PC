package GuBuffer

import (
	"bytes"
	util "github.com/shenpan233/Tangent-PC/utils"
	"os"
)

func NewGuUnPacket(bin []byte) (u *GuUnPacket) {
	if len(bin) == 0 {
		return nil
	}
	u = new(GuUnPacket)
	u.r = bytes.NewReader(bin)
	return
}

func (r *GuUnPacket) GetLen() int {
	if r.r == nil {
		return 0
	}
	return r.r.Len()
}

func (r *GuUnPacket) GetAll() []byte {
	return r.GetBin(r.GetLen())
}

func (r *GuUnPacket) GetAllHex() string {
	bin := r.GetAll()
	if bin == nil {
		return ""
	}
	return util.BinToHex2(&bin)
}

func (r *GuUnPacket) GetAllHexTmp() string {
	bin := r.GetAll()
	if bin == nil {
		return ""
	}
	r.r.Seek(-int64(len(bin)), 1)
	return util.BinToHex2(&bin)
}

func (r *GuUnPacket) GetBin(i int) []byte {
	if i <= 0 || r.GetLen() < i {
		return nil
	}
	data := make([]byte, i)
	if _, err := r.r.Read(data); err != nil {
		return nil
	}
	return data
}

func (r *GuUnPacket) GetStr(i int32) string {
	return string(r.GetBin(int(i)))
}

func (r *GuUnPacket) GetUint8() uint8 {
	if bArr := r.GetBin(1); bArr != nil {
		return bArr[0] & 255
	} else {
		return 0
	}
}

func (r *GuUnPacket) GetUint16() uint16 {
	//int8(bArr[0] & 255)
	if bArr := r.GetBin(2); bArr != nil {
		return uint16(((int(bArr[0]) << 8) & 65280) + ((int(bArr[1]) << 0) & 255))
	} else {
		return 0
	}
}

func (r *GuUnPacket) GetUint32() uint32 {
	if bArr := r.GetBin(4); bArr != nil {
		return uint32((int(bArr[0])<<24)&-16777216 + ((int(bArr[1]) << 16) & 16711680) + ((int(bArr[2]) << 8) & 65280) + (int(bArr[3]<<0) & 255))
	} else {
		return 0
	}
}

func (r *GuUnPacket) GetToken() []byte {
	if r.GetLen() < 2 {
		return make([]byte, 0)
	}
	return r.GetBin(int(r.GetUint16()))
}

func (r *GuUnPacket) GetTlv() (t *Tlv) {
	if r.GetLen() < 4 {
		return nil
	}
	t = &Tlv{
		Tag:   uint(r.GetUint16()),
		Len:   uint(r.GetUint16()),
		Value: nil,
	}
	t.Value = r.GetBin(int(t.Len))
	return
}

func (this *GuUnPacket) Skip(len int64) {
	this.r.Seek(len, os.SEEK_CUR)
}

// NewGuUnPacketFun 仿Mirai,这种还挺好用的
func NewGuUnPacketFun(Buffer []byte, fun GuUnPackFun) *GuUnPacket {
	pack := NewGuUnPacket(Buffer)
	if pack == nil {
		return nil
	}
	if fun != nil {
		fun(pack)
	}
	return pack
}

//TlvEnum Tlv结构枚举
func TlvEnum(bin []byte, store map[uint16]func(pack *GuUnPacket)) {
	NewGuUnPacketFun(bin, func(pack *GuUnPacket) {
		for pack.GetLen() > 0 {
			Tag := pack.GetUint16()
			if callBack := store[Tag]; callBack != nil {
				callBack(NewGuUnPacket(pack.GetToken()))
			} else {
				//无需处理即跳过
				pack.Skip(int64(pack.GetUint16()))
			}
		}
	})
}
