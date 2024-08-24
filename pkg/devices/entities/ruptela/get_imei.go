package ruptela

func (d *Device) GetImei() string {
	return d.Header.Imei
}
