package ruptela

func (d *Device) IsValidSize() bool {
	crcSize := 2
	packLenSize := 2

	a := d.Header.PackageLen
	b := d.BytesReaded - uint32(
		crcSize+packLenSize,
	)
	return a == b
}
