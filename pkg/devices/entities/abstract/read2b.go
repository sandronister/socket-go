package abstract

import (
	"encoding/binary"
	"unsafe"
)

func (d *Device) Read2B(output *int16) {
	binary.Read(d.Buffer, binary.BigEndian, output)
	sz := unsafe.Sizeof(*output)
	d.BytesReaded += uint32(sz)
}
