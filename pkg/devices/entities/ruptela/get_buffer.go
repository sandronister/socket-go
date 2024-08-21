package ruptela

import "bytes"

func (d *Device) GetBuffer() *bytes.Buffer {
	return d.Buffer
}
