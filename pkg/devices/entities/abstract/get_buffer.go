package abstract

func (d *Device) GetBuffer() []byte {
	return d.OriginalBuff
}
