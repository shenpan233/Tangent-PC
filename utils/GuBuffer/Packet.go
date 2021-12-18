package GuBuffer

import (
	util "Tangent-PC/utils"
	"bytes"
)

func NewGuPacket() (p *GuPacket) {
	p = new(GuPacket)
	p.w = new(bytes.Buffer)
	return
}

//仿Mirai,这种还挺好用的
func NewGuPacketFun(fun GuPackFun) []byte {
	pack := NewGuPacket()
	if fun != nil {
		fun(pack)
	}
	return pack.GetAll()
}

func (p *GuPacket) Reset() {
	p.w = new(bytes.Buffer)
}

func (p *GuPacket) JmpHead() {
	p.tmp = p.w.Bytes()
	p.w.Reset()
}

//func (p *GuPacket) SetHex(hex string) {
//	p.SetBytes(util.HexToBin(hex))
//}
func (p *GuPacket) SetBytes(bin []byte) {
	p.w.Write(bin)
}

func (p *GuPacket) SetToken(bin []byte) {
	p.SetUint16(uint16(len(bin)))
	p.SetBytes(bin)
}

func (p *GuPacket) SetSToken(bin string) {
	p.SetUint16(uint16(len(bin)))
	p.SetBytes([]byte(bin))
}

func (p *GuPacket) SetString(bin string) {
	p.w.WriteString(bin)
}

func (p *GuPacket) SetTlv(t *Tlv) {
	p.SetUint16(uint16(t.Tag))
	p.SetUint16(uint16(len(t.Value)))
	p.w.Write(t.Value)
}

func (p *GuPacket) SetUint8(i uint8) {
	p.w.Write([]byte{(byte)(i >> 0)})
}

func (p *GuPacket) SetUint16(i uint16) {
	p.w.Write([]byte{(byte)(i >> 8), (byte)(i >> 0)})
}

func (p *GuPacket) SetUint32(i uint32) {
	p.w.Write([]byte{(byte)(i >> 24), (byte)(i >> 16), (byte)(i >> 8), (byte)(i >> 0)})
}

func (p *GuPacket) SetUint64(i uint64) {
	p.w.Write([]byte{(byte)(i >> 56), (byte)(i >> 48), (byte)(i >> 40), (byte)(i >> 32), (byte)(i >> 24), (byte)(i >> 16), (byte)(i >> 8), (byte)(i >> 0)})
}

func (p *GuPacket) GetAll() (bin []byte) {
	p.w.Write(p.tmp)
	p.tmp = nil
	bin = p.w.Bytes()
	return
}

func (p *GuPacket) GetHex() string {
	return util.BinToHex(p.GetAll())
}

/**/
func (p *GuPacket) ToTlv(t uint16) (bin []byte) {
	bin = p.GetAll()
	p.Reset()
	p.SetUint16(t)
	p.SetUint16(uint16(len(bin)))
	p.SetBytes(bin)
	return p.GetAll()
}
