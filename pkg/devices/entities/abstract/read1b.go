package abstract

import (
	"encoding/binary"
	"unsafe"
)

func (d *Device) Read1B(output *uint8) {
	binary.Read(d.Buffer, binary.BigEndian, output)
	sz := unsafe.Sizeof(*output)
	d.BytesReaded += uint32(sz)
}
