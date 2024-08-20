package abstract

import (
	"encoding/binary"
	"unsafe"
)

func (d *Device) Read8Bu(output *uint64) {
	binary.Read(d.Buffer, binary.BigEndian, output)
	sz := unsafe.Sizeof(output)
	d.BytesReaded += uint32(sz)
}
