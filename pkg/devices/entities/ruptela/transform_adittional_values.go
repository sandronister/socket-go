package ruptela

func (d *Device) is65535(in int32) int32 {

	if in == 65535 {
		return 1
	}

	return 2
}

func (d *Device) is102(in float64) int32 {

	if in == 102 {
		return 1
	}

	return 2
}

func (d *Device) is20(in byte) int32 {

	if in == 20 {
		return 1
	}

	return 0
}

func (d *Device) TransformAdditionalValues(hdop byte) {

	d.Parameter.Lls1Status = d.is65535(d.Parameter.Lls1Value)
	d.Parameter.Lls2Status = d.is65535(d.Parameter.Lls2Value)
	d.Parameter.Lls3Status = d.is65535(d.Parameter.Lls3Value)
	d.Parameter.Lls4Status = d.is65535(d.Parameter.Lls4Value)
	d.Parameter.Lls5Status = d.is65535(d.Parameter.Lls5Value)
	d.Parameter.Lls6Status = d.is102(d.Parameter.CanFuelLevel)
	d.Parameter.GpsStatus = d.is20(hdop)

}
